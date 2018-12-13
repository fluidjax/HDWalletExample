//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

// func aTestDepositAPI(t *testing.T) {
// 	//Gte next unused address for deposit
//
// 	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
//
// 	mnemonic := Entropy2Mnemonic(entropy)
// 	seed := Mnemonic2Seed(mnemonic)
//
// 	nextUnusedAddress := nextUnusedAddress(seed, 0)
// 	assert.Equal(t, nextUnusedAddress, "1E7NvpF3u87rbpfYxt3HDmpFasPiU2JhMp", "Incorrect Exchange Withdraw Address ")
// }
//
// func aTestGetBalance(t *testing.T) {
// 	//get the total balance for all used addresses for a given seed
//
// 	entropy, _ := hex.DecodeString("c0ed16aaeb28289fbfd1bfaf40166cbf")
// 	mnemonic := Entropy2Mnemonic(entropy)
// 	seed := Mnemonic2Seed(mnemonic)
// 	balance := getWalletBalance(seed)
// 	print(balance)
//
// }
//
// func aTestWithdrawal(t *testing.T) {
// 	//Create a signed Raw Withdrawal Transaction
//
// 	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
// 	mnemonic := Entropy2Mnemonic(entropy)
// 	seed := Mnemonic2Seed(mnemonic)
// 	Use(seed)
// }
