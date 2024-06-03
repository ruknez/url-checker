package main

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"

	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . checker:CheckerMock
type checker interface {
	SaveToCheck(ctx context.Context, urlInfo entity.UrlInfo) error
	GetStatus(ctx context.Context, url string) (entity.Status, error)
}

func main() {
	//myCheker := checker2.NewChecker()
	//myCheker.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
	debug.ReadBuildInfo()
}
