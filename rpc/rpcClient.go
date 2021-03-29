package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// add IHttpClient for mock unit test
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RpcClient struct {
	Endpoint   *url.URL
	httpClient IHttpClient
}

func NewClient(endpoint string) *RpcClient {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil
	}
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}
	return &RpcClient{Endpoint: u, httpClient: netClient}
}

func (n *RpcClient) makeRequest(method string, params []interface{}, out interface{}) error {
	request := NewRequest(method, params)
	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", n.Endpoint.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Set("Connection", "close")
	req.Close = true
	res, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}
	return nil
}

func (n *RpcClient) ClaimGas(address string) ClaimGasResponse {
	response := ClaimGasResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("claimgas", params, &response)
	return response
}

// GetAccountState returns global assets info of an address
func (n *RpcClient) GetAccountState(address string) GetAccountStateResponse {
	response := GetAccountStateResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("getaccountstate", params, &response)
	return response
}

// the endpoint needs to use ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	response.NetError = n.makeRequest("getapplicationlog", params, &response)
	return response
}

func (n *RpcClient) GetAssetState(assetId string) GetAssetStateResponse {
	response := GetAssetStateResponse{}
	params := []interface{}{assetId}
	response.NetError = n.makeRequest("getassetstate", params, &response)
	return response
}

// the endpoint needs to use RpcWallet plugin
func (n *RpcClient) GetBalance(assetId string) GetBalanceResponse {
	response := GetBalanceResponse{}
	params := []interface{}{assetId}
	response.NetError = n.makeRequest("getbalance", params, &response)
	return response
}

func (n *RpcClient) GetBestBlockHash() GetBestBlockHashResponse {
	response := GetBestBlockHashResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getbestblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlockByHash(blockHash string) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{blockHash, 1}
	response.NetError = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockByIndex(index uint32) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{index, 1}
	response.NetError = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockCount() GetBlockCountResponse {
	response := GetBlockCountResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getblockcount", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeaderByHash(blockHash string) GetBlockHeaderResponse {
	response := GetBlockHeaderResponse{}
	params := []interface{}{blockHash, 1}
	response.NetError = n.makeRequest("getblockheader", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeaderByIndex(index uint32) GetBlockHeaderResponse {
	hash := n.GetBlockHash(index).Result
	return n.GetBlockHeaderByHash(hash)
}

func (n *RpcClient) GetBlockHash(index uint32) GetBlockHashResponse {
	response := GetBlockHashResponse{}
	params := []interface{}{index}
	response.NetError = n.makeRequest("getblockhash", params, &response)
	return response
}

func (n *RpcClient) GetClaimable(address string) GetClaimableResponse {
	response := GetClaimableResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("getclaimable", params, &response)
	return response
}

func (n *RpcClient) GetConnectionCount() GetConnectionCountResponse {
	response := GetConnectionCountResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getconnectioncount", params, &response)
	return response
}
func (n *RpcClient) GetContractState(scriptHash string) GetContractStateResponse {
	response := GetContractStateResponse{}
	params := []interface{}{scriptHash}
	response.NetError = n.makeRequest("getcontractstate", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Balances(address string) GetNep5BalancesResponse {
	response := GetNep5BalancesResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("getnep5balances", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Transfers(address string) GetNep5TransfersResponse {
	response := GetNep5TransfersResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("getnep5balances", params, &response)
	return response
}

// need RpcWallet plugin
func (n *RpcClient) GetNewAddress() GetNewAddressResponse {
	response := GetNewAddressResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getnewaddress", params, &response)
	return response
}

func (n *RpcClient) GetRawMemPool() GetRawMemPoolResponse {
	response := GetRawMemPoolResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getrawmempool", params, &response)
	return response
}

func (n *RpcClient) GetRawTransaction(txid string) GetRawTransactionResponse {
	response := GetRawTransactionResponse{}
	params := []interface{}{txid, 1}
	response.NetError = n.makeRequest("getrawtransaction", params, &response)
	return response
}

func (n *RpcClient) GetStorage(scripthash string, key string) GetStorageResponse {
	response := GetStorageResponse{}
	params := []interface{}{scripthash, key}
	response.NetError = n.makeRequest("getstorage", params, &response)
	return response
}

func (n *RpcClient) GetTransactionHeight(txid string) GetTransactionHeightResponse {
	response := GetTransactionHeightResponse{}
	params := []interface{}{txid}
	response.NetError = n.makeRequest("gettransactionheight", params, &response)
	return response
}

func (n *RpcClient) GetTxOut(txid string, index int) GetTxOutResponse {
	response := GetTxOutResponse{}
	params := []interface{}{txid, index}
	response.NetError = n.makeRequest("gettxout", params, &response)
	return response
}

func (n *RpcClient) GetPeers() GetPeersResponse {
	response := GetPeersResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getpeers", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimedGas() GetUnclaimedGasResponse {
	response := GetUnclaimedGasResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getunclaimedgas", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimed(address string) GetUnclaimedResponse {
	response := GetUnclaimedResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("getunclaimed", params, &response)
	return response
}

func (n *RpcClient) GetUnspents(adddress string) GetUnspentsResponse {
	response := GetUnspentsResponse{}
	params := []interface{}{adddress}
	response.NetError = n.makeRequest("getunspents", params, &response)
	return response
}

func (n *RpcClient) GetValidators() GetValidatorsResponse {
	response := GetValidatorsResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getvalidators", params, &response)
	return response
}

func (n *RpcClient) GetVersion() GetVersionResponse {
	response := GetVersionResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getversion", params, &response)
	return response
}

func (n *RpcClient) GetWalletHeight() GetWalletHeightResponse {
	response := GetWalletHeightResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getwalletheight", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) ImportPrivKey(wif string) ImportPrivKeyResponse {
	response := ImportPrivKeyResponse{}
	params := []interface{}{wif}
	response.NetError = n.makeRequest("importprivkey", params, &response)
	return response
}

func (n *RpcClient) InvokeFunction(scriptHash string, method string, checkWitnessHashes string, args ...interface{}) InvokeFunctionResponse {
	response := InvokeFunctionResponse{}
	var params []interface{}
	if args != nil {
		params = []interface{}{scriptHash, method, args, checkWitnessHashes}
	} else {
		params = []interface{}{scriptHash, method, checkWitnessHashes}
	}
	response.NetError = n.makeRequest("invokefunction", params, &response)
	return response
}

func (n *RpcClient) InvokeScript(scriptInHex string, checkWitnessHashes string) InvokeScriptResponse {
	response := InvokeScriptResponse{}
	params := []interface{}{scriptInHex, checkWitnessHashes}
	response.NetError = n.makeRequest("invokescript", params, &response)
	return response
}

func (n *RpcClient) ListPlugins() ListPluginsResponse {
	response := ListPluginsResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("listplugins", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) ListAddress() ListAddressResponse {
	response := ListAddressResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("listaddress", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) SendFrom(assetId string, from string, to string, amount uint32, fee float32, changeAddress string) SendFromResponse {
	response := SendFromResponse{}
	params := []interface{}{assetId, from, to, amount, fee, changeAddress}
	response.NetError = n.makeRequest("sendfrom", params, &response)
	return response
}

func (n *RpcClient) SendRawTransaction(rawTransactionInHex string) SendRawTransactionResponse {
	response := SendRawTransactionResponse{}
	params := []interface{}{rawTransactionInHex, 1}
	response.NetError = n.makeRequest("sendrawtransaction", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) SendToAddress(assetId string, to string, amount uint32, fee float32, changeAddress string) SendToAddressResponse {
	response := SendToAddressResponse{}
	params := []interface{}{assetId, to, amount, fee, changeAddress}
	response.NetError = n.makeRequest("sendtoaddress", params, &response)
	return response
}

// TODO: sendmany

func (n *RpcClient) SubmitBlock(blockHex string) SubmitBlockResponse {
	response := SubmitBlockResponse{}
	params := []interface{}{blockHex}
	response.NetError = n.makeRequest("submitblock", params, &response)
	return response
}

func (n *RpcClient) ValidateAddress(address string) ValidateAddressResponse {
	response := ValidateAddressResponse{}
	params := []interface{}{address}
	response.NetError = n.makeRequest("validateaddress", params, &response)
	return response
}

func (n *RpcClient) GetProof(stateroot, contractScriptHash, storeKey string) CrossChainProofResponse {
	response := CrossChainProofResponse{}
	params := []interface{}{stateroot, contractScriptHash, storeKey}
	response.NetError = n.makeRequest("getproof", params, &response)
	return response
}

func (n *RpcClient) GetStateHeight() StateHeightResponse {
	response := StateHeightResponse{}
	params := []interface{}{}
	response.NetError = n.makeRequest("getstateheight", params, &response)
	return response
}

func (n *RpcClient) GetStateRootByIndex(blockHeight uint32) StateRootResponse {
	response := StateRootResponse{}
	params := []interface{}{blockHeight}
	response.NetError = n.makeRequest("getstateroot", params, &response)
	return response
}

func (n *RpcClient) GetStateRootByHash(blockHash string) StateRootResponse {
	response := StateRootResponse{}
	params := []interface{}{blockHash}
	response.NetError = n.makeRequest("getstateroot", params, &response)
	return response
}
