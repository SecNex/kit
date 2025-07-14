package main

import (
	"github.com/secnex/kit/config"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/api"
	"github.com/secnex/kit/server/handler"
	"github.com/secnex/kit/server/middlewares"
)

func main() {
	config := config.NewConfig(
		database.DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DBName:   "kit",
		},
		"8081",
		"secret",
		true,
		"http.log",
	)

	db := database.NewDatabaseConnectionWithConfig(config.GetDatabaseConfig())

	h := handler.NewHandler(db)

	server := api.NewServer(config.Port, db)

	internalRouter := server.CreateSubRouter("/_")

	// Internal
	internalRouter.HandleFunc("/healthz", h.Healthz).Methods("GET")

	publicRouter := server.CreateSubRouter("/public")

	// Public
	publicRouter.HandleFunc("/hello", h.Hello).Methods("GET")
	publicRouter.HandleFunc("/ip", h.IP).Methods("GET")

	apiV1Router := server.CreateSubRouterWithMiddlewares("/api/v1", middlewares.ContentTypeOnlyJSON)

	// API v1
	// Auth
	apiV1Router.HandleFunc("/auth/login", h.AuthLogin).Methods("POST")
	apiV1Router.HandleFunc("/auth/register", h.AuthRegister).Methods("POST")

	// User
	apiV1Router.HandleFunc("/user", h.UserGet).Methods("GET")
	apiV1Router.HandleFunc("/user", h.UserNew).Methods("POST")

	// Organization
	apiV1Router.HandleFunc("/organization", h.OrganizationGet).Methods("GET")
	apiV1Router.HandleFunc("/organization", h.OrganizationNew).Methods("POST")

	// Domain
	apiV1Router.HandleFunc("/domain", h.DomainGet).Methods("GET")
	apiV1Router.HandleFunc("/domain", h.DomainNew).Methods("POST")

	// Tenant
	apiV1Router.HandleFunc("/tenant", h.TenantGet).Methods("GET")
	apiV1Router.HandleFunc("/tenant", h.TenantNew).Methods("POST")

	// Client
	apiV1Router.HandleFunc("/client", h.ClientGet).Methods("GET")
	apiV1Router.HandleFunc("/client", h.ClientNew).Methods("POST")

	// App
	apiV1Router.HandleFunc("/app", h.AppGet).Methods("GET")
	apiV1Router.HandleFunc("/app", h.AppNew).Methods("POST")

	server.Run()
}
