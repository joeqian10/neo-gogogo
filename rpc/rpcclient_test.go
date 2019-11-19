package rpc

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestNewClient(t *testing.T) {
	rpcClient := NewClient("http://seed1.ngd.network:20332")
	assert.NotNil(t, rpcClient)
	endpoint := rpcClient.Endpoint
	assert.NotNil(t, endpoint)
	assert.Equal(t, "seed1.ngd.network:20332", endpoint.Host)
	assert.Equal(t, "http", endpoint.Scheme)
}

func TestRpcClient_ClaimGas(t *testing.T) {
	var client = new(ClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "txid": "0xa6d16b98dbd4b2d140bd8316f595de3c6770456454d5aa48e1d3dbe11c1acd3e",
        "size": 543,
        "type": "ClaimTransaction",
        "version": 0,
        "attributes": [],
        "vin": [],
        "vout": [
            {
                "n": 0,
                "asset": "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
                "value": "0.499384",
                "address": "AH2TGXkKgiWm4xMQzVT5R9zV1yxwVXNAPT"
            }
        ],
        "sys_fee": "0",
        "net_fee": "0",
        "scripts": [
            {
                "invocation": "4000b71a53738761ec03d9cd5d85fc8ea2acf42cda076e7a5f690adb897b2ee8ed018a0000494d6ee047cd66ec926806a0c934fe1a84aec171f66eedf8309d6a9a",
                "verification": "2102bac395577bf47e7dd6058b7dfc38d2b351966ac03ea84ba7bca410e8a8aded20ac"
            },
            {
                "invocation": "4086cdf2622d901110f2c04e646f8e78f4c995f3be399bbe5d2d4ac81918fd4a5d5c39a7df60f827c00de1f19e3584c3253541c52b88bed54ba1b9a448548b0e5d",
                "verification": "2103991949671a85ba5eb09a982c94e0d205b81c94f958109895b4ebafa747caaf09ac"
            }
        ],
        "claims": [
            {
                "txid": "0xe1e44f41a1f0854063ccdc9beb7537fc40565575e0ae2366b4a93a73c18b6166",
                "vout": 2
            },
            {
                "txid": "0x69fd452c92fb0e5861b27588549a8e55d3f9fee542884ae317600508bbacedbb",
                "vout": 2
            },
            {
                "txid": "0x152d823d5cf1ce58cf33879e23309dc83152cfb8c50ba05cc03c090dcd00198e",
                "vout": 0
            },
            {
                "txid": "0x152d823d5cf1ce58cf33879e23309dc83152cfb8c50ba05cc03c090dcd00198e",
                "vout": 1
            },
            {
                "txid": "0xe1e44f41a1f0854063ccdc9beb7537fc40565575e0ae2366b4a93a73c18b6166",
                "vout": 0
            },
            {
                "txid": "0xe1e44f41a1f0854063ccdc9beb7537fc40565575e0ae2366b4a93a73c18b6166",
                "vout": 1
            },
            {
                "txid": "0x668e9be5185e1cfa1efb08b673062038ce04ebc9db41f75dc74d6faacbaf71ea",
                "vout": 0
            },
            {
                "txid": "0x668e9be5185e1cfa1efb08b673062038ce04ebc9db41f75dc74d6faacbaf71ea",
                "vout": 1
            }
        ]
    }
}`))),
	}, nil)

	response := rpc.ClaimGas("")
	r := response.Result
	assert.Equal(t, "ClaimTransaction", r.Type)
	assert.Equal(t, 0, r.Version)
}

func TestRpcClient_GetAccountState(t *testing.T) {
	var client = new(ClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "version": 0,
        "script_hash": "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a",
        "frozen": false,
        "votes": [],
        "balances": [
            {
                "asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
                "value": "94"
            }
        ]
    }
}`))),
	}, nil)

	response := rpc.GetAccountState("")
	r := response.Result

	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
}
//
//func TestRpcClient_GetApplicationLog(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetAssetState(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBalance(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBestBlockHash(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBlockByHash(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBlockByIndex(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBlockCount(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBlockHash(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetBlockHeaderByHash(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetClaimable(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetConnectionCount(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetContractState(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetNep5Balances(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetNep5Transfers(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetNewAddress(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetPeers(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetRawMemPool(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetRawTransaction(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetStorage(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetTransactionHeight(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetTxOut(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetUnclaimed(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetUnclaimedGas(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetUnspents(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetValidators(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetVersion(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_GetWalletHeight(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_ImportPrivKey(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_InvokeFunction(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_InvokeScript(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_ListAddress(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_ListPlugins(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_SendFrom(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_SendRawTransaction(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_SendToAddress(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_SubmitBlock(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_ValidateAddress(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
//
//func TestRpcClient_makeRequest(t *testing.T) {
//	var client = new(ClientMock)
//	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
//	client.On("Do", mock.Anything).Return(&http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(``))),
//	}, nil)
//
//	response := rpc.GetAccountState("")
//	r := response.Result
//
//	assert.Equal(t, 0, r.Version)
//	assert.Equal(t, "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a", r.ScriptHash)
//}
