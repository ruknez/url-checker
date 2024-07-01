package app

import (
	"time"

	"url-checker/internal/api/http/ping"
	"url-checker/internal/api/http/url_checker"
	main_http_server "url-checker/internal/app/http"
	"url-checker/internal/config"
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
				fx.As(new(http.Checker)),
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
				fx.As(new(http.Logger)),
			),
			fx.Annotate(
				main_http_server.NewMainHttpService,
				fx.As(new(http.Server)),
			),
			fx.Annotate(
				config.NewConfigService,
				fx.As(new(main_http_server.MainServiceConfigInterface)),
				fx.As(new(main_http_server.PingServiceConfigInterface)),
			),
			fx.Annotate(
				main_http_server.NewPingHttpService,
				fx.As(new(ping.PingerTransport)),
			),

			ping.NewPingHandler,
		),
		fx.Invoke(func(*ping.PingHandlerSt) {}),
	).Run()

}
