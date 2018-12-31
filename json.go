package main

import (
	"io/ioutil"
	"log"
	"os"
)

// OpenJSON returns a pointer to the json file
func OpenJSON() *os.File {

	JSONFile, err := os.Open(Conf.dataFile)
	if err != nil {
		log.Fatal("open failed")
	}
	return JSONFile
}

// ReadJSON unmarshels jsons and return a byte
func ReadJSON(JF *os.File) {
	byteValue, err := ioutil.ReadAll(JF)
	if err != nil {
		log.Fatalf("reading failed")
	}
	err = JSON.Unmarshal(byteValue, &Forte)
	if err != nil {
		log.Fatal(err)
	}
}

// writeJSON writes json
func writeJSON(Forte Assets) {
	wjson, _ := JSON.Marshal(Forte)
	err := ioutil.WriteFile(Conf.dataFile, wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
