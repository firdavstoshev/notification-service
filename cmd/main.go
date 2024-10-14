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

	"notification-service/domain"
	"notification-service/handler"
	"notification-service/server"
	"notification-service/storage"
	"notification-service/worker"
)

func main() {
	storage := storage.New()
	events := make(chan domain.Event)

	handler := handler.New(storage, events)
	server := server.New(handler)

	go worker.ProcessEvents(events)

	go func() {
		log.Println("Starting server on", server.HttpServer.Addr)
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server run failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	close(events)

	log.Println("Server exiting")
}
