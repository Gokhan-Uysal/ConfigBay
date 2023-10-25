package payload

import "net/http"

type HTTPError struct {
	StatusCode    int
	StatusMessage string
}

var InternalServerErr = HTTPError{
	StatusCode:    http.StatusInternalServerError,
	StatusMessage: "Internal server error",
}

var PageNotFound = HTTPError{
	StatusCode:    http.StatusNotFound,
	StatusMessage: "Page not found",
}

var Unauthorized = HTTPError{
	StatusCode:    http.StatusUnauthorized,
	StatusMessage: "Unauthorized access",
}

var Forbidden = HTTPError{
	StatusCode:    http.StatusForbidden,
	StatusMessage: "Forbidden access",
}
