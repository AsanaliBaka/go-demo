package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var paylod T

	err := json.NewDecoder(body).Decode(&paylod)

	if err != nil {
		return paylod, err
	}

	return paylod, nil
}
