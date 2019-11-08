package models

type RpcClaimable struct {
	Claimables []Claimable `json:"claimable"`
	Address    string      `json:"address"`
	Unclaimed  float32     `json:"unclaimed"`
}

type Claimable struct {
	TxId        string  `json:"txid"`
	N           int     `json:"n"`
	Value       uint32  `json:"value"`
	StartHeight uint32  `json:"start_height"`
	EndHeight   uint32  `json:"end_height"`
	Generated   float32 `json:"generated"`
	SysFee      float32 `json:"sys_fee"`
	Unclaimed   float32 `json:"unclaimed"`
}
