package repository

import (
	"context"
	"fmt"

	"url-checker/internal/repository/entity"
)

//go:generate ./../../../../bin/moq -stub -skip-ensure -pkg mocks -out ./mocks/url_repository_mock.go . UrlRepository:UrlRepositoryMock
type UrlRepository interface {
	Get(ctx context.Context, url string) (*entity.UrlInfo, error)
	GetAllUrls(ctx context.Context) []string
	GetAllUrlsToCheck(ctx context.Context) []string
	UpdateStatus(ctx context.Context, url string, status entity.Status) error
}

//go:generate ./../../../../bin/moq -stub -skip-ensure -pkg mocks -out ./mocks/url_repository_bd_mock.go . urlRepositoryBd:UrlRepositoryBdMock
type urlRepositoryBd interface {
	Get(ctx context.Context, url string) (entity.UrlInBd, error)
	GetAllUrls(ctx context.Context) []string
	GetAllUrlsToCheck(ctx context.Context) []string
	UpdateStatus(ctx context.Context, url string, status int) error
}

type urlRepositoryImplement struct {
	bd urlRepositoryBd
}

func NewUrlRepositoryImplement(bd urlRepositoryBd) *urlRepositoryImplement {
	return &urlRepositoryImplement{
		bd: bd,
	}
}

func GetLOLO() {
	fmt.Println("LOLO")
}
