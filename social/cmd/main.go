package main

import (
	"context"
	"log"

	"social/cmd/wire"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := wire.InitializeApp(ctx)
	if err != nil {
		log.Fatalf("failed to init server: %v", err)
	}

	go app.Worker.Run(ctx)

	if err = app.Server.Run(ctx); err != nil {
		log.Fatalf("server run: %v", err)
	}
}
