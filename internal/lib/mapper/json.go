package mapper

import (
	"encoding/json"
	"fmt"
	"os"
)

func JSON[data interface{}](path string) (*data, error) {
	var (
		bytes []byte
		model = new(data)
		err   error
	)

	bytes, err = os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config file not found in '%s'", path)
	}

	err = json.Unmarshal(bytes, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
