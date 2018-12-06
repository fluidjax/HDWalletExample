//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockCypherInterface(t *testing.T) {
	testAddress := "1AozLV7krw87WKxjCzzygM29BrYFxbxPwh"
	unusedAddress := "161PrtobFv8DWaYLDXChzZg9euXeg7URxd"

	balance1, _ := blockCypherGetBalance(testAddress)
	assert.Equal(t, balance1, "44333771577", "Invalid Balance")

	balance2, _ := blockCypherGetBalance(unusedAddress)
	assert.Equal(t, balance2, "0", "Invalid Balance")

	isUsed1 := blockCypherHasBeenUsed(testAddress)
	assert.Equal(t, isUsed1, true, "incorrectly detected as not used")

	isUsed2 := blockCypherHasBeenUsed(unusedAddress)
	assert.Equal(t, isUsed2, false, "incorrectly detected as  used")
}
