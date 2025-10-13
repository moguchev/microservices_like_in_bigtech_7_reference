package main

import (
	"context"
	"log"

	"social/cmd/wire"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv, err := wire.InitializeServer(ctx)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("server run: %v", err)
	}
}
