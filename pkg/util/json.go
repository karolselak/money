package util

import (
	"encoding/json"
)

// Unmarshal converts a Json file to an array of bytes to be read
func Unmarshal(bytes []byte, v interface{}) error {
	err := json.Unmarshal(bytes, v)
	return err
}

// Marshal converts a byte array to a json to be writtrn
func Marshal(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	return bytes, err
}
