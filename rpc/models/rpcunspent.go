package models

type RpcUnspent struct {
	Balance []UnspentBalance `json:"balance"`
	Address string `json:"address"`
}

type UnspentBalance struct {
	Unspent     []Unspent `json:"unspent"`
	AssetHash   string    `json:"asset_hash"`
	Asset       string    `json:"asset"`
	AssetSymbol string    `json:"asset_symbol"`
	Amount      int       `json:"amount"`
}

type Unspent struct {
	Txid  string  `json:"txid"`
	N     int     `json:"n"`
	Value float32 `json:"value"`
}
