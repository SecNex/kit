package handler

import (
	"net/http"
)

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := h.db.TestConnection(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not OK"))
		return
	}

	w.Write([]byte("OK"))
}
