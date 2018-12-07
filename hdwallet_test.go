//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
)

func aTestBTCTestnetVectors(t *testing.T) {
	assert.NotEqual(t, Random256Bits(), Random256Bits(), "Random isn't working")

	for _, testVector := range vectorsTESTNET {
		startingEntropy, _ := hex.DecodeString(testVector.entropy)

		mnemonic := Entropy2Mnemonic(startingEntropy)
		assert.Equal(t, testVector.mnemonic, mnemonic, "Mnemonic is incorrect")
		print("\n" + mnemonic)

		seed := Mnemonic2Seed(mnemonic)
		print("\n" + hex.EncodeToString(seed))

		assert.Equal(t, hex.EncodeToString(seed), testVector.seed, "Seed from Mnemonic is incorrect")

		xpriv := MasterKeyFromSeed(seed)
		assert.Equal(t, xpriv.String(), testVector.bip32Root, "Invalid xPriv")
		print("\n" + xpriv.String())

		//	bip32Extended := Bip32Extended(seed, testVector.coin, testVector.account, testVector.change)
		bip32Extended := Bip32Extended(seed, 0, 0, 0)

		assert.Equal(t, bip32Extended.String(), testVector.bip32ExPriv, "Invalid BIP32 Priv")
		assert.Equal(t, bip32Extended.PublicKey().String(), testVector.bip32ExPub, "Invalid BIP32 Public")
		print("\n" + bip32Extended.String())
		print("\n" + bip32Extended.PublicKey().String())

		addressDerivedUsingxPub := Bip44AddressFromXPub(bip32Extended.PublicKey(), testVector.addressIndex)
		assert.Equal(t, addressDerivedUsingxPub, testVector.address, " Incorrect Derviced XPUB address")

		btcAdd, btcPrivKey := Bip44Address(seed, testVector.coin, testVector.account, testVector.change, testVector.addressIndex)
		assert.Equal(t, btcAdd, testVector.address, "Invalid BTC Address")
		wifComp, _ := btcutil.NewWIF(btcPrivKey, &chaincfg.MainNetParams, true)
		assert.Equal(t, wifComp.String(), testVector.privKey, "Invalid BTC Address")
	}

}

// TestBTCVectors - Cycle through the list of test vectors below taken from https://www.coinomi.com/recovery-phrase-tool.html
// Need to add significantly more vectors for more edge cases, coins, accounts, change etc.
func aTestBTCVectors(t *testing.T) {
	assert.NotEqual(t, Random256Bits(), Random256Bits(), "Random isn't working")

	for _, testVector := range vectorsMAINNET {
		startingEntropy, _ := hex.DecodeString(testVector.entropy)

		mnemonic := Entropy2Mnemonic(startingEntropy)
		assert.Equal(t, mnemonic, testVector.mnemonic, "Mnemonic is incorrect")

		seed := Mnemonic2Seed(mnemonic)
		assert.Equal(t, hex.EncodeToString(seed), testVector.seed, "Seed from Mnemonic is incorrect")

		xpriv := MasterKeyFromSeed(seed)
		assert.Equal(t, xpriv.String(), testVector.bip32Root, "Invalid xPriv")

		bip32Extended := Bip32Extended(seed, testVector.coin, testVector.account, testVector.change)
		assert.Equal(t, bip32Extended.String(), testVector.bip32ExPriv, "Invalid BIP32 Priv")
		assert.Equal(t, bip32Extended.PublicKey().String(), testVector.bip32ExPub, "Invalid BIP32 Public")

		addressDerivedUsingxPub := Bip44AddressFromXPub(bip32Extended.PublicKey(), testVector.addressIndex)
		assert.Equal(t, addressDerivedUsingxPub, testVector.address, " Incorrect Derviced XPUB address")

		btcAdd, btcPrivKey := Bip44Address(seed, testVector.coin, testVector.account, testVector.change, testVector.addressIndex)
		assert.Equal(t, btcAdd, testVector.address, "Invalid BTC Address")
		wifComp, _ := btcutil.NewWIF(btcPrivKey, &chaincfg.MainNetParams, true)
		assert.Equal(t, wifComp.String(), testVector.privKey, "Invalid BTC Address")
	}

}

func aTestBench(t *testing.T) {

	testVector := vectorsMAINNET[0]

	assert.NotEqual(t, Random256Bits(), Random256Bits(), "Random isn't working")
	startingEntropy, _ := hex.DecodeString(testVector.entropy)
	mnemonic := Entropy2Mnemonic(startingEntropy)
	assert.Equal(t, mnemonic, testVector.mnemonic, "Mnemonic is incorrect")
	seed := Mnemonic2Seed(mnemonic)
	assert.Equal(t, hex.EncodeToString(seed), testVector.seed, "Seed from Mnemonic is incorrect")

	for i := 0; i < 10; i++ {
		print("\n", i)

		xpriv := MasterKeyFromSeed(seed)
		bip32Extended := Bip32Extended(seed, testVector.coin, testVector.account, testVector.change)
		addressDerivedUsingxPub := Bip44AddressFromXPub(bip32Extended.PublicKey(), testVector.addressIndex)
		btcAdd, btcPrivKey := Bip44Address(seed, testVector.coin, testVector.account, testVector.change, testVector.addressIndex)
		wifComp, _ := btcutil.NewWIF(btcPrivKey, &chaincfg.MainNetParams, true)
		Use(bip32Extended, xpriv, btcAdd, wifComp, addressDerivedUsingxPub)
	}
	Use(seed)
}

var vectorsMAINNET = []struct {
	coin         int
	account      int
	change       int
	addressIndex int
	entropy      string
	mnemonic     string
	seed         string
	bip32Root    string
	bip32ExPriv  string
	bip32ExPub   string
	address      string
	privKey      string
}{
	{
		0, 0, 0, 0,
		"000102030405060708090a0b0c0d0e0f",
		"abandon amount liar amount expire adjust cage candy arch gather drum buyer",
		"3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0",
		"xprv9s21ZrQH143K2XojduRLQnU8D8K59KSBoMuQKGx8dW3NBitFDMkYGiJPwZdanjZonM7eXvcEbxwuGf3RdkCyyXjsbHSkwtLnJcsZ9US42Gd",
		"xprvA2QWrMvVn11Cnc8Wv5XH22Phaz1eLLYUtUVCJxjRu3eSbPZk3WphdkqGBnAKiKtg3bxkL48zbf9C8jJKtbDhB4kTJuNfv3KZVRjxseHNNWk",
		"xpub6FPsFsTPcNZW16Cz274HPALS91r8joGLFhQo7M93TPBRUBttb48xBZ9k34oiG29Bvqfry9QyXPsGXSRE1kjut92Dgik1w6Whm1GU4F122n8",
		"128BCBZndgrPXzEgF4QbVR3jnQGwzRtEz5",
		"L35qaFLpbCc9yCzeTuWJg4qWnTs9BaLr5CDYcnJ5UnGmgLo8JBgk",
	},
	{
		1, 1, 1, 19,
		"000102030405060708090a0b0c0d0e0f",
		"abandon amount liar amount expire adjust cage candy arch gather drum buyer",
		"3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0",
		"xprv9s21ZrQH143K2XojduRLQnU8D8K59KSBoMuQKGx8dW3NBitFDMkYGiJPwZdanjZonM7eXvcEbxwuGf3RdkCyyXjsbHSkwtLnJcsZ9US42Gd",
		"xprv9zzh2iB8WtfwQLHwF49QpYX7smwzazXf6BUDN7GUNzBSLwkyz9uZesMAAFqWt1zCGba6XercMGatdXuAHfCoEjJzqgCUv2yDYrLET4itKR8",
		"xpub6Dz3SDi2MGEEcpNQM5gRBgTrRonUzTFWTQPpAVg5wKiRDk68XhDpCffe1WXnpjDakCvgi9dNxjLfFJ8K8dk6PzXgJq8AYiejXzTgsXCH3fi",
		"1H1L8sviYk59vfbT8Uo9hdshtbDHTnjMwi",
		"KxvwKJ9ZfAeU54rKPTpuWftrwoYjBrnGQFTrgGD4BanTjqaeppme",
	},
}

var vectorsTESTNET = []struct {
	coin         int
	account      int
	change       int
	addressIndex int
	entropy      string
	mnemonic     string
	seed         string
	bip32Root    string
	bip32ExPriv  string
	bip32ExPub   string
	address      string
	privKey      string
}{
	{
		1, 0, 0, 0,
		"000102030405060708090a0b0c0d0e0f",
		"abandon amount liar amount expire adjust cage candy arch gather drum buyer",
		"3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0",
		"xprv9s21ZrQH143K2XojduRLQnU8D8K59KSBoMuQKGx8dW3NBitFDMkYGiJPwZdanjZonM7eXvcEbxwuGf3RdkCyyXjsbHSkwtLnJcsZ9US42Gd",
		"xprvA2QWrMvVn11Cnc8Wv5XH22Phaz1eLLYUtUVCJxjRu3eSbPZk3WphdkqGBnAKiKtg3bxkL48zbf9C8jJKtbDhB4kTJuNfv3KZVRjxseHNNWk",
		"xpub6FPsFsTPcNZW16Cz274HPALS91r8joGLFhQo7M93TPBRUBttb48xBZ9k34oiG29Bvqfry9QyXPsGXSRE1kjut92Dgik1w6Whm1GU4F122n8",
		"128BCBZndgrPXzEgF4QbVR3jnQGwzRtEz5",
		"L35qaFLpbCc9yCzeTuWJg4qWnTs9BaLr5CDYcnJ5UnGmgLo8JBgk",
	},
	{
		1, 1, 1, 19,
		"000102030405060708090a0b0c0d0e0f",
		"abandon amount liar amount expire adjust cage candy arch gather drum buyer",
		"3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0",
		"xprv9s21ZrQH143K2XojduRLQnU8D8K59KSBoMuQKGx8dW3NBitFDMkYGiJPwZdanjZonM7eXvcEbxwuGf3RdkCyyXjsbHSkwtLnJcsZ9US42Gd",
		"xprv9zzh2iB8WtfwQLHwF49QpYX7smwzazXf6BUDN7GUNzBSLwkyz9uZesMAAFqWt1zCGba6XercMGatdXuAHfCoEjJzqgCUv2yDYrLET4itKR8",
		"xpub6Dz3SDi2MGEEcpNQM5gRBgTrRonUzTFWTQPpAVg5wKiRDk68XhDpCffe1WXnpjDakCvgi9dNxjLfFJ8K8dk6PzXgJq8AYiejXzTgsXCH3fi",
		"1H1L8sviYk59vfbT8Uo9hdshtbDHTnjMwi",
		"KxvwKJ9ZfAeU54rKPTpuWftrwoYjBrnGQFTrgGD4BanTjqaeppme",
	},
}
