package handler

import (
	"encoding/json"
	"net/http"

	"github.com/secnex/kit/database"
)

type Handler struct {
	db *database.DatabaseConnection
}

func NewHandler(db *database.DatabaseConnection) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RequiredBodyFields(w http.ResponseWriter, r *http.Request, fields ...string) bool {
	var result map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		BadRequest(w, r, "Invalid request body")
		return false
	}

	for _, field := range fields {
		if result[field] == nil {
			BadRequest(w, r, "Missing required field: "+field)
			return false
		}
	}

	return true
}
