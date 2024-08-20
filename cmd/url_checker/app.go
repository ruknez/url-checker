package main

import (
	"context"
	"log/slog"
	"url-checker/pkg/logger"
)

func main() {

	//l := slog.New(logger.NewLogger())

	slog.SetDefault(logger.NewLoggerMain())

	slog.Error("LOLO")
	slog.Error("LOLO2")

	ctx := context.Background()

	ctx = context.WithValue(ctx, "1234", "lolo data")

	slog.ErrorContext(ctx, "with data")

	// transaction out box (river)
	/*	child := slog.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
		),
	)*/

	//app.Run()
}
