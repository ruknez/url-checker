package main_http_server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	http.Server
}

func NewHttpService(ctx context.Context, host string, port int, handlers http.Handler) *Server {
	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: handlers,
	}

	slog.Info("server start on", "addr", httpServer.Addr)

	go func() {
		<-ctx.Done()
		slog.Info("http server Shutdown")
		httpServer.Shutdown(ctx)
	}()

	return &Server{
		httpServer,
	}
}
