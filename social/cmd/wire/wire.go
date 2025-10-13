//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	auth_context "lib/auth/context"
	lib_grpc_middleware "lib/middleware/grpc"
	social_controller "social/internal/app/controllers/social"
	social_repository "social/internal/app/repositories/social"
	"social/internal/app/server"
	"social/internal/app/usecases/social"
	grpc_middleware "social/internal/middleware/grpc"
	socialv1 "social/pkg/api/social/v1"

	"github.com/bufbuild/protovalidate-go"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

const address = ":8080"

func InitializeServer(ctx context.Context) (*server.Server, error) {
	wire.Build(
		// ========== Repositories ==========
		social_repository.NewRepository,
		wire.Bind(new(social.SocialRepository), new(*social_repository.Repository)),

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
	)

	return &server.Server{}, nil
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
	return server.Config{
		GRPCPort:               address,
		ChainUnaryInterceptors: mws,
	}
}

func newUserIDProvider() social.UserIDProvider {
	return auth_context.MyUserIDProvider{}
}
