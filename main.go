package main

import (
	"fmt"

	"github.com/secnex/kit/config"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/api"
	"github.com/secnex/kit/server/handler"
	"github.com/secnex/kit/server/middlewares"
	"github.com/secnex/kit/utils/initializer"
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
		config.DefaultConfig{
			OrganizationName: "Default Organization",
			TenantName:       "Default Tenant",
			DomainName:       "Default Domain",
		},
	)

	fmt.Printf("🔥 Environment: %s\n\n", config.Environment)

	db := database.NewDatabaseConnectionWithConfig(config.GetDatabaseConfig())

	i := initializer.NewInitializer(db, config)
	i.Initialize()

	// Initialize handlers
	h := handler.NewHandler(db)

	server := api.NewServer(config.Port, db)

	internalRouter := server.CreateSubRouter("/_")

	// Internal
	internalRouter.HandleFunc("/healthz", h.Healthz).Methods("GET")

	publicRouter := server.CreateSubRouter("/public")

	// Public
	publicRouter.HandleFunc("/hello", h.Hello).Methods("GET")
	publicRouter.HandleFunc("/ip", h.IP).Methods("GET")

	apiV1Router := server.CreateApiServerWithMiddlewares("/api", 1, middlewares.ContentTypeOnlyJSON)

	// v1: Auth
	apiV1Router.HandleFunc("/auth/login", h.AuthLogin).Methods("POST")
	apiV1Router.HandleFunc("/auth/register", h.AuthRegister).Methods("POST")

	// v1: User
	apiV1Router.HandleFunc("/user", h.UserGet).Methods("GET")
	apiV1Router.HandleFunc("/user", h.UserNew).Methods("POST")

	// v1: Organization
	apiV1Router.HandleFunc("/organization", h.OrganizationGet).Methods("GET")
	apiV1Router.HandleFunc("/organization", h.OrganizationNew).Methods("POST")

	// v1: Domain
	apiV1Router.HandleFunc("/domain", h.DomainGet).Methods("GET")
	apiV1Router.HandleFunc("/domain", h.DomainNew).Methods("POST")

	// v1: Tenant
	apiV1Router.HandleFunc("/tenant", h.TenantGet).Methods("GET")
	apiV1Router.HandleFunc("/tenant", h.TenantNew).Methods("POST")

	// v1: Client
	apiV1Router.HandleFunc("/client", h.ClientGet).Methods("GET")
	apiV1Router.HandleFunc("/client", h.ClientNew).Methods("POST")

	// v1: App
	apiV1Router.HandleFunc("/app", h.AppGet).Methods("GET")
	apiV1Router.HandleFunc("/app", h.AppNew).Methods("POST")

	// v1: HTTPLog
	apiV1Router.HandleFunc("/log/http", h.HTTPLogGet).Methods("GET")

	server.Run()
}
