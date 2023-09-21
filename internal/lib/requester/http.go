package requester

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Make(req *http.Request) (*http.Response, error) {
	var (
		client *http.Client
		resp   *http.Response
		err    error
	)

	client = &http.Client{}
	defer client.CloseIdleConnections()

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func UnmarshalBody[TData interface{}](body io.ReadCloser) (*TData, error) {
	var (
		bytes   = make([]byte, 0)
		bodyObj = new(TData)
		err     error
	)

	if body == nil {
		return nil, errors.New("body is missing")
	}

	defer func(reader io.ReadCloser) {
		if err := reader.Close(); err != nil {
			return
		}
	}(body)

	bytes, err = io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, bodyObj)
	if err != nil {
		return nil, err
	}

	return bodyObj, nil
}
