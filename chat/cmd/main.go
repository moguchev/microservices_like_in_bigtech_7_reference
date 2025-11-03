package main

import (
	"context"
	"log"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	_ = container.Provide(provideContext)
	_ = container.Provide(providePostgresConnection)
	_ = container.Provide(provideTransactionManager)
	_ = container.Provide(provideChatRepository)
	_ = container.Provide(provideChatUsecase)
	_ = container.Provide(provideChatController)
	_ = container.Provide(provideValidator)
	_ = container.Provide(provideGRPCMiddlewares)
	_ = container.Provide(provideServerConfig)
	_ = container.Provide(provideServer)
	_ = container.Provide(provideOutboxRepository)
	_ = container.Provide(provideOutboxProcessor)
	_ = container.Provide(provideKafkaProducer)
	_ = container.Provide(provideChatMessageEventsHandler)
	_ = container.Provide(provideOutboxWorker)
	_ = container.Provide(provideApp)

	err := container.Invoke(func(ctx context.Context, cancel context.CancelFunc, app *App) {
		go app.Worker.Run(ctx)

		defer cancel()
		if err := app.Server.Run(ctx); err != nil {
			log.Fatalf("run server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("failed to invoke container: %v", err)
	}
}
