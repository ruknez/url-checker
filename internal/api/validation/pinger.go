package validation

import (
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
)

type PingHandlerSt struct{}

func NewPingHandler() *PingHandlerSt {
	return &PingHandlerSt{}
}

func (p *PingHandlerSt) PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("X-Test-ping", "ok")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("OK logs_to_ch"))
	if err != nil {
		slog.Error(errors.Wrap(err, "PingHandler w.Write").Error())
	}
}
