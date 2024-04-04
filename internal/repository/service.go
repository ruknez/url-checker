package repository

import (
	"context"
	"fmt"

	"url-checker/internal/repository/entity"
)

type UrlRepository interface {
	Get(ctx context.Context, url string) (*entity.UrlInfo, error)
}

type urlRepositoryBd interface {
	Get(ctx context.Context, url string) (entity.UrlInBd, error)
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
