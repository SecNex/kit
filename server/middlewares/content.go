package middlewares

import (
	"net/http"

	"github.com/secnex/kit/server/handler"
)

func ContentTypeOnlyJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "DELETE" || r.Method == "OPTIONS" || r.Method == "HEAD" {
			next.ServeHTTP(w, r)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			handler.WrongContentType(w, r, "application/json")
			return
		}
		next.ServeHTTP(w, r)
	})
}
