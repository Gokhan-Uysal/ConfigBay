package controller

import (
	"encoding/json"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"net/http"
)

type baseController struct {
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
