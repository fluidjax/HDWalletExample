//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func aTestBlockCypherInterface(t *testing.T) {

	testAddress := "13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM"
	unusedAddress := "161PrtobFv8DWaYLDXChzZg9euXeg7URxd"

	balance1, _ := addressGetBalance(testAddress)
	assert.Equal(t, balance1, "5004532403", "Invalid Balance")

	balance2, _ := addressGetBalance(unusedAddress)
	assert.Equal(t, balance2, "0", "Invalid Balance")

	isUsed1 := addressHasBeenUsed(testAddress)
	assert.Equal(t, isUsed1, true, "incorrectly detected as not used")

	isUsed2 := addressHasBeenUsed(unusedAddress)
	assert.Equal(t, isUsed2, false, "incorrectly detected as  used")
}
