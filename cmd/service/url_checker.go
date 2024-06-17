package main

import (
	"context"
	"log"

	"url-checker/internal/app"
	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . checker:CheckerMock
type checker interface {
	SaveToCheck(ctx context.Context, urlInfo entity.UrlInfo) error
	GetStatus(ctx context.Context, url string) (entity.Status, error)
}

func main() {
	ctx := context.Background()

	app := app.NewApp(ctx)

	err := app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
