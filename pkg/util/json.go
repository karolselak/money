package util

import (
	"io/ioutil"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

// JSON accessor to json-iterator lib
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

// OpenJSON returns a pointer to the json file
func OpenJSON(f string) *os.File {

	JSONFile, err := os.Open(f)
	if err != nil {
		log.Fatal("open failed")
	}
	return JSONFile
}

// ReadJSON unmarshels jsons and return a byte
func ReadJSON(JF *os.File, v interface{}) {
	byteValue, err := ioutil.ReadAll(JF)
	if err != nil {
		log.Fatalf("reading failed")
	}
	err = JSON.Unmarshal(byteValue, v)
	if err != nil {
		log.Fatal(err)
	}
}

// writeJSON writes json
func WriteJSON(s string, v interface{}) {
	wjson, _ := JSON.Marshal(&v)
	err := ioutil.WriteFile(s, wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
