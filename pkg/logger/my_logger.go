package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const LOGGER_CONFIG = "LOGGER_CONFIG"

type Logger struct {
	log   *slog.Logger
	level *slog.Level
}

// json  -> yaml -> env

func NewLogger() *Logger {
	return &Logger{
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

}

func NewLoggerMain() *slog.Logger {
	return slog.New(&Logger{
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	})
}

func (l *Logger) Error(ctx context.Context, args ...interface{}) {
	l.log.ErrorContext(ctx, "ERROR", args...)
}

type Handler interface {
	Enabled(context.Context, slog.Level) bool
	Handle(context.Context, slog.Record) error
	WithAttrs(attrs []slog.Attr) Handler
	WithGroup(name string) Handler
}

func SomeFuncMain() {
	l := slog.New(NewLogger())

	l.Log(context.Background(), slog.LevelDebug, "SomeFuncMain")

}

func (l *Logger) Enabled(ctx context.Context, level slog.Level) bool {
	fmt.Println("level", level.String())
	return true
}

func (l *Logger) Handle(ctx context.Context, message slog.Record) error {
	fmt.Printf("%#v", message)

	ctxData := ctx.Value("123")

	fmt.Println("level", message.Level, " ctxData ", ctxData)
	fmt.Println("message", message.Message)

	return nil
}

func (l *Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	fmt.Println("WithAttrs")
	return l
}
func (l *Logger) WithGroup(name string) slog.Handler {
	fmt.Println("WithGroup")
	return l
}
