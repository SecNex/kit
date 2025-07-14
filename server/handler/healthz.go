package handler

import (
	"fmt"
	"net/http"
)

type HealthzResponse struct {
	Status string `json:"status"`
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	acceptJSON := r.Header.Get("Accept") == "application/json"

	databaseStatus := h.db.TestConnection()
	databaseStatusText := "OK"

	if databaseStatus != nil {
		databaseStatusText = "Not OK"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if acceptJSON {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"status":"%s"}`, databaseStatusText)))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
