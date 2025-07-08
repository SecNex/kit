package handler

import "github.com/secnex/kit/database"

type Handler struct {
	DB *database.DatabaseConnection
}

func NewHandler(db *database.DatabaseConnection) *Handler {
	return &Handler{
		DB: db,
	}
}
