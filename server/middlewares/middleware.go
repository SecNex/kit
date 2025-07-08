package middlewares

import "github.com/secnex/kit/database"

type Middleware struct {
	DB      *database.DatabaseConnection
	LogFile string
}

func NewMiddleware(db *database.DatabaseConnection, logFile string) *Middleware {
	return &Middleware{DB: db, LogFile: logFile}
}
