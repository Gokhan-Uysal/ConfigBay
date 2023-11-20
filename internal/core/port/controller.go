package port

import "net/http"

type (
	PageController interface {
		Root(w http.ResponseWriter, r *http.Request)
		Home(w http.ResponseWriter, r *http.Request)
		Signup(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
	}

	OnboardController interface {
		SignupWith(w http.ResponseWriter, r *http.Request)
		LoginWith(w http.ResponseWriter, r *http.Request)
		RedirectGoogle(w http.ResponseWriter, r *http.Request)
	}
)
