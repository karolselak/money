package util

import (
	"io/ioutil"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/"
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

