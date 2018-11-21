package main

import (
  "fmt"
   "encoding/hex"
	 "github.com/FactomProject/go-bip32" 
	 "github.com/FactomProject/go-bip39"
   "github.com/btcsuite/btcutil"
   "github.com/btcsuite/btcd/chaincfg"
   "github.com/btcsuite/btcd/btcec"
)

/* 
   Some Useful tools
   Mnemonic Code Converter tool
   https://www.coinomi.com/recovery-phrase-tool.html

   Key Convertor
   https://www.bitaddress.org
   */


func main(){
  //mnemonic
 //tv_seed           := "000102030405060708090a0b0c0d0e0f"
 //tv_BIP39_tv_seed  := "3779b041fab425e9c0fd55846b2a03e9a388fb12784067bd8ebdb464c2574a05bcc7a8eb54d7b2a2c8420ff60f630722ea5132d28605dbc996c8ca7d7a8311c0"
 //tv_BIP39_mnemonic := "abandon amount liar amount expire adjust cage candy arch gather drum buyer"
 //tv_BIP32_rootKey  := "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"
 
 //entropy, _ := bip39.NewEntropy(256)
 //entropy, _    := hex.DecodeString(tv_seed)
 //mnemonic, _   := bip39.NewMnemonic(entropy)
 
 mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"
 seed          := bip39.NewSeed(mnemonic, "")
 masterPrivateKey, _ := bip32.NewMasterKey(seed)
 
 // fmt.Println(hex.EncodeToString(seed))
 // fmt.Println(mnemonic)
  fmt.Println(masterPrivateKey)

  //fKey, _ := NewKeyFromMasterKey(masterPrivateKey, 0x00000000, 0x00000000, 0, 0)
  //xprvA2SBJXYRzBjhZQ1MvmKLKYcgnWpn9DRVLZkrsBqhEU4yzgwmYKwcwppzPS5hPRctN9PZjzVyYJiXPrwDiSMUDKLxWqVQb1Prd9diJxUFed6
  
  // fKey1, _ := NewKeyFromMasterKey(masterPrivateKey, 0, 0, 0, 0)
  // fmt.Println(fKey1)
  // 
  // fKey2, _ := NewKeyFromMnemonic(mnemonic, 0, 0, 0, 0)
  // fmt.Println(fKey2)
  // 
  // fKey2
  
  
  seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
  masterKey, err := bip32.NewMasterKey(seed)
  child1, err := masterKey.NewChildKey(0x8000002C) //purpose 44
  child2, err := child1.NewChildKey(0x80000000)    //cointype 0 = bitcoin
  child3, err := child2.NewChildKey(0x80000000)   //account 0 
  child4, err := child3.NewChildKey(0)   // change

  //fmt.Println("BIP32 Extended:", child4)


  add1, err := child4.NewChildKey(0)
    
  
  //fmt.Println(hex.EncodeToString(add1.ChainCode))
  fmt.Println(hex.EncodeToString(add1.Key))  //Private Key Hexadecimal Format
  
  //fmt.Println(add1)
  
  
  
//  fmt.Println(hex.EncodeToString(add.Key))
  
  fmt.Println(err)
  
  //addrString := "326f12c6fd4341c1544fcf241caf8ad53df7476067ab8ad6988ba23be858ef35"
  
  privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), add1.Key)
  
  wifComp, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
  fmt.Println("Compressed WIF", wifComp)
  wifUnComp, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
  fmt.Println("Uncompressed WIF", wifUnComp)
  
  caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
  uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)
  
  fmt.Println(caddr.EncodeAddress())
  fmt.Println(uaddr.EncodeAddress())
  
  
  //wif := btcutil.NewWIF(addrString,&chaincfg.MainNetParams)
  
  // addr, err := btcutil.DecodeAddress(addrString, &chaincfg.MainNetParams)
  // if err != nil {
	// fmt.Println(err)
	// return
  // }
  // fmt.Println(addr.EncodeAddress())
  
  // 326F12C6FD4341C1544FCF241CAF8AD53DF7476067AB8AD6988BA23BE858EF35
  // Mm8Sxv1DQcFUT88kHK+K1T33R2Bnq4rWmIuiO+hY7zU=
  // KxukKhTPU11xH2Wfk2366e375166QE4r7y8FWojU9XPbzLYYSM3j
  // 5JCVr6YMVms4TjwJPSA5nENmskX9vkHjPsn9GpHsZHrXarbj3hf
  // 02F7AA1EB14DE438735C026C7CC719DB11BAF82E47F8FA2C86B55BFF92B677EAE2
  // 04F7AA1EB14DE438735C026C7CC719DB11BAF82E47F8FA2C86B55BFF92B677EAE27FC055C6AD3B3D8770DC9A8291C39CE909094CEF7CCE12DF8C90B90DDDF1B2BA
  // 12UGT3397x9TvWpheYQvS2HfKysjWHENSr
  // 1Hdh7MEWekWD4qiHVRa2H8Ar3JR8sXunE
  
  
  //private key KxukKhTPU11xH2Wfk2366e375166QE4r7y8FWojU9XPbzLYYSM3j
  //looking for 1Hdh7MEWekWD4qiHVRa2H8Ar3JR8sXunE
  //public 02f7aa1eb14de438735c026c7cc719db11baf82e47f8fa2c86b55bff92b677eae2
}


