package checker

import (
	"context"

	"github.com/pkg/errors"
	entity "url-checker/internal/domain"
)

func (c *Checker) GetStatus(ctx context.Context, url string) (entity.Status, error) {
	urlInfo, err := c.urlRepo.Get(ctx, url)
	if err != nil {
		return entity.NotCheck, errors.Wrap(err, "urlRepo.Get")
	}

	return urlInfo.Status, nil
}
