package models

type RpcTransaction struct {
	Txid          string                    `json:"txid"`
	Size          uint32                    `json:"size"`
	Type          string                    `json:"type"`
	Version       uint32                    `json:"version"`
	Attributes    []RpcTransactionAttribute `json:"attributes"`
	Vin           []RpcTransactionInput     `json:"vin"`
	Vout          []RpcTransactionOutput    `json:"vout"`
	SysFee        string                    `json:"sys_fee"`
	NetFee        string                    `json:"net_fee"`
	Scripts       []RpcWitness              `json:"scripts"`
	Nonce         uint32                    `json:"nonce"`
	BlockHash     string                    `json:"blockhash"`
	Confirmations uint32                    `json:"confirmations"`
	Blocktime     uint32                    `json:"blocktime"`
	Script        string                    `json:"script"`
	Gas           string                    `json:"gas"`
	Claims        []RpcClaim                `json:"claims"`
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

type RpcWitness struct {
	Invocation   string `json:"invocation"`
	Verification string `json:"verification"`
}

type RpcClaim struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}
