// To parse and unparse this JSON data, add this code to your project and do:
//
//    hDWallet, err := UnmarshalHDWallet(bytes)
//    bytes, err = hDWallet.Marshal()

package main

import "encoding/json"

func UnmarshalHDWallet(data []byte) (HDWallet, error) {
	var r HDWallet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *HDWallet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type HDWallet struct {
	Username         string    `json:"username"`
	Seed             string    `json:"seed"`
	XPUB             string    `json:"xPUB"`
	LastUpdatedBlock int64     `json:"lastUpdatedBlock"`
	CoinName         string    `json:"coinName"`
	Purpose          int64     `json:"purpose"`
	CoinType         int64     `json:"coinType"`
	Account          int64     `json:"account"`
	Address          []Address `json:"address"`
	Change           []Change  `json:"change"`
}

type Address struct {
	Index        int64         `json:"index"`
	Type         string        `json:"type"`
	Address      string        `json:"address"`
	Balance      string        `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Txid          string `json:"txid"`
	Hash          string `json:"hash"`
	Version       int64  `json:"version"`
	Size          int64  `json:"size"`
	Vsize         int64  `json:"vsize"`
	Weight        int64  `json:"weight"`
	Locktime      int64  `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations int64  `json:"confirmations"`
	Time          int64  `json:"time"`
	Blocktime     int64  `json:"blocktime"`
}

type Vin struct {
	Txid      string    `json:"txid"`
	Vout      int64     `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  int64     `json:"sequence"`
	Address   Vout      `json:"address"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int64        `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	ASM       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int64    `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

type ScriptSig struct {
	ASM string `json:"asm"`
	Hex string `json:"hex"`
}

type Change struct {
	CurrrenlyUnused bool `json:"currrenly unused"`
}
