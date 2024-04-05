package repository

import (
	"context"
)

func (u *urlRepositoryImplement) GetAllUrlsToCheck(ctx context.Context) []string {
	return u.bd.GetAllUrlsToCheck(ctx)
}
