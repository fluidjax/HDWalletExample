//  (c)Christopher Morris 2018
//  cmorris@qredo.com
//  Qredo Limited

package main

import (
	"encoding/hex"
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

//{"status":200,"msg":"success","data":{"address":"1ALfUfD9NL5TKitZVxec9vuzgmndHstxJb","balance":0.00533265}}
func getBalance(address string) string {
	response, err := http.Get("https://btc.mousebelt.com/api/v1/balance/" + address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		value := gjson.Get(string(data), "data.balance")
		return value.String()
	}
	return "0.00000000"
}

func getWalletAddresses(seed []byte, coin int, account int, change int, addressIndexStart int, addressIndexEnd int) []string {
	var res []string
	for add := addressIndexStart; add < addressIndexEnd; add++ {
		btcAdd, _ := Bip44Address(seed, coin, account, change, add)

		res = append(res, btcAdd)
	}
	return res
}

func hasTransactions(address string) bool {
	url := "http://btc.mousebelt.com/api/v1/address/txs/" + address
	print(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		value := gjson.Get(string(data), "data.total")

		if value.Num > 0 {
			return true
		}
	}
	return false
}

func main() {
	entropy, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	mnemonic := Entropy2Mnemonic(entropy)
	seed := Mnemonic2Seed(mnemonic)
	Use(seed)

	//Setup backend https://github.com/norestlabs/mousexplore-vcoins
	//No Rest API https://github.com/norestlabs/mousexplore-vcoins/blob/master/Bitcoin/API.md
	//Endpoints http://mousexplore.mousebelt.com/btc

	//test backend from NoRest
	//API https://github.com/norestlabs/mousexplore-vcoins/wiki/Bitcoin
	//Example call http://127.0.0.1:8080/monitor/db
	//Online call https://btc.mousebelt.com/monitor
	//test address 1MqBLR1eMhDzZF3mzdjhZk8VrLj7MzrGwb

	// https://btc.mousebelt.com/api/v1/block/552379

	addresses := getWalletAddresses(seed, 0, 0, 0, 0, 9)
	address0 := addresses[0]
	address1 := addresses[1]

	fmt.Printf("Get Balance %s = %s \n", address0, getBalance(address0))
	fmt.Printf("Get Balance %s = %s \n", address1, getBalance(address1))

	fmt.Printf("Has Transactions %s = %t \n", address0, hasTransactions(address0))
	fmt.Printf("Has Transactions %s = %t \n", address1, hasTransactions(address1))

	//http://btc.mousebelt.com/api/v1/address/txs/1ALfUfD9NL5TKitZVxec9vuzgmndHstxJb?offset=0
	//https://btc.mousebelt.com/api/v1/balance/1ALfUfD9NL5TKitZVxec9vuzgmndHstxJb
	//print("Get Balance Bitcoin address")

	//addresses := getWalletAddresses(seed, 0, 0, 0, 0, 9)

	//fmt.Printf("Has transactions %s = %t \n", addresses[0], hasTransactions(addresses[0]))

	// print("Get addresses in wallet")
	// print("Get Balance of HD Wallet")

	// print("Get next unused deposit address")

	// print("Withdraw Transction")
	// print("BroadCast Transaction")

	// 	//see tests (comments denote labels on https://iancoleman.io/bip39/)
	// 	startingEntropy := Random256Bits()
	// 	mnemonic := Entropy2Mnemonic(startingEntropy)         //BIP39 Mnemonic
	// 	seed := Mnemonic2Seed(mnemonic)                       //BIP39 Seed
	// 	address, privateKey := Bip44Address(seed, 0, 0, 0, 0) //m/44'/0'/0'/0/0
	// 	xPrivBIP39 := MasterKeyFromSeed(seed)                 //BIP32 Root Key
	// 	keyPairBIP32 := Bip32Extended(seed, 0, 0, 0)          // Generate BIP32 Extended Keypair (xPub/xPriv)
	// 	exPrivateKeyBIP32 := keyPairBIP32.String()            //BIP32 Extended Private Key
	// 	exPublicKeyBIP32 := keyPairBIP32.PublicKey().String() //BIP32 Extended Public Key

	// 	addressDerivedUsingxPub := Bip44AddressFromXPub(keyPairBIP32.PublicKey(), 0) //derive 'address' using xPub (no private info)

	// 	wifComp, _ := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true) //Generate WIF format Private key

	// 	print("hello")
	// 	Use(address, privateKey, exPrivateKeyBIP32, exPublicKeyBIP32, keyPairBIP32, xPrivBIP39, addressDerivedUsingxPub)
	// 	Use(address, privateKey, wifComp)
}
