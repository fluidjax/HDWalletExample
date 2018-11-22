package main

import (
	"fmt"

	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

//BIP 32 - xPub/xPriv from seed
//BIP 39 - Mnemonic Wordlist
//BIP 44 - m / purpose' / coin_type' / account' / change / address_index

/*
   Some Useful tools for testsing
   Mnemonic Code Converter tool
   https://www.coinomi.com/recovery-phrase-tool.html

   Key Convertor
   https://www.bitaddress.org
*/

//GetBitcoinAddress get Bitcoin (BIP32,39,44) address & associated private key for given seed, account & index
func Bip44Address(seed []byte, coin int, account int, change int, addressIndex int) (string, *btcec.PrivateKey) {
	bip32extended := Bip32Extended(seed, coin, account, change)
	//fmt.Println("bip32extended: ", bip32extended.PublicKey())
	add1, _ := bip32extended.NewChildKey(uint32(addressIndex))
	privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), add1.Key)
	Use(privKey)
	caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
	return caddr.EncodeAddress(), privKey
}

func masterKeyFromSeed(seed []byte) *bip32.Key {
	masterKey, _ := bip32.NewMasterKey(seed)
	return masterKey
}

//Bip44AddressFromXPub Generate Bitcoin address from XPub
func Bip44AddressFromXPub(key *bip32.Key, addressIndex int) string {
	xpubRaw := "xpub6FPsFsTPcNZW16Cz274HPALS91r8joGLFhQo7M93TPBRUBttb48xBZ9k34oiG29Bvqfry9QyXPsGXSRE1kjut92Dgik1w6Whm1GU4F122n8"
	offset := 0

	xpubKey, _ := hdkeychain.NewKeyFromString(xpubRaw)
	keys := make([]*hdkeychain.ExtendedKey, 10)
	for i := uint32(0); i < uint32(len(keys)); i++ {
		chPubKey, _ := xpubKey.Child(i + uint32(offset))
		fmt.Println(chPubKey.Address(&chaincfg.MainNetParams))
	}

	// return keys
	//
	// fmt.Println("XPUB: ", xpubRaw)
	//
	// devKey, _ := key.NewChildKey(uint32(addressIndex))
	// fmt.Println("PUBLIC0:", devKey.PublicKey())
	// fmt.Println("PUBLIC1:", devKey.String())
	// fmt.Println("PUBLIC2:", hex.EncodeToString(devKey.ChainCode))
	// fmt.Println("PUBLIC3:", hex.EncodeToString(devKey.FingerPrint))
	// fmt.Println("PUBLIC4:", hex.EncodeToString(devKey.Key))
	// //privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), a.Key)
	// //Use(privKey)
	//
	// //fmt.Println("PRIVATE:", privKey)
	//caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
	return "hello"
}

//Bip32Extended get Bip32 extended Keys for path
func Bip32Extended(seed []byte, coin int, account int, change int) *bip32.Key {
	masterKey := masterKeyFromSeed(seed)
	child1, _ := masterKey.NewChildKey(0x8000002C)                //purpose 44
	child2, _ := child1.NewChildKey(0x80000000 + uint32(coin))    //cointype 0 = bitcoin
	child3, _ := child2.NewChildKey(0x80000000 + uint32(account)) //account 0
	bip32extended, _ := child3.NewChildKey(uint32(change))
	return bip32extended
}

//Entropy2Mnemonic convert a seed (entropy) to a Mmemonic String
func Entropy2Mnemonic(entropy []byte) string {
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return mnemonic
}

//Mnemonic2Seed convert BIP39 mnemonic (recovery phrase) to a seed byte array
func Mnemonic2Seed(mnemonic string) []byte {
	password := "" //we arent using passwords
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, password)
	return seed
}

//RandomSeed generate a randomseed
func RandomSeed() []byte {
	entropy, _ := bip39.NewEntropy(256)
	return entropy
}

func main() {
	mnemonic := "abandon amount liar amount expire adjust cage candy arch gather drum buyer"
	seed := Mnemonic2Seed(mnemonic)
	fmt.Println("mnemonic: ", mnemonic)

	add1, privKey := Bip44Address(seed, 0, 0, 0, 1)
	fmt.Println("Address1: ", add1)

	Use(privKey)
}

// func TestSum(t *testing.T) {
// 	fmt.Println("DONE")
// }
//
//Use - helper to remove warnings
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
