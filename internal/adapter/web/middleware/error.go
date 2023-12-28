package middleware

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, httpErr payload.HTTPError) {
	w.WriteHeader(httpErr.StatusCode)
	_, err := w.Write([]byte(httpErr.StatusMessage))
	if err != nil {
		return
	}
}
