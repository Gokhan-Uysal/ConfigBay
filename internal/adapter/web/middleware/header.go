package middleware

import (
	"net/http"
)

func EnableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
			w.Header().Add("Access-Control-Allow-Headers", "origin")
			w.Header().Add("Access-Control-Allow-Headers", "accept")
			w.Header().Add("Access-Control-Allow-Headers", "authorization")

			w.Header().Set("Access-Control-Allow-Credentials", "true")

			w.Header().Add("Access-Control-Allow-Methods", http.MethodGet)
			w.Header().Add("Access-Control-Allow-Methods", http.MethodPost)

			handler.ServeHTTP(w, r)
		},
	)
}

func AddCors(
	w http.ResponseWriter,
	url string,
) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Origin", url)
	return w
}
