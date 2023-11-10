package controller

import (
	"encoding/json"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"net/http"
)

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

func (bc baseController) handleError(w http.ResponseWriter, httpErr payload.HTTPError) {
	var (
		data []byte
		err  error
	)
	data, err = json.Marshal(httpErr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpErr.StatusCode)
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func (bc baseController) handleResponse(w http.ResponseWriter, body interface{}) error {
	var (
		data []byte
		err  error
	)

	data, err = json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
