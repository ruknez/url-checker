package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	entity "url-checker/internal/domain"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/checker_mock.go . checker:CheckerMock
type checker interface {
	GetStatus(ctx context.Context, url string) (entity.Status, error)
}

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/logger_mock.go . logger:LoggerMock
type logger interface {
	Error(ctx context.Context, args ...interface{})
}

type HttpServer struct {
	ctx            context.Context
	checkerService checker
	logger         logger
}

func NewHttpServer(ctx context.Context, checker checker, logger logger) *HttpServer {
	return &HttpServer{
		ctx:            ctx,
		checkerService: checker,
		logger:         logger,
	}
}

func (h *HttpServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("LogHandle.io.ReadAll: %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	in := InResource{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("GetHandler json.Unmarshal: %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	out := OutStatus{}
	out.Status, err = h.checkerService.GetStatus(r.Context(), in.Url)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("GetHandler GetStatus: %w", err).Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	res, jError := json.Marshal(&out)
	if jError != nil {
		h.logger.Error(h.ctx, fmt.Errorf("GetHandler json.Marshal: %w", err).Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(res)
}

func (h *HttpServer) SetHandler() {

}
