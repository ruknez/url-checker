package checker

import (
	"context"
	"time"

	entity "url-checker/internal/domain"

	"go.uber.org/fx"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/url_repository_mock.go . UrlRepository:UrlRepositoryMock
type UrlRepository interface {
	Get(ctx context.Context, url string) (entity.UrlInfo, error)
	GetAllUrls(ctx context.Context) []string
	GetAllUrlsToCheck(ctx context.Context) []string
	UpdateStatus(ctx context.Context, url string, status entity.Status) error
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/logger_mock.go . Logger:LoggerMock
type Logger interface {
	Error(ctx context.Context, args ...interface{})
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/get_url_statuser_mock.go . GetUrlStatuser:GetUrlStatuserMock
type GetUrlStatuser interface {
	GetUrlStatus(ctx context.Context, url string) (entity.Status, error)
}

type Checker struct {
	urlRepo         UrlRepository
	tickDuration    time.Duration
	logger          Logger
	statuserService GetUrlStatuser
}

func NewChecker(
	lc fx.Lifecycle,
	urlRepo UrlRepository,
	logger Logger,
	tickDuration time.Duration,
	statuserService GetUrlStatuser,
) *Checker {
	ch := &Checker{
		urlRepo:         urlRepo,
		logger:          logger,
		tickDuration:    tickDuration,
		statuserService: statuserService,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ch.gRun(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return ch
}

// gRun запускает итерационный процесс проверки урлов из базы.
func (c *Checker) gRun(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(c.tickDuration)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case <-ctx.Done():
				return
			default:
				c.checkAllUrls(ctx)
				ticker.Reset(c.tickDuration)
			}
		}
	}()
}
