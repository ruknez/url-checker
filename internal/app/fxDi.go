package app

import (
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
			checker.NewChecker,
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
