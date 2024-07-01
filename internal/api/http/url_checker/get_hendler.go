package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *HttpServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("LogHandle.io.ReadAll: %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	in := inResource{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		h.logger.Error(h.ctx, fmt.Errorf("GetHandler json.Unmarshal: %w", err).Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	out := outStatus{}
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
