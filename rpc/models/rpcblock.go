package models

type RpcBlockHeader struct {
	Hash              string `json:"hash"`
	Size              uint32    `json:"size"`
	Version           uint32 `json:"version"`
	Previousblockhash string `json:"previousblockhash"`
	Merkleroot        string `json:"merkleroot"`
	Time              uint32 `json:"time"`
	Index             uint32 `json:"index"`
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
