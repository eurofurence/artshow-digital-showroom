package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/server"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Handle exit if processed is killed
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv, err := server.NewServer(cfg, ctx)
	if err != nil {
		log.Fatalf("error setting up server: %v", err)
	}

	errCh := make(chan error, 1)

	go func() {
		errCh <- srv.Start()
	}()

	select {
	case <-ctx.Done():
		// Normal shutdown requested.
		if err := srv.Close(); err != nil {
			log.Printf("server shutdown error: %v", err)
		}

	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}
}
