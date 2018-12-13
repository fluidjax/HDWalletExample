//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited
//  Using Blockcypher as a temporary block chain indexer

package main

// var token = "?token=2ea83a3e4efa439f9ecc18deb5781baf"
//
// func addressHasBeenUsed(address string) bool {
// 	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address + token
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("The get Balance request failed with error %s\n", err)
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		value := gjson.Get(string(data), "final_n_tx")
//
// 		if value.Num > 0 {
// 			fmt.Printf("Used %s \n", url)
// 			return true
// 		}
// 	}
// 	fmt.Printf("Unused %s \n", url)
// 	return false
// }
//
// func addressGetBalance(address string) (string, error) {
// 	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address + token
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("The get Balance request failed with error %s\n", err)
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		value := gjson.Get(string(data), "balance")
// 		return value.String(), nil
// 	}
// 	return "0", errors.New("Failed to check Address Balance")
// }
//
// func nextUnusedAddress(seed []byte, startIndex int) string {
// 	for i := startIndex; ; i++ {
// 		address, _ := Bip44Address(seed, 0, 0, 0, i)
// 		used := addressHasBeenUsed(address)
// 		if used == false {
// 			return address
// 		}
// 	}
// }
//
// func getWalletBalance(seed []byte) string {
// 	//loop addresses until unused, cummulating balance total
// 	var totalBalance float64
// 	for i := 0; ; i++ {
// 		address, _ := Bip44Address(seed, 0, 0, 0, i)
// 		if addressHasBeenUsed(address) == false {
// 			break
// 		}
// 		time.Sleep(1000 * time.Millisecond)
// 		balanceString, err := addressGetBalance(address)
//
// 		if err != nil {
// 			fmt.Printf("Error checking balance in Get Wallet Balance %s\n", err)
// 		} else {
// 			balance, _ := strconv.ParseFloat(balanceString, 64)
// 			totalBalance = totalBalance + balance
// 			fmt.Printf("Address:%s balance:%f total:%f \n", address, balance, totalBalance)
// 		}
// 	}
// 	return fmt.Sprintf("%.8f", totalBalance/100000000)
// }
