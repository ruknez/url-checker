package checker

import (
	"context"
	"time"

	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/url_repository_mock.go . urlRepository:UrlRepositoryMock
type urlRepository interface {
	Get(ctx context.Context, url string) (*entity.UrlInfo, error)
	GetAllUrls(ctx context.Context) []string
	GetAllUrlsToCheck(ctx context.Context) []string
	UpdateStatus(ctx context.Context, url string, status entity.Status) error
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/logger_mock.go . logger:LoggerMock
type logger interface {
	Error(ctx context.Context, args ...interface{})
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/get_url_statuser_mock.go . getUrlStatuser:GetUrlStatuserMock
type getUrlStatuser interface {
	GetUrlStatus(ctx context.Context, url string) (entity.Status, error)
}

type checker struct {
	urlRepo         urlRepository
	tickDuration    time.Duration
	logger          logger
	statuserService getUrlStatuser
}

func NewChecker(
	ctx context.Context,
	urlRepo urlRepository,
	logger logger,
	tickDuration time.Duration,
	statuserService getUrlStatuser,
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
