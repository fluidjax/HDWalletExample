//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tidwall/gjson"
)

//BIP 32 - xPub/xPriv from seed
//BIP 39 - Mnemonic Wordlist
//BIP 44 - m / purpose' / coin_type' / account' / change / address_index

/*
   Some Useful tools for testsing
   Mnemonic Code Converter tool
   https://iancoleman.io/bip39/

   Key Convertor
   https://www.bitaddress.org
*/

//Bip44Address get Bitcoin (BIP32,39,44) address & associated private key for given seed, account & index
func Bip44Address(seed []byte, coin int, account int, change int, addressIndex int) (string, *btcec.PrivateKey) {
	bip32extended := Bip32Extended(seed, coin, account, change)
	add1, _ := bip32extended.NewChildKey(uint32(addressIndex))
	privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), add1.Key)
	Use(privKey)
	caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
	return caddr.EncodeAddress(), privKey
}

//MasterKeyFromSeed Generate Bitcoin address from XPub
func MasterKeyFromSeed(seed []byte) *bip32.Key {
	masterKey, _ := bip32.NewMasterKey(seed)
	return masterKey
}

//Bip44AddressFromXPub Generate Bitcoin address from XPub
func Bip44AddressFromXPub(key *bip32.Key, addressIndex int) string {
	xpubRaw := key.PublicKey().String()
	xpubKey, _ := hdkeychain.NewKeyFromString(xpubRaw)
	chPubKey, _ := xpubKey.Child(uint32(addressIndex))
	//return chPubKey.Address(&chaincfg.MainNetParams)
	address, _ := chPubKey.Address(&chaincfg.MainNetParams)
	return address.String()
}

//Bip32Extended get Bip32 extended Keys for path
func Bip32Extended(seed []byte, coin int, account int, change int) *bip32.Key {
	masterKey := MasterKeyFromSeed(seed)
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

//Random256Bits generate a randomseed
func Random256Bits() []byte {
	entropy, _ := bip39.NewEntropy(256)
	return entropy
}

//Use - helper to remove warnings
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func blockCypherHasBeenUsed(address string) bool {
	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "final_n_tx")

		if value.Num > 0 {
			return true
		}
	}
	return false
}

func blockCypherGetBalance(address string) (string, error) {
	url := "https://api.blockcypher.com/v1/btc/main/addrs/" + address + "/balance"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The get Balance request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "balance")
		return value.String(), nil
	}
	return "0", errors.New("Failed to check Address Balance")
}

func nextUnusedAddress(seed []byte, startIndex int) string {
	for i := startIndex; ; i++ {
		address, _ := Bip44Address(seed, 0, 0, 0, i)
		used := blockCypherHasBeenUsed(address)
		if used == false {
			return address
		}
	}
}

func main() {
	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)

	nextUnusedAddress := nextUnusedAddress(seed, 0)
	fmt.Printf("Next Unused Address = %s\n", nextUnusedAddress)

}
