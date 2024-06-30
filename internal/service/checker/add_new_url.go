package checker

import (
	"context"

	"github.com/pkg/errors"
	entity "url-checker/internal/domain"
)

func (c *Checker) AddUrl(ctx context.Context, url string) error {
	if _, err := c.urlRepo.Get(ctx, url); errors.Is(err, entity.NoDataErr) {
		return errors.Wrap(c.urlRepo.UpdateStatus(ctx, url, entity.NotCheck), "urlRepo.UpdateStatus")

	}

	return errors.New("AddUrl cannot add url")
}
