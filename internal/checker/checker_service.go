package checker

import (
	"context"
	"time"

	"url-checker/internal/check_client"
	"url-checker/internal/repository"
	"url-checker/internal/repository/entity"
	"url-checker/pkg/logger"
)

//go:generate ./../../../../bin/moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . Checker:CheckerMock
type Checker interface {
	SaveToCheck(ctx context.Context, urlInfo entity.UrlInfo) error
	GetStatus(ctx context.Context, url string) (entity.Status, error)
}

type checker struct {
	urlRepo         repository.UrlRepository
	tickDuration    time.Duration
	logger          logger.ContextualLogger
	statuserService check_client.GetUrlStatuser
}

func NewChecker(
	ctx context.Context,
	urlRepo repository.UrlRepository,
	logger logger.ContextualLogger,
	tickDuration time.Duration,
	statuserService check_client.GetUrlStatuser,
) *checker {
	ch := &checker{
		urlRepo:         urlRepo,
		logger:          logger,
		tickDuration:    tickDuration,
		statuserService: statuserService,
	}

	ch.gRun(ctx)

	return ch
}

func (c *checker) gRun(ctx context.Context) {
	go func() {
		for range time.Tick(c.tickDuration) {
			select {
			case <-ctx.Done():
				return
			default:
				c.checkAllUrls(ctx)
			}
		}
	}()
}
