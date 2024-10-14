package server

import (
	"context"
	"net/http"
	"time"

	"notification-service/handler"
)

type Server struct {
	HttpServer *http.Server
}

func (s *Server) Run() error {
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func New(handler *handler.Handler) *Server {
	return &Server{HttpServer: &http.Server{
		Addr:         ":" + "8080",
		Handler:      handler.InitRoutes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}}
}
