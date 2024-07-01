package app

import (
	"time"

	"url-checker/internal/api/http"
	main_http_server "url-checker/internal/app/http"
	"url-checker/internal/dependensis/http/check_client"
	"url-checker/internal/repository/in_memory_bd"
	"url-checker/internal/service/checker"
	"url-checker/pkg/logger"

	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				checker.NewChecker,
				fx.ParamTags(`name:"tickDuration"`),
			),
			fx.Annotate(
				func() time.Duration {
					return 1 * time.Second
				},
				fx.ResultTags(`name:"tickDuration"`),
			),
			http.NewHttpServer,
			fx.Annotate(
				check_client.NewCheckClient,
				fx.As(new(checker.GetUrlStatuser)),
			),
			fx.Annotate(
				in_memory_bd.NewCache,
				fx.As(new(checker.UrlRepository)),
			),
			fx.Annotate(
				logger.NewLogger,
				fx.As(new(checker.Logger)),
			),
			fx.Annotate(
				main_http_server.NewHttpService,
				fx.As(new(http.Server)),
				fx.ParamTags(`name:"host"`, `name:"port"`),
			),
			func() string {
				return "localhost"
			},
			func() int {
				return 8080
			},
		),
		fx.Invoke(func(*checker.Checker) {}),
	).Run()

}
