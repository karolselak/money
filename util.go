package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	cmc "github.com/coincircle/go-coinmarketcap"
)

// Close closes open resource
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// NumDig returns the numbers of digits
func NumDig(n float64) int {
	return len(strconv.FormatFloat(n, 'f', 0, 64))
}

// OpenJSON returns a pointer to the json file
func OpenJSON() *os.File {

	usr, _ := user.Current()
	dir := usr.HomeDir
	file := filepath.Join(dir, "go/src/github.com/mohfunk/netWorth/data/assets.json")
	JSONFile, err := os.Open(file)
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
		log.Fatalf("Unmarshal failed")
	}
}

// writeJSON writes json
func writeJSON(Forte Assets) {
	wjson, _ := JSON.Marshal(Forte)
	err := ioutil.WriteFile("data/assets.json", wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// getPrice fetches the price of a currency
func getPrice(sym string) float64 {

	price, err := cmc.Price(&cmc.PriceOptions{
		Symbol:  sym,
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}
	return price
}

// stf converts string to float
func stf(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
