package checker

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	entity "url-checker/internal/domain"
)

// checkAllUrls Ассинхронно проверяет все урлы из базы.
func (c *Checker) checkAllUrls(ctx context.Context) {
	wg := &sync.WaitGroup{}

	for _, urls := range c.urlRepo.GetAllUrls(ctx) {
		u := urls
		wg.Add(1)
		go func() {
			defer wg.Done()
			status, err := c.checkUrl(ctx, u)
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "c.checkUrl: "+u))
				return
			}

			err = c.saveStatus(ctx, u, status)
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "c.saveStatus: "+u))
				return
			}
		}()
	}

	wg.Wait()
}

func (c *Checker) checkUrl(ctx context.Context, url string) (entity.Status, error) {
	status, err := c.statuserService.GetUrlStatus(ctx, url)
	if err != nil {
		return status, errors.Wrap(err, "statuserService.GetUrlStatus")
	}

	return status, nil
}

func (c *Checker) saveStatus(ctx context.Context, url string, status entity.Status) error {
	return c.urlRepo.UpdateStatus(ctx, url, status)
}
