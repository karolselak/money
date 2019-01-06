package util

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Convert provides currency conversion
func Convert(from, to string, amount float64) float64 {

	url := "https://free.currencyconverterapi.com/api/v5/convert?q=" + from + "_" + to + "&compact=y"

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		if err2 == nil {
			strval := strings.Replace(bodyString, "{\""+from+"_"+to+"\":{\"val\":", "", 1)
			strval = strings.Replace(strval, "}}", "", 1)
			rate, err3 := strconv.ParseFloat(strval, 64)
			if err3 == nil {
				return rate * amount
			}
		}
	}
	return 0
}
