package rpc

// add IRpcClient for mock UT
type IRpcClient interface {
	ClaimGas(s string) ClaimGasResponse
	GetAccountState(s string) GetAccountStateResponse
	GetApplicationLog(s string) GetApplicationLogResponse
	GetAssetState(s string) GetAssetStateResponse
	GetBalance(s string) GetBalanceResponse
	GetBestBlockHash() GetBestBlockHashResponse
	GetBlockByHash(s string) GetBlockResponse
	GetBlockByIndex(n uint32) GetBlockResponse
	GetBlockCount() GetBlockCountResponse
	GetBlockHeaderByHash(s string) GetBlockHeaderResponse
	GetBlockHash(n uint32) GetBlockHashResponse
	GetClaimable(s string) GetClaimableResponse
	GetConnectionCount() GetConnectionCountResponse
	GetContractState(s string) GetContractStateResponse
	GetNep5Balances(s string) GetNep5BalancesResponse
	GetNep5Transfers(s string) GetNep5TransfersResponse
	GetNewAddress() GetNewAddressResponse
	GetPeers() GetPeersResponse
	GetRawMemPool() GetRawMemPoolResponse
	GetRawTransaction(s string) GetRawTransactionResponse
	GetStorage(s1 string, s2 string) GetStorageResponse
	GetTransactionHeight(s string) GetTransactionHeightResponse
	GetTxOut(s string, n int) GetTxOutResponse
	GetUnclaimed(s string) GetUnclaimedResponse
	GetUnclaimedGas() GetUnclaimedGasResponse
	GetUnspents(s string) GetUnspentsResponse
	GetValidators() GetValidatorsResponse
	GetVersion() GetVersionResponse
	GetWalletHeight() GetWalletHeightResponse
	ImportPrivKey(s string) ImportPrivKeyResponse
	InvokeFunction(s1 string, s2 string, args ...interface{}) InvokeFunctionResponse
	InvokeScript(s string) InvokeScriptResponse
	ListPlugins() ListPluginsResponse
	ListAddress() ListAddressResponse
	SendFrom(assetId string, from string, to string, amount uint32, fee float32, changeAddress string) SendFromResponse
	SendRawTransaction(s string) SendRawTransactionResponse
	SendToAddress(assetId string, to string, amount uint32, fee float32, changeAddress string) SendToAddressResponse
	SubmitBlock(s string) SubmitBlockResponse
	ValidateAddress(s string) ValidateAddressResponse
	IsTxConfirmed(txId string) (bool, error)
	WaitForTransactionConfirmed(txId string) error
}