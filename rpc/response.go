package rpc

import "github.com/joeqian10/neo-gogogo/rpc/models"

type RpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type ErrorResponse struct {
	Error RpcError `json:"error"`
	NetError error
}

func (r *ErrorResponse) HasError() bool {
	if len(r.Error.Message) == 0 && r.NetError == nil {
		return false
	}
	return true
}

func (r *ErrorResponse) GetErrorInfo() string {
	if r.NetError != nil {
		return r.NetError.Error()
	}
	return r.Error.Message
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ClaimGasResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
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

type GetAssetStateResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcAssetState `json:"result"`
}

type GetBalanceResponse struct {
	RpcResponse
	ErrorResponse
	Result models.AssetBalance `json:"result"`
}

type GetBestBlockHashResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
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

type GetBlockHeaderResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcBlockHeader `json:"result"`
}

type GetBlockHashResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetClaimableResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcClaimable `json:"result"`
}

type GetConnectionCountResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetContractStateResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcContractState `json:"result"`
}

type GetNep5BalancesResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep5Balances `json:"result"`
}

type GetNep5TransfersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep5Transfers `json:"result"`
}

type GetNewAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetPeersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcPeers `json:"result"`
}

type GetRawMemPoolResponse struct {
	RpcResponse
	ErrorResponse
	Result []string `json:"result"`
}

type GetRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type GetStorageResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetTransactionHeightResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetTxOutResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransactionOutput `json:"result"`
}

type GetUnclaimedGasResponse struct {
	RpcResponse
	ErrorResponse
	Result models.UnclaimedGasInWallet `json:"result"`
}

type GetUnclaimedResponse struct {
	RpcResponse
	ErrorResponse
	Result models.UnclaimedGasInAddress `json:"result"`
}

type GetUnspentsResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcUnspent `json:"result"`
}

type GetValidatorsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcValidator `json:"result"`
}

type GetVersionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcVersion `json:"result"`
}

type GetWalletHeightResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type ImportPrivKeyResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcAddress `json:"result"`
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

type ListPluginsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcListPlugin `json:"result"`
}

type ListAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcAddress `json:"result"`
}

type SendFromResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SendRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type SendToAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SubmitBlockResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type ValidateAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.ValidateAddress `json:"result"`
}

type CrossChainProofResponse struct {
	RpcResponse
	ErrorResponse
	CrosschainProof models.MPTProof `json:"result"`
}

type StateHeightResponse struct {
	RpcResponse
	ErrorResponse
	Result models.StateHeight `json:"result"`
}

type StateRootResponse struct {
	RpcResponse
	ErrorResponse
	Result models.StateRootState `json:"result"`
}
