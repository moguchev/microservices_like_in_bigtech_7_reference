package main

import (
	"context"

	chat_controller "chat/internal/app/controllers/chat"
	chat_repository "chat/internal/app/repositories/chat"
	"chat/internal/app/server"
	"chat/internal/app/usecases/chat"
	grpc_middleware "chat/internal/middleware/grpc"
	auth_context "lib/auth/context"
	lib_grpc_middleware "lib/middleware/grpc"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

const address = ":8080"

func provideContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func provideChatRepository() *chat_repository.Repository {
	return chat_repository.NewRepository()
}

func provideChatUsecase(repo *chat_repository.Repository) chat.Usecase {
	return chat.NewUsecase(chat.Deps{
		ChatRepository: repo,
		UserIDProvider: auth_context.MyUserIDProvider{},
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
		GRPCPort:               address,
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
