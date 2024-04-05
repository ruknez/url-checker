package logger

import (
	"context"
)

//go:generate ./../../../../bin/moq -stub -skip-ensure -pkg mocks -out ./mocks/contextual_logger_mock.go . ContextualLogger:ContextualLoggerMock
type ContextualLogger interface {
	Info(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	WithError(ctx context.Context, err error) context.Context
	WithFields(ctx context.Context, fields map[string]interface{}) context.Context
	WithField(ctx context.Context, k string, v interface{}) context.Context
}
