package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	auth_context "lib/auth/context"
	lib_grpc_middleware "lib/middleware/grpc"
	"lib/postgres"
	"lib/postgres/transaction_manager"
	users_controller "users/internal/app/controllers/users"
	profiles_repository "users/internal/app/repositories/profiles"
	"users/internal/app/server"
	"users/internal/app/usecases/users"
	grpc_middleware "users/internal/middleware/grpc"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

var (
	grpcListenPort = os.Getenv("GRPC_LISTEN_PORT")

	pgUser     = os.Getenv("USERS_POSTGRES_USER")
	pgPassword = os.Getenv("USERS_POSTGRES_PASSWORD")
	pgDB       = os.Getenv("USERS_POSTGRES_DB")
	pgHost     = os.Getenv("USERS_POSTGRES_HOST")
	pgPort     = os.Getenv("POSTGRES_PORT")
	pgSSLmode  = os.Getenv("POSTGRES_SSLMODE")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// infra
	// =========================

	pgConn, err := postgres.NewConnectionPool(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgUser, pgPassword, pgHost, pgPort, pgDB, pgSSLmode),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatalf("failed to create connection pool: %v", err)
	}

	// =========================
	// adapters
	// =========================

	txManager := transaction_manager.New(pgConn)

	// =========================
	// repository
	// =========================

	profilesRepo := profiles_repository.NewRepository(txManager)

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
		GRPCPort:               ":" + grpcListenPort,
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
