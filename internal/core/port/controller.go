package port

import "net/http"

type (
	PageController interface {
		Home(w http.ResponseWriter, r *http.Request)
	}
)
