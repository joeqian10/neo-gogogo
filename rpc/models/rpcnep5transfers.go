package models

type RpcNep5Transfers struct {
	Sent     []Nep5Transfer `json:"sent"`
	Received []Nep5Transfer `json:"received"`
	Address  string         `json:"address"`
}

type Nep5Transfer struct {
	Timestamp           uint32 `json:"timestamp"`
	AssetHash           string `json:"asset_hash"`
	TransferAddress     string `json:"transfer_address"`
	Amount              string `json:"amount"`
	BlockIndex          uint32 `json:"block_index"`
	TransferNotifyIndex uint32 `json:"transfer_notify_index"`
	TxHash              string `json:"tx_hash"`
}
