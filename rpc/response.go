package rpc

import "github.com/joeqian10/neo-gogogo/rpc/models"

type RpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type GetAccountStateResponse struct {
	RpcResponse
	ErrorResponse
	Result models.AccountState `json:"result"`
}

type GetApplicationLogResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcApplicationLog `json:"result"`
}

type GetBlockHeaderResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcBlockHeader `json:"result"`
}

type GetBlockResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcBlock `json:"result"`
}

type GetBlockCountResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetContractStateResponse struct {
	RpcResponse
	ErrorResponse
	Result models.ContractState `json:"result"`
}

type GetRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}
type InvokeFunctionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.InvokeResult `json:"result"`
}

type InvokeScriptResponse struct {
	RpcResponse
	ErrorResponse
	Result models.InvokeResult `json:"result"`
}

type SendRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type GetStorageResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetUnspentsResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcUnspent `json:"result"`
}

type GetCrosschainProofResponse struct {
	RpcResponse
	ErrorResponse
	CrosschainProof string `json:"result"`
}
