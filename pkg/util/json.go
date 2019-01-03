package util

import (
	jsoniter "github.com/json-iterator/go"
)

var j = jsoniter.ConfigCompatibleWithStandardLibrary

func Unmarshal(bytes []byte, v interface{}) error {
	err := j.Unmarshal(bytes, v)
	return err
}

func Marshal(v interface{}) ([]byte, error) {
	bytes, err := j.Marshal(&v)
	return bytes, err
}
