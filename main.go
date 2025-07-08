package main

import (
	"github.com/gorilla/mux"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server"
	"github.com/secnex/kit/server/handler"
	"github.com/secnex/kit/server/middlewares"
)

func main() {
	db := database.NewDatabaseConnectionWithConfig(database.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "kit",
	})

	middleware := middlewares.NewMiddleware(db, "log.txt")

	server := server.NewServer(8080, db, middleware)

	handler := handler.NewHandler(db)

	server.NewVersionRouter(1, func(router *mux.Router) {
		router.HandleFunc("/test", handler.TestDatabaseConnection)
	})

	server.Run()
}
