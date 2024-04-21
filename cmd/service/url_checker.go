package main

import (
	"context"

	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . checker:CheckerMock
type checker interface {
	SaveToCheck(ctx context.Context, urlInfo entity.UrlInfo) error
	GetStatus(ctx context.Context, url string) (entity.Status, error)
}

func main() {}
