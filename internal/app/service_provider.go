package app

import (
	"context"
	"time"

	"url-checker/internal/dependensis/http/check_client"
	"url-checker/internal/repository/in_memory_bd"
	"url-checker/internal/service/checker"
	"url-checker/pkg/logger"
)

type serviceProvider struct {
	repositoryInMemory *in_memory_bd.Cache
	checker            *checker.Checker
	logger             *logger.Logger
	checkClient        *check_client.CheckClient
}

func createServiceProvider(ctx context.Context) *serviceProvider {
	sp := &serviceProvider{}

	sp.run(ctx)
	return sp
}

func (s *serviceProvider) run(ctx context.Context) {
	if s.checker == nil {
		s.checker = s.createChecker(ctx)
	}
}

func (s *serviceProvider) createChecker(ctx context.Context) *checker.Checker {
	if s.repositoryInMemory == nil {
		s.repositoryInMemory = s.createRepositoryInMemory(ctx)
	}

	if s.logger == nil {
		s.logger = logger.NewLogger()
	}

	if s.checkClient == nil {
		s.checkClient = s.createCheckClient()
	}

	return checker.NewChecker(ctx, s.repositoryInMemory, s.logger, time.Second*1, s.checkClient)
}

func (s *serviceProvider) createRepositoryInMemory(_ context.Context) *in_memory_bd.Cache {
	return in_memory_bd.NewCache()
}

func (s *serviceProvider) createCheckClient() *check_client.CheckClient {
	return check_client.NewCheckClient()
}
