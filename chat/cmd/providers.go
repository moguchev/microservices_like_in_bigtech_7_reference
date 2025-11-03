package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	chat_message_events_handler "chat/internal/app/adapters/chat_message_events_handler"
	chat_controller "chat/internal/app/controllers/chat"
	"chat/internal/app/modules/outbox"
	chat_repository "chat/internal/app/repositories/chat"
	outbox_repository "chat/internal/app/repositories/outbox"
	"chat/internal/app/server"
	"chat/internal/app/usecases/chat"
	grpc_middleware "chat/internal/middleware/grpc"
	auth_context "lib/auth/context"
	"lib/kafka"
	lib_grpc_middleware "lib/middleware/grpc"
	"lib/postgres"
	"lib/postgres/transaction_manager"

	"github.com/IBM/sarama"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

var (
	grpcListenPort = os.Getenv("GRPC_LISTEN_PORT")

	pgUser     = os.Getenv("CHAT_POSTGRES_USER")
	pgPassword = os.Getenv("CHAT_POSTGRES_PASSWORD")
	pgDB       = os.Getenv("CHAT_POSTGRES_DB")
	pgHost     = os.Getenv("CHAT_POSTGRES_HOST")
	pgPort     = os.Getenv("POSTGRES_PORT")
	pgSSLmode  = os.Getenv("POSTGRES_SSLMODE")

	kafkaBrokers         = os.Getenv("KAFKA_BROKERS")
	messageSentTopicName = os.Getenv("KAFKA_CHAT_MESSAGE_SENT_TOPIC_NAME")
)

type App struct {
	Server *server.Server
	Worker *outbox.OutboxChatMessageWorker
}

func provideContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func providePostgresConnection(ctx context.Context) (*postgres.Connection, error) {
	return postgres.NewConnectionPool(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			pgUser, pgPassword, pgHost, pgPort, pgDB, pgSSLmode),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
}

func provideTransactionManager(pool *postgres.Connection) *transaction_manager.TransactionManager {
	return transaction_manager.New(pool)
}

func provideChatRepository(tm *transaction_manager.TransactionManager) *chat_repository.Repository {
	return chat_repository.NewRepository(tm, tm)
}

func provideChatUsecase(repo *chat_repository.Repository, tm *transaction_manager.TransactionManager, outboxProc *outbox.Processor) chat.Usecase {
	return chat.NewUsecase(chat.Deps{
		ChatRepository:     repo,
		TransactionManager: tm,
		OutboxRepository:   outboxProc,
		UserIDProvider:     auth_context.MyUserIDProvider{},
	})
}

func provideChatController(usecase chat.Usecase) *chat_controller.Controller {
	return chat_controller.New(chat_controller.Deps{
		ChatUsecase: usecase,
	})
}

func provideValidator() (*protovalidate.Validator, error) {
	return protovalidate.New(protovalidate.WithDisableLazy(false))
}

func provideGRPCMiddlewares(validator *protovalidate.Validator) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_middleware.ErrorsUnaryServerInterceptor(),
		lib_grpc_middleware.ValidateUnaryServerInterceptor(validator),
	}
}

func provideServerConfig(mws []grpc.UnaryServerInterceptor) server.Config {
	return server.Config{
		GRPCPort:               ":" + grpcListenPort,
		ChainUnaryInterceptors: mws,
	}
}

func provideServer(
	ctx context.Context,
	cfg server.Config,
	chatController *chat_controller.Controller,
) (*server.Server, error) {
	return server.New(ctx, cfg, server.Contollers{
		ChatServiceServer: chatController,
	})
}

func provideOutboxRepository(tm *transaction_manager.TransactionManager) *outbox_repository.Repository {
	return outbox_repository.NewRepository(tm)
}

func provideOutboxProcessor(outboxRepo *outbox_repository.Repository) *outbox.Processor {
	return outbox.NewProcessor(outbox.Deps{
		Repository: outboxRepo,
	})
}

func provideKafkaProducer() (sarama.SyncProducer, error) {
	return kafka.NewNewSyncProducer(strings.Split(kafkaBrokers, ","), nil)
}

func provideChatMessageEventsHandler(producer sarama.SyncProducer) *chat_message_events_handler.KafkaChatMessageBatchHandler {
	return chat_message_events_handler.NewKafkaChatMessageBatchHandler(producer,
		chat_message_events_handler.WithMaxBatchSize(100),
		chat_message_events_handler.WithTopic(messageSentTopicName),
	)
}

func provideOutboxWorker(
	ctx context.Context,
	outboxRepo *outbox_repository.Repository,
	txManager *transaction_manager.TransactionManager,
	handler *chat_message_events_handler.KafkaChatMessageBatchHandler,
) *outbox.OutboxChatMessageWorker {
	return outbox.NewOutboxChatMessageWorker(
		outboxRepo,
		txManager,
		handler,
		outbox.WithBatchSize(10),
		outbox.WithMaxRetry(10),
		outbox.WithRetryInterval(30*time.Second),
		outbox.WithWindow(time.Hour),
	)
}

func provideApp(srv *server.Server, worker *outbox.OutboxChatMessageWorker) *App {
	return &App{
		Server: srv,
		Worker: worker,
	}
}
