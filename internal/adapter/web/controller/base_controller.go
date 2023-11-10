package controller

import "net/http"

type baseController struct {
}

func (bc baseController) enableCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Headers", "content-yype")
	w.Header().Add("Access-Control-Allow-Headers", "origin")
	w.Header().Add("Access-Control-Allow-Headers", "accept")
	w.Header().Add("Access-Control-Allow-Headers", "authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Add("Access-Control-Allow-Methods", http.MethodGet)
	w.Header().Add("Access-Control-Allow-Methods", http.MethodPost)
	return w
}

func (bc baseController) addCors(
	w http.ResponseWriter,
	url string,
) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Origin", url)
	return w
}
