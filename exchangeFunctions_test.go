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
	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)

	nextUnusedAddress := nextUnusedAddress(seed, 0)
	assert.Equal(t, nextUnusedAddress, "1E7NvpF3u87rbpfYxt3HDmpFasPiU2JhMp", "Incorrect Exchange Withdraw Address ")
}
