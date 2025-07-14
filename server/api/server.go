package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/secnex/kit/database"
	"github.com/secnex/kit/server/logger"
)

type Server struct {
	Database *database.DatabaseConnection
	Router   *mux.Router
	Port     int
}

func NewServer(port int, database *database.DatabaseConnection) *Server {
	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))

	router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	))

	return &Server{
		Database: database,
		Router:   router,
		Port:     port,
	}
}

func (s *Server) CreateSubRouter(path string) *mux.Router {
	fmt.Println("🚀 Creating sub router for path: ", path)
	subRouter := s.Router.PathPrefix(path).Subrouter()

	return subRouter
}

func (s *Server) CreateSubRouterWithMiddleware(path string, middleware func(http.Handler) http.Handler) *mux.Router {
	subRouter := s.CreateSubRouter(path)

	fmt.Println("🚀 Adding middleware to sub router for path: ", path)
	subRouter.Use(middleware)

	return subRouter
}

func (s *Server) CreateSubRouterWithMiddlewares(path string, middlewares ...func(http.Handler) http.Handler) *mux.Router {
	subRouter := s.CreateSubRouter(path)
	for _, middleware := range middlewares {
		subRouter.Use(middleware)
	}
	return subRouter
}

func (s *Server) Run() {
	httpLogger := logger.NewHTTPLogger(s.Database, "http.log")
	fmt.Println("🚀 Running server on port: ", s.Port)
	handler := httpLogger.LogHTTPRequest(s.Router)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.Port), handler)
}
