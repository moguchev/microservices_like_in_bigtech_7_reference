package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lib/postgres"
	"lib/postgres/transaction_manager"
	"notification/internal/app/adapters/consumer"
	"notification/internal/app/modules/inbox"
	inboxrepository "notification/internal/app/repositories/inbox"
	"notification/internal/app/usecases/notification"
)

var (
	pgUser     = os.Getenv("NOTIFICATION_POSTGRES_USER")
	pgPassword = os.Getenv("NOTIFICATION_POSTGRES_PASSWORD")
	pgDB       = os.Getenv("NOTIFICATION_POSTGRES_DB")
	pgHost     = os.Getenv("NOTIFICATION_POSTGRES_HOST")
	pgPort     = os.Getenv("POSTGRES_PORT")
	pgSSLmode  = os.Getenv("POSTGRES_SSLMODE")

	kafkaBrokers                      = os.Getenv("KAFKA_BROKERS")
	kafkaSocialFriendRequestTopicName = os.Getenv("KAFKA_SOCIAL_FRIEND_REQUEST_TOPIC_NAME")
	kafkaSocialFriendUpdatedTopicName = os.Getenv("KAFKA_SOCIAL_FRIEND_UPDATED_TOPIC_NAME")
	kafkaChatMessageSentTopicName     = os.Getenv("KAFKA_CHAT_MESSAGE_SENT_TOPIC_NAME")

	kafkaConsumerGroup = os.Getenv("KAFKA_CONSUMER_GROUP")
	kafkaConsumerName  = os.Getenv("KAFKA_CONSUMER_NAME")
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// infra
	pgConn, err := postgres.NewConnectionPool(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgUser, pgPassword, pgHost, pgPort, pgDB, pgSSLmode),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatalf("failed to create connection pool: %v", err)
	}

	// adapters/repositories
	dedup := consumer.NewInMemoryDeduper(ctx, 24*time.Hour)
	txManager := transaction_manager.New(pgConn)
	inboxRepo := inboxrepository.NewRepository(txManager)

	// usecases
	handler := notification.NewUsecase()

	worker := inbox.NewInboxWorker(inboxRepo, txManager, handler,
		inbox.WithBatchSize(10),
		inbox.WithMaxAttempts(10),
		inbox.WithPollInterval(10*time.Second),
	)

	go worker.Run(ctx)

	processor := inbox.NewProcessor(inbox.Deps{
		Repository: inboxRepo,
	})

	consumer, err := consumer.NewInboxConsumer([]string{kafkaBrokers},
		kafkaConsumerGroup,
		kafkaConsumerName,
		dedup,
		processor,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()
	if err := consumer.Run(
		ctx,
		kafkaSocialFriendRequestTopicName,
		kafkaSocialFriendUpdatedTopicName,
		kafkaChatMessageSentTopicName,
	); err != nil && ctx.Err() == nil {
		log.Println("consumer stopped:", err)
	}

	log.Println("shutdown")
}
