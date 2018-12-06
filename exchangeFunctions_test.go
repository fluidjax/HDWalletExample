//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepositAPI(t *testing.T) {
	//Gte next unused address for deposit

	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)

	nextUnusedAddress := nextUnusedAddress(seed, 0)
	assert.Equal(t, nextUnusedAddress, "1E7NvpF3u87rbpfYxt3HDmpFasPiU2JhMp", "Incorrect Exchange Withdraw Address ")
}

func TestGetBalance(t *testing.T) {
	//get the total balance for all used addresses for a given seed

	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)
	Use(seed)
}

func TestWithdrawal(t *testing.T) {
	//Create a signed Raw Withdrawal Transaction

	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)
	Use(seed)
}
