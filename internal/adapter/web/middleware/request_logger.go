package middleware

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"net/http"
)

func RequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.INFO.Printf("[%s] [%s] %s\n", r.Method, r.RequestURI, r.RemoteAddr)
			handler.ServeHTTP(w, r)
		},
	)
}
