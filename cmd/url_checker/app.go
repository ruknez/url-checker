package main

import (
	"context"
	"log"

	"url-checker/internal/app"
)

func main() {
	ctx := context.Background()

	app := app.NewApp(ctx)

	err := app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
