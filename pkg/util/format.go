package util

import (
	"log"
	"strconv"
)

// Stf converts string to float
func Stf(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
