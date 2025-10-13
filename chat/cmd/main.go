package main

import (
	"context"
	"log"

	"chat/internal/app/server"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	_ = container.Provide(provideContext)
	_ = container.Provide(provideChatRepository)
	_ = container.Provide(provideChatUsecase)
	_ = container.Provide(provideChatController)
	_ = container.Provide(provideValidator)
	_ = container.Provide(provideGRPCMiddlewares)
	_ = container.Provide(provideServerConfig)
	_ = container.Provide(provideServer)

	err := container.Invoke(func(ctx context.Context, cancel context.CancelFunc, srv *server.Server) {
		defer cancel()
		if err := srv.Run(ctx); err != nil {
			log.Fatalf("run server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("failed to invoke container: %v", err)
	}
}
