package controller

import "net/http"

type baseController struct {
}

func (bc baseController) enableCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", " Authorization")

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
