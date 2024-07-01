package ping

import (
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
)

//go:generate moq -stub -skip-ensure -pkg mocks -out ./mocks/server_mock.go . Server:ServerMock
type PingerTransport interface {
	RegisterHandlers(mux *http.ServeMux)
}

type PingHandlerSt struct{}

func NewPingHandler(server PingerTransport) *PingHandlerSt {
	slog.Info("NewPingHandler constructor")
	p := &PingHandlerSt{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", p.PingHandler)
	server.RegisterHandlers(mux)

	return p
}

func (p *PingHandlerSt) PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("X-Test-ping", "ok")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("OK logs_to_ch"))
	if err != nil {
		slog.Error(errors.Wrap(err, "PingHandler w.Write").Error())
	}
}
