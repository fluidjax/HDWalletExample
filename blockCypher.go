package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func addressHasBeenUsed(address string) bool {
	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "final_n_tx")

		if value.Num > 0 {
			return true
		}
	}
	return false
}

func addressGetBalance(address string) (string, error) {
	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address + "/balance?token=2ea83a3e4efa439f9ecc18deb5781baf"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "balance")
		return value.String(), nil
	}
	return "0", errors.New("Failed to check Address Balance")
}

func nextUnusedAddress(seed []byte, startIndex int) string {
	for i := startIndex; ; i++ {
		address, _ := Bip44Address(seed, 0, 0, 0, i)
		used := addressHasBeenUsed(address)
		if used == false {
			return address
		}
	}
}
