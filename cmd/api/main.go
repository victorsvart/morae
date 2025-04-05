// Package main is the entry point of the Morae application.
// It starts the HTTP server and handles graceful shutdown.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app, cfg := mount()
	go func() {
		if err := app.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server close error: %v", err)
		}
	}()

	log.Printf("Server started at: %s%s", cfg.Host, app.Addr)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	s := <-sig
	log.Printf("Received signal: %s. Initiating shutdown...", s)

	const shutdownTimeout = 10 * time.Second
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownRelease()

	if err := app.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped. Bye bye!")
}
