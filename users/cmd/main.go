package main

import (
	"context"
	"log"

	auth_context "lib/auth/context"
	lib_grpc_middleware "lib/middleware/grpc"
	users_controller "users/internal/app/controllers/users"
	profiles_repository "users/internal/app/repositories/profiles"
	"users/internal/app/server"
	"users/internal/app/usecases/users"
	grpc_middleware "users/internal/middleware/grpc"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

const address = ":8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository
	profilesRepo := profiles_repository.NewRepository()

	// =========================
	// usecases
	// =========================

	usersUsecase := users.NewUsecase(users.Deps{
		ProfilesRepository: profilesRepo,
		UserIDProvider:     auth_context.MyUserIDProvider{},
	})

	// =========================
	// delivery
	// =========================

	// controllers
	usersController := users_controller.New(users_controller.Deps{
		UsersUsecase: usersUsecase,
	})

	// middlewares
	validator, err := protovalidate.New(protovalidate.WithDisableLazy(false))
	if err != nil {
		log.Fatalf("server: failed to initialize validator: %s", err)
	}
	mws := []grpc.UnaryServerInterceptor{
		grpc_middleware.ErrorsUnaryServerInterceptor(),
		lib_grpc_middleware.ValidateUnaryServerInterceptor(validator),
	}

	// infrastructure server
	config := server.Config{
		GRPCPort:               address,
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.New(ctx, config, server.Contollers{
		UserServiceServer: usersController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
