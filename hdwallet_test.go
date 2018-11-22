package main

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
)

// func TestBitcoinWallet(t *testing.T) {
// 	mnemonic := "abandon amount liar amount expire adjust cage candy arch gather drum buyer"
// 	seed := Mnemonic2Seed(mnemonic)
// 	fmt.Println("mnemonic: ", mnemonic)
// 	add1, privKey := Bip44Address(seed, 0, 1)
// 	fmt.Println("Address1: ", add1)
// 	Use(privKey)
//
// 	// seedString := hex.EncodeToString(seed)
// 	// entropy = "000102030405060708090a0b0c0d0e0f"
// 	// mnemonic, _ := bip39.NewMnemonic(entropy)
// 	//
// 	//assert.Equal(t, "000102030405060708090a0b0c0d0e0f", seedString, "Invalid seed")
//
// }

func TestBTCVectors(t *testing.T) {
	for _, testVector := range vectors {

		startingEntropy, _ := hex.DecodeString(testVector.entropy)

		mnemonic := Entropy2Mnemonic(startingEntropy)
		assert.Equal(t, mnemonic, testVector.mnemonic, "Invalid Seed")

		seed := Mnemonic2Seed(mnemonic)
		assert.Equal(t, hex.EncodeToString(seed), testVector.seed, "Invalid Seed")

		xpriv := masterKeyFromSeed(seed)
		assert.Equal(t, xpriv.String(), testVector.bip32Root, "Invalid X Priv")

		btcAdd1, btcPrivKey1 := Bip44Address(seed, 0, 0, 0, 0)
		assert.Equal(t, btcAdd1, testVector.add000BTCAdd, "Invalid BTC Address")
		wifComp, _ := btcutil.NewWIF(btcPrivKey1, &chaincfg.MainNetParams, true)
		assert.Equal(t, wifComp.String(), testVector.add000privKey, "Invalid BTC Address")

		btcAdd2, btcPrivKey2 := Bip44Address(seed, 1, 1, 1, 19)
		assert.Equal(t, btcAdd2, testVector.add1119BTCAdd, "Invalid BTC Address")
		wifComp2, _ := btcutil.NewWIF(btcPrivKey2, &chaincfg.MainNetParams, true)
		assert.Equal(t, wifComp2.String(), testVector.add1119privKey, "Invalid BTC Address")

	}
}

//Vectors manually made from https://www.coinomi.com/recovery-phrase-tool.html
var vectors = []struct {
	entropy        string
	mnemonic       string
	seed           string
	bip32Root      string
	bip32ExPriv    string
	bip32ExPub     string
	add000BTCAdd   string
	add000privKey  string
	add1119BTCAdd  string //address coin=1, account=1, addressindex = 19
	add1119privKey string
}{
	{
		"000102030405060708090a0b0c0d0e0f",
		"abandon amount liar amount expire adjust cage candy arch gather drum buyer",
		"3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0",
		"xprv9s21ZrQH143K2XojduRLQnU8D8K59KSBoMuQKGx8dW3NBitFDMkYGiJPwZdanjZonM7eXvcEbxwuGf3RdkCyyXjsbHSkwtLnJcsZ9US42Gd",
		"xprvA2QWrMvVn11Cnc8Wv5XH22Phaz1eLLYUtUVCJxjRu3eSbPZk3WphdkqGBnAKiKtg3bxkL48zbf9C8jJKtbDhB4kTJuNfv3KZVRjxseHNNWk",
		"xpub6FPsFsTPcNZW16Cz274HPALS91r8joGLFhQo7M93TPBRUBttb48xBZ9k34oiG29Bvqfry9QyXPsGXSRE1kjut92Dgik1w6Whm1GU4F122n8",
		"128BCBZndgrPXzEgF4QbVR3jnQGwzRtEz5",
		"L35qaFLpbCc9yCzeTuWJg4qWnTs9BaLr5CDYcnJ5UnGmgLo8JBgk",
		"1H1L8sviYk59vfbT8Uo9hdshtbDHTnjMwi",
		"KxvwKJ9ZfAeU54rKPTpuWftrwoYjBrnGQFTrgGD4BanTjqaeppme",
	},
}
