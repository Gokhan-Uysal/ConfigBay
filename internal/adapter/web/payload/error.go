package payload

import "net/http"

type HTTPError struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

var InternalServerErr = HTTPError{
	StatusCode:    http.StatusInternalServerError,
	StatusMessage: "internal server error",
}

var PageNotFound = HTTPError{
	StatusCode:    http.StatusNotFound,
	StatusMessage: "page not found",
}

var Unauthorized = HTTPError{
	StatusCode:    http.StatusUnauthorized,
	StatusMessage: "unauthorized access",
}

var Forbidden = HTTPError{
	StatusCode:    http.StatusForbidden,
	StatusMessage: "forbidden access",
}
