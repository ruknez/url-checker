package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api_http "url-checker/internal/api/http"
	"url-checker/internal/api/validation"
	main_http_server "url-checker/internal/app/http"
	"url-checker/internal/profiling"
)

type App struct {
	ctx             context.Context
	serviceProvider *serviceProvider
	mainServer      *main_http_server.Server
}

func NewApp(ctx context.Context) *App {
	a := App{
		ctx:             ctx,
		serviceProvider: createServiceProvider(ctx),
	}

	return &a
}

func (a *App) Run() error {
	chError := make(chan error)

	go a.runValidator(chError)
	if err := <-chError; err != nil {
		return fmt.Errorf("runValidator %w", err)
	}

	profiling.GoPProfStart(3451)
	return a.runMainHttpServer()
}

func (a *App) runMainHttpServer() error {
	mux := http.NewServeMux()

	httpClientServer := api_http.NewHttpServer(a.ctx, a.serviceProvider.checker, a.serviceProvider.logger)

	mux.HandleFunc("/getStatus", httpClientServer.GetHandler)

	httpServer := main_http_server.NewHttpService(a.ctx, "", 8080, mux)

	if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("httpServer.ListenAndServe %w", err)
	}

	return nil
}

func (a *App) runValidator(chError chan<- error) {
	mux := http.NewServeMux()

	pinger := validation.NewPingHandler()

	mux.HandleFunc("/ping", pinger.PingHandler)

	httpServer := main_http_server.NewHttpService(a.ctx, "", 8085, mux)

	if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		chError <- fmt.Errorf("httpServer.ListenAndServe %w", err)
	}

	chError <- nil
}
