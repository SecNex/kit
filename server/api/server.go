package api

import (
	"github.com/gorilla/mux"
	"github.com/secnex/kit/database"
)

type ApiServerConfig struct {
	Port int
}

type ApiServer struct {
	Config ApiServerConfig
	Router *mux.Router
}

func NewApiServer(port int, db *database.DatabaseConnection) *ApiServer {
	router := mux.NewRouter()

	return &ApiServer{
		Config: ApiServerConfig{
			Port: port,
		},
		Router: router,
	}
}
