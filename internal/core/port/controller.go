package port

import "net/http"

type (
	PageController interface {
		Home(w http.ResponseWriter, r *http.Request)
		Root(w http.ResponseWriter, r *http.Request)
	}

	OnboardController interface {
		Signup(w http.ResponseWriter, r *http.Request)
	}
)
