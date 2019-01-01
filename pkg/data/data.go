package data

import (
	"io/ioutil"
	"log"
	"os"
)

type Data interface {
	readData()
	writeData()
}

func (d Assets) readData() {

}

func (d Assets) writeData() {

}

// ReadJSON unmarshels jsons and return a byte
func ReadData(JF *os.File) {
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
func writeData(Forte Assets) {
	wjson, _ := JSON.Marshal(Forte)
	err := ioutil.WriteFile(Conf.dataFile, wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
