package repository

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"

	"url-checker/internal/repository/entity"
)

func (u *urlRepositoryImplement) Get(ctx context.Context, url string) (*entity.UrlInfo, error) {
	res, err := u.bd.Get(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "bd.Get")
	}

	var tm *time.Time
	if res.LastCheck != 0 {
		tm = pointer.To(time.Unix(res.LastCheck, 0))
	}

	return &entity.UrlInfo{
		URL:       res.URL,
		Duration:  castToDuration(res.Duration),
		Headers:   res.Headers,
		LastCheck: tm,
		Status:    entity.Status(res.Status),
	}, nil
}

func (u *urlRepositoryImplement) GetAllUrls(ctx context.Context) []string {
	return u.bd.GetAllUrls(ctx)
}

func castToDuration(t int64) time.Duration {
	return time.Duration(t)
}

func castFromDuration(t time.Duration) int64 {
	return t.Microseconds()
}
