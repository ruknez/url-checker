package main_http_server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/fx"
)

type MainServiceConfigInterface interface {
	MainServiceConfig() (string, int)
}

type PingServiceConfigInterface interface {
	PingServiceConfig() (string, int)
}

type ServerTransport struct {
	*http.Server
}

// TODO очень сомнительное решение надо бы обсутить как сделать лучше.
// Мб передавалть host and  port как аргументы и как-то рулить на уровне di.
func NewPingHttpService(lc fx.Lifecycle, conf PingServiceConfigInterface) *ServerTransport {
	slog.Info("NewPingHttpService constructor")
	host, port := conf.PingServiceConfig()
	return httpService(lc, host, port)
}

func NewMainHttpService(lc fx.Lifecycle, conf MainServiceConfigInterface) *ServerTransport {
	host, port := conf.MainServiceConfig()
	return httpService(lc, host, port)
}

func httpService(lc fx.Lifecycle, host string, port int) *ServerTransport {
	errCh := make(chan error)
	httpServer := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				errCh <- httpServer.ListenAndServe()
			}()

			slog.Info("server start on", "addr", httpServer.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server at", httpServer.Addr)

			if err := httpServer.Shutdown(ctx); err != nil {
				return errors.Wrap(err, fmt.Sprintf("http server %s shutdown failed", httpServer.Addr))
			}
			return <-errCh
		},
	})

	return &ServerTransport{
		httpServer,
	}
}

func (s *ServerTransport) RegisterHandlers(mux *http.ServeMux) {
	s.Handler = mux
}
