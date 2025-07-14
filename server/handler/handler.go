package handler

import "github.com/secnex/kit/database"

type Handler struct {
	db *database.DatabaseConnection
}

func NewHandler(db *database.DatabaseConnection) *Handler {
	return &Handler{db: db}
}
