package port

import "net/http"

type (
	ProjectController interface {
	}

	PageController interface {
		Home(w http.ResponseWriter, r *http.Request)
	}
)
