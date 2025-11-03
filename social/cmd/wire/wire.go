//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	auth_context "lib/auth/context"
	"lib/kafka"
	lib_grpc_middleware "lib/middleware/grpc"
	"lib/postgres"
	"lib/postgres/transaction_manager"

	social_controller "social/internal/app/controllers/social"
	social_repository "social/internal/app/repositories/social"
	"social/internal/app/server"
	"social/internal/app/usecases/social"
	grpc_middleware "social/internal/middleware/grpc"
	socialv1 "social/pkg/api/social/v1"

	friend_request_events_handler "social/internal/app/adapters/friend_request_events_handler"
	"social/internal/app/modules/outbox"
	outbox_repository "social/internal/app/repositories/outbox"

	"github.com/IBM/sarama"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type App struct {
	Server *server.Server
	Worker *outbox.OutboxFriendRequestWorker
}

func InitializeApp(ctx context.Context) (*App, error) {
	wire.Build(
		// ========== infra ==========
		newPostgresConnection,
		newTransactionManager,

		wire.Bind(new(postgres.QueryEngineProvider), new(*transaction_manager.TransactionManager)),
		wire.Bind(new(social.TransactionManager), new(*transaction_manager.TransactionManager)),

		// ========== Repositories ==========
		social_repository.NewRepository,
		wire.Bind(new(social.SocialRepository), new(*social_repository.Repository)),

		// ===== outbox =====
		newOutboxRepository,
		newKafkaProducer,
		newFriendRequestEventsHandler,
		newOutboxWorker,
		newOutboxProcessor,
		wire.Bind(new(social.OutboxRepository), new(*outbox.Processor)),

		// ========== UserIDProvider ==========
		newUserIDProvider,

		// ========== Usecases ==========
		wire.Struct(new(social.Deps), "*"),
		social.NewUsecase,

		// ========== Controllers ==========
		wire.Struct(new(social_controller.Deps), "*"),
		social_controller.New,
		wire.Bind(new(socialv1.SocialServiceServer), new(*social_controller.Controller)),

		// ========== Middlewares ==========
		newValidator,
		newMiddlewares,

		// ========== Server ==========
		newConfig,
		wire.Struct(new(server.Contollers), "*"),
		server.New,

		// ===== wrap в App =====
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

// ========== Providers ==========

func newValidator() (*protovalidate.Validator, error) {
	return protovalidate.New(protovalidate.WithDisableLazy(false))
}

func newMiddlewares(validator *protovalidate.Validator) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_middleware.ErrorsUnaryServerInterceptor(),
		lib_grpc_middleware.ValidateUnaryServerInterceptor(validator),
	}
}

func newConfig(mws []grpc.UnaryServerInterceptor) server.Config {
	grpcListenPort := os.Getenv("GRPC_LISTEN_PORT")

	return server.Config{
		GRPCPort:               ":" + grpcListenPort,
		ChainUnaryInterceptors: mws,
	}
}

func newUserIDProvider() social.UserIDProvider {
	return auth_context.MyUserIDProvider{}
}

func newPostgresConnection(ctx context.Context) (*postgres.Connection, error) {
	pgUser := os.Getenv("SOCIAL_POSTGRES_USER")
	pgPassword := os.Getenv("SOCIAL_POSTGRES_PASSWORD")
	pgDB := os.Getenv("SOCIAL_POSTGRES_DB")
	pgHost := os.Getenv("SOCIAL_POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgSSLmode := os.Getenv("POSTGRES_SSLMODE")

	return postgres.NewConnectionPool(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			pgUser, pgPassword, pgHost, pgPort, pgDB, pgSSLmode),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
}

func newTransactionManager(pool *postgres.Connection) *transaction_manager.TransactionManager {
	return transaction_manager.New(pool)
}

func newOutboxRepository(txManager *transaction_manager.TransactionManager) *outbox_repository.Repository {
	return outbox_repository.NewRepository(txManager)
}

func newKafkaProducer() (sarama.SyncProducer, error) {
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	return kafka.NewNewSyncProducer(strings.Split(kafkaBrokers, ","), nil)
}

func TopicResolver(e *outbox.Event) (topic string, key string) {
	friendRequestTopicName := os.Getenv("KAFKA_SOCIAL_FRIEND_REQUEST_TOPIC_NAME")
	friendUpdatedTopicName := os.Getenv("KAFKA_SOCIAL_FRIEND_UPDATED_TOPIC_NAME")

	switch e.EventType {
	case outbox.EventTypeFriendRequestCreated:
		return friendRequestTopicName, e.AggregateID
	case outbox.EventTypeFriendRequestUpdated:
		return friendUpdatedTopicName, e.AggregateID
	}

	return "social.unknown.events", e.AggregateID
}

func newFriendRequestEventsHandler(producer sarama.SyncProducer) *friend_request_events_handler.KafkaFriendRequestBatchHandler {
	return friend_request_events_handler.NewKafkaFriendRequestBatchHandler(producer,
		friend_request_events_handler.WithMaxBatchSize(100),
		friend_request_events_handler.WithTopicResolver(TopicResolver),
	)
}

func newOutboxWorker(
	ctx context.Context,
	outboxRepo *outbox_repository.Repository,
	txManager *transaction_manager.TransactionManager,
	handler *friend_request_events_handler.KafkaFriendRequestBatchHandler,
) *outbox.OutboxFriendRequestWorker {
	return outbox.NewOutboxFriendRequestWorker(
		outboxRepo,
		txManager,
		handler,
		outbox.WithBatchSize(10),
		outbox.WithMaxRetry(10),
		outbox.WithRetryInterval(30*time.Second),
		outbox.WithWindow(time.Hour),
	)
}

func newOutboxProcessor(outboxRepo *outbox_repository.Repository) *outbox.Processor {
	return outbox.NewProcessor(outbox.Deps{
		Repository: outboxRepo,
	})
}
