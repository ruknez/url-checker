package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *HttpServer) AddUrlHandler(w http.ResponseWriter, r *http.Request) {
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

	err = h.checkerService.AddUrl(r.Context(), in.Url)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("AddUrlHandler AddUrl: %w", err).Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
