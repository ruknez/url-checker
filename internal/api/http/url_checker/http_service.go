package http

import (
	"context"
	"log/slog"
	"net/http"

	entity "url-checker/internal/domain"

	"go.uber.org/fx"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . Checker:CheckerMock
type Checker interface {
	GetStatus(ctx context.Context, url string) (entity.Status, error)
	AddUrl(ctx context.Context, url string) error
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/logger_mock.go . Logger:LoggerMock
type Logger interface {
	Error(ctx context.Context, args ...interface{})
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/server_mock.go . Server:ServerMock
type Server interface {
	RegisterHandlers(mux *http.ServeMux)
}

type HttpServer struct {
	ctx            context.Context
	checkerService Checker
	logger         Logger
}

func NewHttpServer(lc fx.Lifecycle, checker Checker, logger Logger, server Server) *HttpServer {
	h := &HttpServer{
		checkerService: checker,
		logger:         logger,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				h.ctx = ctx
				mux := http.NewServeMux()

				mux.HandleFunc("GET /getStatus", h.GetHandler)
				mux.HandleFunc("/addUrl", h.AddUrlHandler)
				server.RegisterHandlers(mux)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				slog.Info("stopping http server url_checker")
				return nil
			},
		})

	return h
}
