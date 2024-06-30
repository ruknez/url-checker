package http

import (
	"context"

	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . checker:CheckerMock
type checker interface {
	GetStatus(ctx context.Context, url string) (entity.Status, error)
	AddUrl(ctx context.Context, url string) error
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/logger_mock.go . logger:LoggerMock
type logger interface {
	Error(ctx context.Context, args ...interface{})
}

type HttpServer struct {
	ctx            context.Context
	checkerService checker
	logger         logger
}

func NewHttpServer(ctx context.Context, checker checker, logger logger) *HttpServer {
	return &HttpServer{
		ctx:            ctx,
		checkerService: checker,
		logger:         logger,
	}
}
