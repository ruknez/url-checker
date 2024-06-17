package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	l.log.ErrorContext(ctx, "ERROR", args...)
}
