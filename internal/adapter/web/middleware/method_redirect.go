package middleware

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"net/http"
)

func Get(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Post(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Put(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Patch(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPatch {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Delete(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Connect(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodConnect {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Head(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodHead {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}

func Trace(handler http.Handler) http.Handler {
	handler = RequestLogger(handler)
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodOptions {
				ErrorHandler(w, payload.PageNotFound)
				return
			}
			handler.ServeHTTP(w, r)
		},
	)
}
