package repository

import (
	"context"

	"url-checker/internal/repository/entity"
)

func (u *urlRepositoryImplement) UpdateStatus(ctx context.Context, url string, status entity.Status) error {
	return u.bd.UpdateStatus(ctx, url, int(status))
}
