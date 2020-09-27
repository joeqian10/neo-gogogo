package models

type RpcNep5Balances struct {
	Balances []Nep5Balance `json:"balance"`
	Address  string        `json:"address"`
}

type Nep5Balance struct {
	AssetHash        string `json:"asset_hash"`
	Amount           string    `json:"amount"`
	LastUpdatedBlock int    `json:"last_updated_block"`
}
