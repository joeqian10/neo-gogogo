package models

type RpcTransaction struct {
	Txid          string                     `json:"txid"`
	Size          int                        `json:"size"`
	Type          string                     `json:"type"`
	Version       int                        `json:"version"`
	Attributes    []RpcTransactionAttribute `json:"attributes"`
	Vin           []RpcTransactionInput     `json:"vin"`
	Vout          []RpcTransactionOutput    `json:"vout"`
	SysFee        string                     `json:"sys_fee"`
	NetFee        string                     `json:"net_fee"`
	Scripts       []RpcWitness              `json:"scripts"`
	Nonce         int                        `json:"nonce"`
	BlockHash     string                     `json:"blockhash"`
	Confirmations int                        `json:"confirmations"`
	Blocktime     int                        `json:"blocktime"`
	Script        string                     `json:"script"`
	Gas           string                     `json:"gas"`
	Claims        []RpcClaim                 `json:"claims"`
}

type RpcTransactionAttribute struct {
	Usage string `json:"usage"`
	Data  string `json:"data"`
}

type RpcTransactionInput struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type RpcTransactionOutput struct {
	N       int    `json:"n"`
	Asset   string `json:"asset"`
	Value   string `json:"value"`
	Address string `json:"address"`
}

type RpcClaim struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}
