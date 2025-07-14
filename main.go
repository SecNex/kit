package main

import (
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/api"
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

	h := handler.NewHandler(db)

	server := api.NewServer(8081, db)

	internalRouter := server.CreateSubRouter("/_")

	internalRouter.HandleFunc("/healthz", h.Healthz).Methods("GET")

	internalRouter.HandleFunc("/hello", h.Hello).Methods("GET")

	apiV1Router := server.CreateSubRouterWithMiddlewares("/api/v1", middlewares.OnlyJSON)

	apiV1Router.HandleFunc("/ip", h.IP).Methods("GET")
	apiV1Router.HandleFunc("/auth/login", h.AuthLogin).Methods("POST")
	apiV1Router.HandleFunc("/user", h.UserGet).Methods("GET")
	apiV1Router.HandleFunc("/user", h.UserNew).Methods("POST")

	server.Run()
}
