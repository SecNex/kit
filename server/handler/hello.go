package handler

import (
	"fmt"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, "Hello, World!")))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}
}
