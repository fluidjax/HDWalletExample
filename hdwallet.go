package main

import (
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

//Bip44Address get Bitcoin (BIP32,39,44) address & associated private key for given seed, account & index
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
	xpubRaw := key.PublicKey().String()
	xpubKey, _ := hdkeychain.NewKeyFromString(xpubRaw)
	chPubKey, _ := xpubKey.Child(uint32(addressIndex))
	//return chPubKey.Address(&chaincfg.MainNetParams)
	address, _ := chPubKey.Address(&chaincfg.MainNetParams)
	return address.String()
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

func main() {
	//see tests (comments denote labels on https://iancoleman.io/bip39/)
	startingEntropy := Random256Bits()
	mnemonic := Entropy2Mnemonic(startingEntropy)         //BIP39 Mnemonic
	seed := Mnemonic2Seed(mnemonic)                       //BIP39 Seed
	address, privateKey := Bip44Address(seed, 0, 0, 0, 0) //m/44'/0'/0'/0/0
	xPrivBIP39 := masterKeyFromSeed(seed)                 //BIP32 Root Key
	keyPairBIP32 := Bip32Extended(seed, 0, 0, 0)          // Generate BIP32 Extended Keypair (xPub/xPriv)
	exPrivateKeyBIP32 := keyPairBIP32.String()            //BIP32 Extended Private Key
	exPublicKeyBIP32 := keyPairBIP32.PublicKey().String() //BIP32 Extended Public Key

	addressDerivedUsingxPub := Bip44AddressFromXPub(keyPairBIP32.PublicKey(), 0) //derive 'address' using xPub (no private info)

	wifComp, _ := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true) //Generate WIF format Private key

	Use(address, privateKey, exPrivateKeyBIP32, exPublicKeyBIP32, keyPairBIP32, xPrivBIP39, addressDerivedUsingxPub)
	Use(address, privateKey, wifComp)
}
