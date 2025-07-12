package main

import (
	"net/http"

	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/api"
)

func main() {
	db := database.NewDatabaseConnectionWithConfig(database.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "kit",
	})

	server := api.NewServer(8081, db)

	internalRouter := server.CreateSubRouter("/_")

	internalRouter.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	internalRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	apiV1Router := server.CreateSubRouter("/api/v1")

	apiV1Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	server.Run()
}
