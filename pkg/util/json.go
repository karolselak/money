package util

import (
	"encoding/json"
)

func Unmarshal(bytes []byte, v interface{}) error {
	err := json.Unmarshal(bytes, v)
	return err
}

func Marshal(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}
