//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func aTestBuildJson(t *testing.T) {
	seed, _ := hex.DecodeString("e82822b1f215de5459c7777d5a18e725b5b369d1f25471b4fa02b6e1d6e6eef531bd76c662a6ced2585be158cb74849898752c5f79e3d547624ce9b2a0771d1e")
	noRestBuildWallet(seed)
}

func TestNoRestInterface(t *testing.T) {
	seed, _ := hex.DecodeString("e82822b1f215de5459c7777d5a18e725b5b369d1f25471b4fa02b6e1d6e6eef531bd76c662a6ced2585be158cb74849898752c5f79e3d547624ce9b2a0771d1e")

	testAddress := "187XjYegq5vMa9gETNmRz6JeE6rPXJZAuM"
	unusedAddress := "161PrtobFv8DWaYLDXChzZg9euXeg7URxd"

	balance1, _ := noRestAddressGetBalance(testAddress)
	assert.Equal(t, "0.00001000", balance1, "Invalid Balance")

	balance2, _ := noRestAddressGetBalance(unusedAddress)
	assert.Equal(t, "0.00000000", balance2, "Invalid Balance")

	isUsed1 := noRestAddressHasBeenUsed(testAddress)
	assert.Equal(t, true, isUsed1, "incorrectly detected as not used")

	isUsed2 := noRestAddressHasBeenUsed(unusedAddress)
	assert.Equal(t, false, isUsed2, "incorrectly detected as used")

	totalBalance := noRestGetWalletBalance(seed)
	assert.Equal(t, "0.00021000", totalBalance, "incorrect total Balance")
	Use(seed)

	transactions := noRestGetTransactions("12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX")
	assert.Equal(t, int64(3), transactions.Data.Total)

}
