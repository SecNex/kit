package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/api"
	"github.com/secnex/kit/server/handler/system"
	"github.com/secnex/kit/server/middlewares"
)

type Server struct {
	ApiServer  *api.ApiServer
	DB         *database.DatabaseConnection
	Middleware *middlewares.Middleware
}

func NewServer(port int, db *database.DatabaseConnection, middleware *middlewares.Middleware) *Server {
	apiServer := api.NewApiServer(port, db)

	return &Server{
		ApiServer:  apiServer,
		DB:         db,
		Middleware: middleware,
	}
}

func (s *Server) Run() {
	s.Migrate()

	s.NewSystemRouter("/_")

	s.ApiServer.Router.Use(s.Middleware.LoggingMiddleware)
	port := s.ApiServer.Config.Port

	fmt.Printf("🚀 Starting server on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.ApiServer.Router)
	if err != nil {
		fmt.Printf("🚨 Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func (s *Server) Migrate() {
	s.DB.AutoMigrateAll()
}

func (s *Server) NewRouter(path string) *mux.Router {
	return s.ApiServer.Router.PathPrefix(path).Subrouter()
}

func (s *Server) NewSystemRouter(path string) *mux.Router {
	router := s.NewRouter(path)

	router.HandleFunc("/healthz", system.HealthzHandler)
	router.HandleFunc("/hello", system.HelloWorldHandler)

	return router
}

func (s *Server) NewVersionRouter(version int, routes ...func(router *mux.Router)) *mux.Router {
	prefix := fmt.Sprintf("/api/v%d", version)
	fmt.Printf("🔍 New version router: %s\n", prefix)
	router := s.NewRouter(prefix)

	for _, route := range routes {
		route(router)
	}

	return router
}
