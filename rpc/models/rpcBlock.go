package models

type RpcBlockHeader struct {
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	Version           int `json:"version"`
	Previousblockhash string `json:"previousblockhash"`
	Merkleroot        string `json:"merkleroot"`
	Time              int `json:"time"`
	Index             int `json:"index"`
	Nonce             string `json:"nonce"`         //ulong = uint64
	Nextconsensus     string `json:"nextconsensus"` //address
	CrossStatesRoot   string `json:"crossstatesroot"`
	ChainID           string `json:"chainid"` //ulong = uint64
	Script            struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	} `json:"script"`
	Confirmations int    `json:"confirmations"`
	NextBlockHash string `json:"nextblockhash"`
}

type RpcBlock struct {
	RpcBlockHeader
	Tx []RpcTransaction `json:"tx"`
}
