// To parse and unparse this JSON data, add this code to your project and do:
//
//    transactions, err := UnmarshalTransactions(bytes)
//    bytes, err = transactions.Marshal()

package NoRestDataStructures

import "encoding/json"

func UnmarshalTransactions(data []byte) (Transactions, error) {
	var r Transactions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Transactions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Transactions struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

type Data struct {
	Total  int64    `json:"total"`
	Result []Result `json:"result"`
}

type Result struct {
	Txid          string `json:"txid"`
	Hash          string `json:"hash"`
	Version       int64  `json:"version"`
	Size          int64  `json:"size"`
	Vsize         int64  `json:"vsize"`
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
	Txid        string    `json:"txid"`
	Vout        int64     `json:"vout"`
	ScriptSig   ScriptSig `json:"scriptSig"`
	Txinwitness []string  `json:"txinwitness"`
	Sequence    int64     `json:"sequence"`
	Address     Vout      `json:"address"`
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
