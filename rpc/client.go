package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type RpcClient struct {
	Endpoint   url.URL
	httpClient *http.Client
}

func NewClient(endpoint string) *RpcClient {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil
	}
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}
	return &RpcClient{Endpoint: *u, httpClient: netClient}
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

// GetAccountState returns global assets info of an address
func (n *RpcClient) GetAccountState(address string) GetAccountStateResponse {
	response := GetAccountStateResponse{}
	params := []interface{}{address}
	n.makeRequest("getaccountstate", params, &response)
	return response
}

// the endpoint needs to use ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	n.makeRequest("getapplicationlog", params, &response)
	return response
}

//
func (n *RpcClient) GetBlockByHash(blockHash string) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{blockHash, 1}
	n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockByIndex(index uint32) GetBlockResponse {
	response := GetBlockResponse{}
	params := []interface{}{index, 1}
	n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeaderByIndex(index uint32) GetBlockHeaderResponse {
	response := GetBlockHeaderResponse{}
	params := []interface{}{index, 1}
	n.makeRequest("getblockheader", params, &response)
	return response
}

func (n *RpcClient) GetBlockCount() GetBlockCountResponse {
	response := GetBlockCountResponse{}
	params := []interface{}{}
	n.makeRequest("getblockcount", params, &response)
	return response
}

func (n *RpcClient) InvokeFunction(scriptHash string, method string, args []interface{}) InvokeFunctionResponse {
	response := InvokeFunctionResponse{}
	params := []interface{}{scriptHash, method, args}
	n.makeRequest("invokefunction", params, &response)
	return response
}

func (n *RpcClient) InvokeScript(scriptInHex string) InvokeScriptResponse {
	response := InvokeScriptResponse{}
	params := []interface{}{scriptInHex, 1}
	n.makeRequest("invokescript", params, &response)
	return response
}

func (n *RpcClient) SendRawTransaction(rawTransactionInHex string) SendRawTransactionResponse {
	response := SendRawTransactionResponse{}
	params := []interface{}{rawTransactionInHex, 1}
	n.makeRequest("sendrawtransaction", params, &response)
	return response
}

func (n *RpcClient) GetStorage(scripthash string, key string) GetStorageResponse {
	response := GetStorageResponse{}
	params := []interface{}{scripthash, key}
	n.makeRequest("getstorage", params, &response)
	return response
}

func (n *RpcClient) GetUnspents(adddress string) GetUnspentsResponse {
	response := GetUnspentsResponse{}
	params := []interface{}{adddress}
	n.makeRequest("getunspents", params, &response)
	return response
}
