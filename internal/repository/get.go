package repository

import (
	"context"

	"github.com/pkg/errors"

	"url-checker/internal/repository/entity"
)

func (u *urlRepositoryImplement) Get(ctx context.Context, url string) (*entity.UrlInfo, error) {

	res, err := u.bd.Get(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "bd.Get")
	}

	return &entity.UrlInfo{
		URL:      res.URL,
		Duration: res.Duration,
		Headers:  res.Headers,
	}, nil
}
