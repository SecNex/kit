package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/secnex/kit/models"
)

type HTTPLogGetResponse struct {
	Count int              `json:"count"`
	Logs  []models.HTTPLog `json:"logs"`
}

func (h *Handler) HTTPLogGet(w http.ResponseWriter, r *http.Request) {
	var result []models.HTTPLog

	countStr := r.URL.Query().Get("count")
	count := 50

	if countStr != "" {
		if parsedCount, err := strconv.Atoi(countStr); err == nil && parsedCount > 0 {
			count = parsedCount
		}
	}

	err := h.db.DB.Order("created_at DESC").Limit(count).Find(&result).Error
	if err != nil {
		InternalServerError(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HTTPLogGetResponse{
		Count: len(result),
		Logs:  result,
	})
}
