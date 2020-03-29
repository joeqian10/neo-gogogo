package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	_ = n.makeRequest("claimgas", params, &response)
	return response
}

// GetAccountState returns global assets info of an address
func (n *RpcClient) GetAccountState(address string) GetAccountStateResponse {
	response := GetAccountStateResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getaccountstate", params, &response)
	return response
}

// the endpoint needs to use ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	_ = n.makeRequest("getapplicationlog", params, &response)
	return response
}

func (n *RpcClient) GetAssetState(assetId string) GetAssetStateResponse {
	response := GetAssetStateResponse{}
	params := []interface{}{assetId}
	_ = n.makeRequest("getassetstate", params, &response)
	return response
}

// the endpoint needs to use RpcWallet plugin
func (n *RpcClient) GetBalance(assetId string) GetBalanceResponse {
	response := GetBalanceResponse{}
	params := []interface{}{assetId}
	_ = n.makeRequest("getbalance", params, &response)
	return response
}

func (n *RpcClient) GetBestBlockHash() GetBestBlockHashResponse {
	response := GetBestBlockHashResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getbestblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlockByHash(blockHash string) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{blockHash, 1}
	_ = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockByIndex(index uint32) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{index, 1}
	_ = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockCount() GetBlockCountResponse {
	response := GetBlockCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getblockcount", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeaderByHash(blockHash string) GetBlockHeaderResponse {
	response := GetBlockHeaderResponse{}
	params := []interface{}{blockHash, 1}
	_ = n.makeRequest("getblockheader", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeaderByIndex(index uint32) GetBlockHeaderResponse {
	hash := n.GetBlockHash(index).Result
	return n.GetBlockHeaderByHash(hash)
}

func (n *RpcClient) GetBlockHash(index uint32) GetBlockHashResponse {
	response := GetBlockHashResponse{}
	params := []interface{}{index}
	_ = n.makeRequest("getblockhash", params, &response)
	return response
}

func (n *RpcClient) GetClaimable(address string) GetClaimableResponse {
	response := GetClaimableResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getclaimable", params, &response)
	return response
}

func (n *RpcClient) GetConnectionCount() GetConnectionCountResponse {
	response := GetConnectionCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getconnectioncount", params, &response)
	return response
}
func (n *RpcClient) GetContractState(scriptHash string) GetContractStateResponse {
	response := GetContractStateResponse{}
	params := []interface{}{scriptHash}
	_ = n.makeRequest("getcontractstate", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Balances(address string) GetNep5BalancesResponse {
	response := GetNep5BalancesResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep5balances", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Transfers(address string) GetNep5TransfersResponse {
	response := GetNep5TransfersResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep5balances", params, &response)
	return response
}

// need RpcWallet plugin
func (n *RpcClient) GetNewAddress() GetNewAddressResponse {
	response := GetNewAddressResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getnewaddress", params, &response)
	return response
}

func (n *RpcClient) GetRawMemPool() GetRawMemPoolResponse {
	response := GetRawMemPoolResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getrawmempool", params, &response)
	return response
}

func (n *RpcClient) GetRawTransaction(txid string) GetRawTransactionResponse {
	response := GetRawTransactionResponse{}
	params := []interface{}{txid, 1}
	_ = n.makeRequest("getrawtransaction", params, &response)
	return response
}

func (n *RpcClient) GetStorage(scripthash string, key string) GetStorageResponse {
	response := GetStorageResponse{}
	params := []interface{}{scripthash, key}
	_ = n.makeRequest("getstorage", params, &response)
	return response
}

func (n *RpcClient) GetTransactionHeight(txid string) GetTransactionHeightResponse {
	response := GetTransactionHeightResponse{}
	params := []interface{}{txid}
	_ = n.makeRequest("gettransactionheight", params, &response)
	return response
}

func (n *RpcClient) GetTxOut(txid string, index int) GetTxOutResponse {
	response := GetTxOutResponse{}
	params := []interface{}{txid, index}
	_ = n.makeRequest("gettxout", params, &response)
	return response
}

func (n *RpcClient) GetPeers() GetPeersResponse {
	response := GetPeersResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getpeers", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimedGas() GetUnclaimedGasResponse {
	response := GetUnclaimedGasResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getunclaimedgas", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimed(address string) GetUnclaimedResponse {
	response := GetUnclaimedResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getunclaimed", params, &response)
	return response
}

func (n *RpcClient) GetUnspents(adddress string) GetUnspentsResponse {
	response := GetUnspentsResponse{}
	params := []interface{}{adddress}
	_ = n.makeRequest("getunspents", params, &response)
	return response
}

func (n *RpcClient) GetValidators() GetValidatorsResponse {
	response := GetValidatorsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getvalidators", params, &response)
	return response
}

func (n *RpcClient) GetVersion() GetVersionResponse {
	response := GetVersionResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getversion", params, &response)
	return response
}

func (n *RpcClient) GetWalletHeight() GetWalletHeightResponse {
	response := GetWalletHeightResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getwalletheight", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) ImportPrivKey(wif string) ImportPrivKeyResponse {
	response := ImportPrivKeyResponse{}
	params := []interface{}{wif}
	_ = n.makeRequest("importprivkey", params, &response)
	return response
}

func (n *RpcClient) InvokeFunction(scriptHash string, method string, args ...interface{}) InvokeFunctionResponse {
	response := InvokeFunctionResponse{}
	var params []interface{}
	if args != nil {
		params = []interface{}{scriptHash, method, args}
	} else {
		params = []interface{}{scriptHash, method}
	}
	_ = n.makeRequest("invokefunction", params, &response)
	return response
}

func (n *RpcClient) InvokeScript(scriptInHex string) InvokeScriptResponse {
	response := InvokeScriptResponse{}
	params := []interface{}{scriptInHex}
	_ = n.makeRequest("invokescript", params, &response)
	return response
}

func (n *RpcClient) ListPlugins() ListPluginsResponse {
	response := ListPluginsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listplugins", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) ListAddress() ListAddressResponse {
	response := ListAddressResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listaddress", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) SendFrom(assetId string, from string, to string, amount uint32, fee float32, changeAddress string) SendFromResponse {
	response := SendFromResponse{}
	params := []interface{}{assetId, from, to, amount, fee, changeAddress}
	_ = n.makeRequest("sendfrom", params, &response)
	return response
}

func (n *RpcClient) SendRawTransaction(rawTransactionInHex string) SendRawTransactionResponse {
	response := SendRawTransactionResponse{}
	params := []interface{}{rawTransactionInHex, 1}
	_ = n.makeRequest("sendrawtransaction", params, &response)
	return response
}

// need RpcWallet
func (n *RpcClient) SendToAddress(assetId string, to string, amount uint32, fee float32, changeAddress string) SendToAddressResponse {
	response := SendToAddressResponse{}
	params := []interface{}{assetId, to, amount, fee, changeAddress}
	_ = n.makeRequest("sendtoaddress", params, &response)
	return response
}

// TODO: sendmany

func (n *RpcClient) SubmitBlock(blockHex string) SubmitBlockResponse {
	response := SubmitBlockResponse{}
	params := []interface{}{blockHex}
	_ = n.makeRequest("submitblock", params, &response)
	return response
}

func (n *RpcClient) ValidateAddress(address string) ValidateAddressResponse {
	response := ValidateAddressResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("validateaddress", params, &response)
	return response
}

func (n *RpcClient) GetProof(stateroot, contractScriptHash, storeKey string) CrossChainProofResponse {
	response := CrossChainProofResponse{}
	params := []interface{}{stateroot, contractScriptHash, storeKey}
	_ = n.makeRequest("getproof", params, &response)
	return response
}

func (n *RpcClient) GetStateHeight() StateHeightResponse {
	response := StateHeightResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getstateheight", params, &response)
	return response
}

func (n *RpcClient) GetStateRootByIndex(blockHeight uint32) StateRootResponse {
	response := StateRootResponse{}
	params := []interface{}{blockHeight}
	_ = n.makeRequest("getstateroot", params, &response)
	return response
}

func (n *RpcClient) GetStateRootByHash(blockHash string) StateRootResponse {
	response := StateRootResponse{}
	params := []interface{}{blockHash}
	_ = n.makeRequest("getstateroot", params, &response)
	return response
}

func (n *RpcClient) IsTxConfirmed(txId string) (bool, error) {
	r1 := n.GetRawTransaction(txId)
	if r1.HasError() {
		return false, fmt.Errorf("GetRawTransaction: %s", r1.Error.Message)
	}
	return r1.Result.Confirmations > 0, nil
}

func (n *RpcClient) WaitForTransactionConfirmed(txId string) error {
	checkPeriod := 9*time.Second
	checkTimeout := 45*time.Second

	start := time.Now()
	first := true
	for time.Since(start) < checkTimeout {
		if first {
			time.Sleep(checkPeriod/2)
			first = false
		} else {
			time.Sleep(checkPeriod)
		}
		accepted, err := n.IsTxConfirmed(txId)
		if err != nil {return err}
		if accepted {return nil}
	}
	return fmt.Errorf("timed out waiting for %s", txId)
}
