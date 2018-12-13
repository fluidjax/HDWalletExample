//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited
//  NORest

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

//var urlStub = "http://127.0.0.1:8080/api/v1/"
var urlStub = "http://btc.mousebelt.com/api/v1/"

func noRestAddressHasBeenUsed(address string) bool {
	url := urlStub + "address/txs/" + address
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "data.total")

		if value.Num > 0 {
			//fmt.Printf("Used %s \n", url)
			return true
		}
	}
	//fmt.Printf("Unused %s \n", url)
	return false
}

func noRestAddressGetBalance(address string) (string, error) {
	url := urlStub + "balance/" + address
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "data.balance")
		coins := fmt.Sprintf("%.8f", value.Float())
		return coins, nil
	}
	return "0.00000000", errors.New("Failed to check Address Balance")
}

func noRestNextUnusedAddress(seed []byte, startIndex int) string {
	for i := startIndex; ; i++ {
		address, _ := Bip44Address(seed, 0, 0, 0, i)
		used := noRestAddressHasBeenUsed(address)
		if used == false {
			return address
		}
	}
}

func noRestGetWalletBalance(seed []byte) string {
	//loop addresses until unused, cummulating balance total
	var totalBalance float64
	for i := 0; ; i++ {
		address, _ := Bip44Address(seed, 0, 0, 0, i)
		if noRestAddressHasBeenUsed(address) == false {
			break
		}
		//time.Sleep(1000 * time.Millisecond)
		balanceString, err := noRestAddressGetBalance(address)

		if err != nil {
			fmt.Printf("Error checking balance in Get Wallet Balance %s\n", err)
		} else {
			balance, _ := strconv.ParseFloat(balanceString, 64)
			totalBalance = totalBalance + balance
			fmt.Printf("Address:%s balance:%f total:%f \n", address, balance, totalBalance)
		}
	}
	return fmt.Sprintf("%.8f", totalBalance)
}
