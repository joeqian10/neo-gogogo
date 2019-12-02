package models

type RpcClaimable struct {
	Claimables []Claimable `json:"claimable"`
	Address    string      `json:"address"`
	Unclaimed  float64     `json:"unclaimed"`
}

type Claimable struct {
	TxId        string  `json:"txid"`
	N           int     `json:"n"`
	Value       int  	`json:"value"`
	StartHeight int  	`json:"start_height"`
	EndHeight   int  	`json:"end_height"`
	Generated   float64 `json:"generated"`
	SysFee      float64 `json:"sys_fee"`
	Unclaimed   float64 `json:"unclaimed"`
}
