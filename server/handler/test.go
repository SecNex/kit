package handler

import "net/http"

func (h *Handler) TestDatabaseConnection(w http.ResponseWriter, r *http.Request) {
	if err := h.DB.TestConnection(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
