package rpc

import (
	"bytes"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type HttpClientMock struct {
	mock.Mock
}

func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
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
	var client = new(HttpClientMock)
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
	var client = new(HttpClientMock)
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

func TestRpcClient_GetApplicationLog(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"txid": "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a",
				"executions": [
					{
						"trigger": "Application",
						"contract": "0x003bd113b3bc841657f3a84db8546daa6e4953c3",
						"vmstate": "HALT",
						"gas_consumed": "2.855",
						"stack": [
							{
								"type": "Integer",
								"value": "1"
							}
						],
						"notifications": [
							{
								"contract": "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263",
								"state": {
									"type": "Array",
									"value": [
										{
											"type": "ByteArray",
											"value": "7472616e73666572"
										},
										{
											"type": "ByteArray",
											"value": "5c564ab204122ddce30eb9a6accbfa23b27cc3ac"
										},
										{
											"type": "ByteArray",
											"value": "8f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc"
										},
										{
											"type": "ByteArray",
											"value": "00203d88792d"
										}
									]
								}
							}
						]
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetApplicationLog("")
	r := response.Result
	assert.Equal(t, "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a", r.TxId)
	e := r.Executions[0]
	assert.Equal(t, "Application", e.Trigger)
	assert.Equal(t, "0x003bd113b3bc841657f3a84db8546daa6e4953c3", e.Contract)
	assert.Equal(t, "HALT", e.VMState)
	n := e.Notifications[0]
	assert.Equal(t, "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", n.Contract)
}

func TestRpcClient_GetAssetState(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"version": 0,
				"id": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
				"type": "GoverningToken",
				"name": [
					{
						"lang": "zh-CN",
						"name": "小蚁股"
					},
					{
						"lang": "en",
						"name": "AntShare"
					}
				],
				"amount": "100000000",
				"available": "100000000",
				"precision": 0,
				"owner": "00",
				"admin": "Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt",
				"issuer": "Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt",
				"expiration": 4000000,
				"frozen": false
			}
		}`))),
	}, nil)

	response := rpc.GetAssetState("")
	r := response.Result
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Id)
	assert.Equal(t, "GoverningToken", r.Type)
	assert.Equal(t, "AntShare", r.Name[1].Name)
	assert.Equal(t, "100000000", r.Amount)
	assert.Equal(t, 0, r.Precision)
	assert.Equal(t, "Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt", r.Admin)
	assert.Equal(t, 4000000, r.Expiration)
}

func TestRpcClient_GetBalance(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"balance": "100000000",
				"confirmed": "100000000"
			}
		}`))),
	}, nil)

	response := rpc.GetBalance("")
	r := response.Result
	assert.Equal(t, "100000000", r.Balance)
}

func TestRpcClient_GetBestBlockHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f"
		}`))),
	}, nil)

	response := rpc.GetBestBlockHash()
	r := response.Result
	assert.Equal(t, "0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f", r)
}

func TestRpcClient_GetBlockByHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179",
				"size": 1521,
				"version": 0,
				"previousblockhash": "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840",
				"merkleroot": "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757",
				"time": 1573123342,
				"index": 3386365,
				"nonce": "a6e6d82b50273b82",
				"nextconsensus": "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X",
				"script": {
					"invocation": "40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd",
					"verification": "552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae"
				},
				"tx": [
					{
						"txid": "0x75f1de0f6aaab785138fb8a5183018d25465278eb38c44917db87b61a7f1c588",
						"size": 10,
						"type": "MinerTransaction",
						"version": 0,
						"attributes": [],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [],
						"nonce": 1344748418
					},
					{
						"txid": "0x18147d0916e1f2fbbcc26a3bb5fd593b90ea86c8a27be4496eebbccb8fe99de9",
						"size": 247,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e45752e389b187108"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40f24367766e1fc0e1a40a9565d7395d5661d29f058668d1374077c4d8a217de9c8848509785f79eda1be3afb8881a29993aeb973963281638277e68ce64b822f7",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "1410d46912932d6ebcd1d3c4a27a1a8ea77e68ac950020000101001a000010020000000056054b2042000019c5420830b10420500444bb53c11063726561746550726f6d6f437574696567f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x452d52a1e8963746e80ea92a45660d85aeb9d6cd04e0065439fc95270db1810d",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e457520059c8d6488"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40da7201b8a684d32a1c6d5a4665ef5f99ef50cfab2e6b42eb9972fa8edbeef8646eb46bb1d0caa74080d898f5affb3f6d67e445691e0aa5f0af71994f5976cebf",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "00029d0652c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x45ea381e940f0e089a72f07f1a05d4bfa2be85410e96eea472c55e7b11872d5a",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e45751f72959c1668"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40ae6931e6fa6398e0ccb448419debd96b2ce7ef76d4f4a172ada527f3da28d2fc77f2ea2aac04186c6ad575cd17d117820ee6d3395f1ebfb900b29972b69d20ac",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "0002460552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x82ad53594683ca5fc44e658fc1ed265b86070e4466417df3f9eadf5e207b87b2",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e457520981613af83"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "404e82e563c8a87155bcfa805e02f45f293d2568632a6aebb8eef6f4a4e51328b43aba45e363fac26461ac821bdc633e4ae9f5836b628820f8b382924c213b6a53",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "0002800552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					}
				],
				"confirmations": 60576,
				"nextblockhash": "0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac"
			}
		}`))),
	}, nil)

	response := rpc.GetBlockByHash("")
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 1521, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.PreviousBlockHash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.MerkleRoot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.NextConsensus)
}

func TestRpcClient_GetBlockByIndex(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179",
				"size": 1521,
				"version": 0,
				"previousblockhash": "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840",
				"merkleroot": "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757",
				"time": 1573123342,
				"index": 3386365,
				"nonce": "a6e6d82b50273b82",
				"nextconsensus": "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X",
				"script": {
					"invocation": "40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd",
					"verification": "552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae"
				},
				"tx": [
					{
						"txid": "0x75f1de0f6aaab785138fb8a5183018d25465278eb38c44917db87b61a7f1c588",
						"size": 10,
						"type": "MinerTransaction",
						"version": 0,
						"attributes": [],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [],
						"nonce": 1344748418
					},
					{
						"txid": "0x18147d0916e1f2fbbcc26a3bb5fd593b90ea86c8a27be4496eebbccb8fe99de9",
						"size": 247,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e45752e389b187108"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40f24367766e1fc0e1a40a9565d7395d5661d29f058668d1374077c4d8a217de9c8848509785f79eda1be3afb8881a29993aeb973963281638277e68ce64b822f7",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "1410d46912932d6ebcd1d3c4a27a1a8ea77e68ac950020000101001a000010020000000056054b2042000019c5420830b10420500444bb53c11063726561746550726f6d6f437574696567f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x452d52a1e8963746e80ea92a45660d85aeb9d6cd04e0065439fc95270db1810d",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e457520059c8d6488"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40da7201b8a684d32a1c6d5a4665ef5f99ef50cfab2e6b42eb9972fa8edbeef8646eb46bb1d0caa74080d898f5affb3f6d67e445691e0aa5f0af71994f5976cebf",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "00029d0652c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x45ea381e940f0e089a72f07f1a05d4bfa2be85410e96eea472c55e7b11872d5a",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e45751f72959c1668"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "40ae6931e6fa6398e0ccb448419debd96b2ce7ef76d4f4a172ada527f3da28d2fc77f2ea2aac04186c6ad575cd17d117820ee6d3395f1ebfb900b29972b69d20ac",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "0002460552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					},
					{
						"txid": "0x82ad53594683ca5fc44e658fc1ed265b86070e4466417df3f9eadf5e207b87b2",
						"size": 196,
						"type": "InvocationTransaction",
						"version": 1,
						"attributes": [
							{
								"usage": "Script",
								"data": "10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95"
							},
							{
								"usage": "Remark",
								"data": "0000016e457520981613af83"
							}
						],
						"vin": [],
						"vout": [],
						"sys_fee": "0",
						"net_fee": "0",
						"scripts": [
							{
								"invocation": "404e82e563c8a87155bcfa805e02f45f293d2568632a6aebb8eef6f4a4e51328b43aba45e363fac26461ac821bdc633e4ae9f5836b628820f8b382924c213b6a53",
								"verification": "2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac"
							}
						],
						"script": "0002800552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3",
						"gas": "0"
					}
				],
				"confirmations": 60576,
				"nextblockhash": "0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac"
			}
		}`))),
	}, nil)

	response := rpc.GetBlockByIndex(0)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 1521, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.PreviousBlockHash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.MerkleRoot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.NextConsensus)
}

func TestRpcClient_GetBlockCount(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 2023
		}`))),
	}, nil)

	response := rpc.GetBlockCount()
	r := response.Result
	assert.Equal(t, 2023, r)
}

func TestRpcClient_GetBlockHeaderByHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179",
				"size": 676,
				"version": 0,
				"previousblockhash": "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840",
				"merkleroot": "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757",
				"time": 1573123342,
				"index": 3386365,
				"nonce": "a6e6d82b50273b82",
				"nextconsensus": "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X",
				"script": {
					"invocation": "40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd",
					"verification": "552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae"
				},
				"confirmations": 60656,
				"nextblockhash": "0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac"
			}
		}`))),
	}, nil)

	response := rpc.GetBlockHeaderByHash("")
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 676, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.PreviousBlockHash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.MerkleRoot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.NextConsensus)
}

func TestRpcClient_GetBlockHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179"
		}`))),
	}, nil)

	response := rpc.GetBlockHash(0)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r)
}

func TestRpcClient_GetClaimable(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"claimable": [
					{
						"txid": "52ba70ef18e879785572c917795cd81422c3820b8cf44c24846a30ee7376fd77",
						"n": 1,
						"value": 800000,
						"start_height": 476496,
						"end_height": 488154,
						"generated": 746.112,
						"sys_fee": 3.92,
						"unclaimed": 750.032
					}
				],
				"address": "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
				"unclaimed": 750.032
			}
		}`))),
	}, nil)

	response := rpc.GetClaimable("")
	r := response.Result
	c := r.Claimables[0]
	assert.Equal(t, "52ba70ef18e879785572c917795cd81422c3820b8cf44c24846a30ee7376fd77", c.TxId)
	assert.Equal(t, 1, c.N)
	assert.Equal(t, "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt", r.Address)
	assert.Equal(t, 750.032, r.Unclaimed)
}

func TestRpcClient_GetConnectionCount(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 48
		}`))),
	}, nil)

	response := rpc.GetConnectionCount()
	r := response.Result
	assert.Equal(t, 48, r)
}

func TestRpcClient_GetContractState(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"version": 0,
				"hash": "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263",
				"script": "011ac56b6c766b00527ac46c766b51527ac4616168164e656f2e52756e74696d652e47657454726967676572009c6c766b54527ac46c766b54c3643e0061145c564ab204122ddce30eb9a6accbfa23b27cc3ac6168184e656f2e52756e74696d652e436865636b5769746e6573736c766b55527ac46259036168164e656f2e52756e74696d652e47657454726967676572609c6c766b56527ac46c766b56c364b702616c766b00c304696e6974876c766b57527ac46c766b57c364110061654e036c766b55527ac46206036c766b00c30a6d696e74546f6b656e73876c766b58527ac46c766b58c36411006165df066c766b55527ac462d8026c766b00c30b746f74616c537570706c79876c766b59527ac46c766b59c3641100616537096c766b55527ac462a9026c766b00c3046e616d65876c766b5a527ac46c766b5ac3641100616594026c766b55527ac46281026c766b00c30673796d626f6c876c766b5b527ac46c766b5bc364110061657d026c766b55527ac46257026c766b00c30a746f74616c546f6b656e876c766b5c527ac46c766b5cc3641100616562026c766b55527ac46229026c766b00c30869636f546f6b656e876c766b5d527ac46c766b5dc364110061658f036c766b55527ac462fd016c766b00c30669636f4e656f876c766b5e527ac46c766b5ec36411006165ca036c766b55527ac462d3016c766b00c306656e6449636f876c766b5f527ac46c766b5fc364110061650c046c766b55527ac462a9016c766b00c3087472616e73666572876c766b60527ac46c766b60c3647900616c766b51c3c0539c009c6c766b0114527ac46c766b0114c3640e00006c766b55527ac46264016c766b51c300c36c766b0111527ac46c766b51c351c36c766b0112527ac46c766b51c352c36c766b0113527ac46c766b0111c36c766b0112c36c766b0113c361527265f2076c766b55527ac46215016c766b00c30962616c616e63654f66876c766b0115527ac46c766b0115c3644d00616c766b51c3c0519c009c6c766b0117527ac46c766b0117c3640e00006c766b55527ac462cd006c766b51c300c36c766b0116527ac46c766b0116c36165cb096c766b55527ac462aa006c766b00c308646563696d616c73876c766b0118527ac46c766b0118c36411006165ad006c766b55527ac4627c00616165990b6c766b52527ac461650d0d6c766b53527ac46c766b53c300907c907ca1630e006c766b52c3c000a0620400006c766b0119527ac46c766b0119c3642f00616c766b52c36c766b53c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f746966796161006c766b55527ac46203006c766b55c3616c756600c56b0b516c696e6b20546f6b656e616c756600c56b03514c43616c756600c56b58616c756600c56b07008048efefd801616c756653c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c3c000a06c766b51527ac46c766b51c3640e00006c766b52527ac462df006168164e656f2e53746f726167652e476574436f6e74657874145c564ab204122ddce30eb9a6accbfa23b27cc3ac0800803dafbe50d300615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c790800803dafbe50d300615272680f4e656f2e53746f726167652e5075746100145c564ab204122ddce30eb9a6accbfa23b27cc3ac0800803dafbe50d300615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b52527ac46203006c766b52c3616c756652c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c30800803dafbe50d300946c766b51527ac46203006c766b51c3616c756652c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c30800803dafbe50d30094050008711b0c966c766b51527ac46203006c766b51c3616c756655c56b61145c564ab204122ddce30eb9a6accbfa23b27cc3ac6168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b52527ac46c766b52c3640e00006c766b53527ac4624e016168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac4080000869eae29d5006c766b00c3946c766b51527ac46c766b51c300a16c766b54527ac46c766b54c3640f0061006c766b53527ac462d7006168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79080000869eae29d500615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e74657874145c564ab204122ddce30eb9a6accbfa23b27cc3ac6c766b51c3615272680f4e656f2e53746f726167652e5075746100145c564ab204122ddce30eb9a6accbfa23b27cc3ac6c766b51c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b53527ac46203006c766b53c3616c75665cc56b61616520076c766b00527ac46c766b00c3c0009c6c766b58527ac46c766b58c3640f0061006c766b59527ac4624f026168184e656f2e426c6f636b636861696e2e4765744865696768746168184e656f2e426c6f636b636861696e2e4765744865616465726168174e656f2e4865616465722e47657454696d657374616d706c766b51527ac46c766b51c304d0013d5a946c766b52527ac46c766b52c36c766b00c3617c656e096c766b53527ac46c766b52c36165b2046c766b54527ac46c766b54c3009c6c766b5a527ac46c766b5ac3643900616c766b00c36c766b53c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b59527ac46274016c766b00c36c766b53c36c766b54c361527265b8046c766b55527ac46c766b55c3009c6c766b5b527ac46c766b5bc3640f0061006c766b59527ac46236016168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b00c36c766b55c36c766b56c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c796c766b55c36c766b57c393615272680f4e656f2e53746f726167652e50757461006c766b00c36c766b55c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b59527ac46203006c766b59c3616c756651c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46203006c766b00c3616c75665bc56b6c766b00527ac46c766b51527ac46c766b52527ac4616c766b52c300a16c766b55527ac46c766b55c3640e00006c766b56527ac46204026c766b00c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b57527ac46c766b57c3640e00006c766b56527ac462c8016c766b00c36c766b51c39c6c766b58527ac46c766b58c3640e00516c766b56527ac462a3016168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b53527ac46c766b53c36c766b52c39f6c766b59527ac46c766b59c3640e00006c766b56527ac46246016c766b53c36c766b52c39c6c766b5a527ac46c766b5ac3643b006168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c68124e656f2e53746f726167652e44656c657465616241006168164e656f2e53746f726167652e476574436f6e746578746c766b00c36c766b53c36c766b52c394615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b51c3617c680f4e656f2e53746f726167652e4765746c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b51c36c766b54c36c766b52c393615272680f4e656f2e53746f726167652e507574616c766b00c36c766b51c36c766b52c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b56527ac46203006c766b56c3616c756652c56b6c766b00527ac4616168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b51527ac46203006c766b51c3616c756654c56b6c766b00527ac4616c766b00c3009f6c766b51527ac46c766b51c3640f0061006c766b52527ac4623b006c766b00c30380de28a0009c6c766b53527ac46c766b53c364140061050008711b0c6c766b52527ac4620f0061006c766b52527ac46203006c766b52c3616c756659c56b6c766b00527ac46c766b51527ac46c766b52527ac4616c766b51c30400e1f505966c766b52c3956c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b54527ac4080000869eae29d5006c766b54c3946c766b55527ac46c766b55c300a16c766b56527ac46c766b56c3643900616c766b00c36c766b51c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b57527ac46276006c766b55c36c766b53c39f6c766b58527ac46c766b58c3644d00616c766b00c36c766b53c36c766b55c3946c766b52c3960400e1f50595617c06726566756e6453c168124e656f2e52756e74696d652e4e6f74696679616c766b55c36c766b53527ac4616c766b53c36c766b57527ac46203006c766b57c3616c756657c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b00527ac46c766b00c361681d4e656f2e5472616e73616374696f6e2e4765745265666572656e6365736c766b51527ac4616c766b51c36c766b52527ac4006c766b53527ac4629d006c766b52c36c766b53c3c36c766b54527ac4616c766b54c36168154e656f2e4f75747075742e47657441737365744964209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc59c6c766b55527ac46c766b55c3642d006c766b54c36168184e656f2e4f75747075742e476574536372697074486173686c766b56527ac4622c00616c766b53c351936c766b53527ac46c766b53c36c766b52c3c09f635aff006c766b56527ac46203006c766b56c3616c756651c56b6161682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b00527ac46203006c766b00c3616c756658c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b00527ac46c766b00c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b51527ac4006c766b52527ac4616c766b51c36c766b53527ac4006c766b54527ac462cd006c766b53c36c766b54c3c36c766b55527ac4616c766b55c36168184e656f2e4f75747075742e47657453637269707448617368616505ff907c907c9e6345006c766b55c36168154e656f2e4f75747075742e47657441737365744964209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc59c620400006c766b56527ac46c766b56c3642d00616c766b52c36c766b55c36168134e656f2e4f75747075742e47657456616c7565936c766b52527ac461616c766b54c351936c766b54527ac46c766b54c36c766b53c3c09f632aff6c766b52c36c766b57527ac46203006c766b57c3616c75665ac56b6c766b00527ac46c766b51527ac46161657cfe6c766b52527ac46c766b00c300a16311006c766b00c3026054a0009c620400006c766b53527ac46c766b53c364870161067072656669786c766b51c37e6c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac404002f68596c766b55c3946c766b56527ac46c766b56c300a0009c6c766b57527ac46c766b57c3643900616c766b51c36c766b52c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b58527ac462e9006c766b56c36c766b52c39f6c766b59527ac46c766b59c3648100616168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b56c36c766b55c393615272680f4e656f2e53746f726167652e507574616c766b51c36c766b52c36c766b56c394617c06726566756e6453c168124e656f2e52756e74696d652e4e6f74696679616c766b56c36c766b58527ac46251006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b52c36c766b55c393615272680f4e656f2e53746f726167652e50757461616c766b52c36c766b58527ac46203006c766b58c3616c7566",
				"parameters": [
					"String",
					"Array"
				],
				"returntype": "ByteArray",
				"name": "QLC",
				"code_version": "1.0",
				"author": "qlink",
				"email": "admin@qlink.mobi",
				"description": "qlink token\t\t",
				"properties": {
					"storage": true,
					"dynamic_invoke": false
				}
			}
		}`))),
	}, nil)

	response := rpc.GetContractState("")
	r := response.Result
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", r.Hash)
	assert.Equal(t, "QLC", r.Name)
}

func TestRpcClient_GetNep5Balances(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"balance": [
					{
						"asset_hash": "b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263",
						"amount": "50000000000000",
						"last_updated_block": 3275044
					}
				],
				"address": "AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY"
			}
		}`))),
	}, nil)

	response := rpc.GetNep5Balances("")
	r := response.Result
	assert.Equal(t, "b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", r.Balances[0].AssetHash)
	assert.Equal(t, "AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY", r.Address)
}

func TestRpcClient_GetNep5Transfers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"sent": [],
				"received": [
					{
						"timestamp": 1555651816,
						"asset_hash": "600c4f5200db36177e3e8a09e9f18e2fc7d12a0f",
						"transfer_address": "AYwgBNMepiv5ocGcyNT4mA8zPLTQ8pDBis",
						"amount": "1000000",
						"block_index": 436036,
						"transfer_notify_index": 0,
						"tx_hash": "df7683ece554ecfb85cf41492c5f143215dd43ef9ec61181a28f922da06aba58"
					}
				],
				"address": "AbHgdBaWEnHkCiLtDZXjhvhaAK2cwFh5pF"
			}
		}`))),
	}, nil)

	response := rpc.GetNep5Transfers("")
	r := response.Result
	c := r.Received[0]
	assert.Equal(t, "AbHgdBaWEnHkCiLtDZXjhvhaAK2cwFh5pF", r.Address)
	assert.Equal(t, 1555651816, c.Timestamp)
}

func TestRpcClient_GetNewAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "AVfppvZcLeynVzsCe5W52yH925rXZARWVd"
		}`))),
	}, nil)

	response := rpc.GetNewAddress()
	r := response.Result
	assert.Equal(t, "AVfppvZcLeynVzsCe5W52yH925rXZARWVd", r)
}

func TestRpcClient_GetPeers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"unconnected": [
					{
						"address": "54.65.248.144",
						"port": 20333
					},
					{
						"address": "104.196.172.132",
						"port": 20333
					},
					{
						"address": "47.254.83.14",
						"port": 20333
					}
				],
				"bad": [],
				"connected": [
					{
						"address": "34.83.126.133",
						"port": 20333
					},
					{
						"address": "18.222.168.189",
						"port": 10333
					},
					{
						"address": "47.254.83.14",
						"port": 20333
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetPeers()
	r := response.Result
	u := r.Connected[0]
	assert.Equal(t, "34.83.126.133", u.Address)
	assert.Equal(t, 20333, u.Port)
}

func TestRpcClient_GetRawMemPool(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				"0x9786cce0dddb524c40ddbdd5e31a41ed1f6b5c8a683c122f627ca4a007a7cf4e",
				"0xb488ad25eb474f89d5ca3f985cc047ca96bc7373a6d3da8c0f192722896c1cd7",
				"0xf86f6f2c08fbf766ebe59dc84bc3b8829f1053f0a01deb26bf7960d99fa86cd6"
			]
		}`))),
	}, nil)

	response := rpc.GetRawMemPool()
	r := response.Result
	assert.Equal(t, "0x9786cce0dddb524c40ddbdd5e31a41ed1f6b5c8a683c122f627ca4a007a7cf4e", r[0])
}

func TestRpcClient_GetRawTransaction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"txid": "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a",
				"size": 242,
				"type": "InvocationTransaction",
				"version": 1,
				"attributes": [
					{
						"usage": "Script",
						"data": "5c564ab204122ddce30eb9a6accbfa23b27cc3ac"
					},
					{
						"usage": "Remark",
						"data": "313537313231383636323935373964363035643631"
					}
				],
				"vin": [],
				"vout": [],
				"sys_fee": "0",
				"net_fee": "0",
				"scripts": [
					{
						"invocation": "4002a84056e9bf04ed47a6307c3030ac92704cb71a8c2fd46f45593c8ce57a403de47e19a4171114e7ec881d9f45d7851712e8eb922d11ce3a0de5ea64b8310025",
						"verification": "2103f19ffa8acecb480ab727b0bf9ee934162f6e2a4308b59c80b732529ebce6f53dac"
					}
				],
				"script": "0600203d88792d148f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc145c564ab204122ddce30eb9a6accbfa23b27cc3ac53c1087472616e736665726763d26113bac4208254d98a3eebaee66230ead7b9",
				"gas": "0",
				"blockhash": "0x1a1d7b2f6d54e7c9084353372dd526301a456900827ce8478fcff1a7a00766f7",
				"confirmations": 172117,
				"blocktime": 1571218675
			}
		}`))),
	}, nil)

	response := rpc.GetRawTransaction("")
	r := response.Result
	assert.Equal(t, "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a", r.Txid)
	assert.Equal(t, 242, r.Size)
	assert.Equal(t, "InvocationTransaction", r.Type)
	assert.Equal(t, 1, r.Version)
}

func TestRpcClient_GetStorage(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "00e1f505"
		}`))),
	}, nil)

	response := rpc.GetStorage("", "")
	r := response.Result
	assert.Equal(t, "00e1f505", r)
}

func TestRpcClient_GetTransactionHeight(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 3275044
		}`))),
	}, nil)

	response := rpc.GetTransactionHeight("")
	r := response.Result
	assert.Equal(t, 3275044, r)
}

func TestRpcClient_GetTxOut(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"n": 0,
				"asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
				"value": "1000000",
				"address": "AazAnhssUfNyBC3rdBKseGuck7voaF5p68"
			}
		}`))),
	}, nil)

	response := rpc.GetTxOut("", 0)
	r := response.Result
	assert.Equal(t, 0, r.N)
	assert.Equal(t, "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Asset)
	assert.Equal(t, "1000000", r.Value)
	assert.Equal(t, "AazAnhssUfNyBC3rdBKseGuck7voaF5p68", r.Address)
}

func TestRpcClient_GetUnclaimed(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"available": 4104,
				"unavailable": 4096,
				"unclaimed": 0
			}
		}`))),
	}, nil)

	response := rpc.GetUnclaimed("")
	r := response.Result
	assert.Equal(t, float64(4104), r.Available)
	assert.Equal(t, float64(4096), r.Unavailable)
}

func TestRpcClient_GetUnclaimedGas(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"available": "4104",
				"unavailable": "4096"
			}
		}`))),
	}, nil)

	response := rpc.GetUnclaimedGas()
	r := response.Result
	assert.Equal(t, "4104", r.Available)
	assert.Equal(t, "4096", r.Unavailable)
}

func TestRpcClient_GetUnspents(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"balance": [
					{
						"unspent": [
							{
								"txid": "4ee4af75d5aa60598fbae40ce86fb9a23ffec5a75dfa8b59d259d15f9e304319",
								"n": 0,
								"value": 27844.821
							},
							{
								"txid": "9906bf2a9f531ac523aad5e9507bd6540acc1c65ae9144918ccc891188578253",
								"n": 0,
								"value": 0.987
							},
							{
								"txid": "184e34eb3f9550d07d03563391d73eb6c438130c7fdca37f0700d5d52ad7deb1",
								"n": 0,
								"value": 243.95598
							},
							{
								"txid": "448abc64412284fb21c9625ac9edd2100090367a551c18ce546c1eded61e77c3",
								"n": 0,
								"value": 369.84904
							},
							{
								"txid": "bd454059e58da4221aaf4effa3278660b231e9af7cea97912f4ac5c4995bb7e4",
								"n": 0,
								"value": 600.41014479
							}
						],
						"asset_hash": "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
						"asset": "GAS",
						"asset_symbol": "GAS",
						"amount": 29060.02316479
					},
					{
						"unspent": [
							{
								"txid": "c3182952855314b3f4b1ecf01a03b891d4627d19426ce841275f6d4c186e729a",
								"n": 0,
								"value": 800000
							}
						],
						"asset_hash": "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
						"asset": "NEO",
						"asset_symbol": "NEO",
						"amount": 800000
					}
				],
				"address": "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt"
			}
		}`))),
	}, nil)

	response := rpc.GetUnspents("")
	r := response.Result
	assert.Equal(t, "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7", r.Balances[0].AssetHash)
	assert.Equal(t, "GAS", r.Balances[0].Asset)
	assert.Equal(t, "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Balances[1].AssetHash)
	assert.Equal(t, "NEO", r.Balances[1].Asset)
}

func TestRpcClient_GetValidators(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"publickey": "025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64",
					"votes": "90101602",
					"active": true
				},
				{
					"publickey": "0266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af07",
					"votes": "90101602",
					"active": true
				},
				{
					"publickey": "03028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec133055652",
					"votes": "90101602",
					"active": true
				},
				{
					"publickey": "030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e36",
					"votes": "90100000",
					"active": true
				},
				{
					"publickey": "03b8bfe058dc404a2f9510606aee0de69e3d8b47e25d7b3af670577373640d51dc",
					"votes": "1000",
					"active": false
				},
				{
					"publickey": "03c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c1",
					"votes": "90100000",
					"active": true
				},
				{
					"publickey": "03fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f",
					"votes": "90100000",
					"active": true
				},
				{
					"publickey": "03fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df9",
					"votes": "90100000",
					"active": true
				}
			]
		}`))),
	}, nil)

	response := rpc.GetValidators()
	r := response.Result
	v := r[0]
	assert.Equal(t, "025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64", v.PublicKey)
	assert.Equal(t, "90101602", v.Votes)
	assert.Equal(t, true, v.Active)
}

func TestRpcClient_GetVersion(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"port": 20333,
				"nonce": 2013624361,
				"useragent": "/Neo:2.10.3/"
			}
		}`))),
	}, nil)

	response := rpc.GetVersion()
	r := response.Result
	assert.Equal(t, 20333, r.Port)
	assert.Equal(t, 2013624361, r.Nonce)
	assert.Equal(t, "/Neo:2.10.3/", r.Useragent)
}

func TestRpcClient_GetWalletHeight(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 2894
		}`))),
	}, nil)

	response := rpc.GetWalletHeight()
	r := response.Result
	assert.Equal(t, 2894, r)
}

func TestRpcClient_ImportPrivKey(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "Ad8S24trcuchyLfEbJWqRP7BUScUT4t2pw",
				"haskey": true,
				"label": null,
				"watchonly": false
			}
		}`))),
	}, nil)

	response := rpc.ImportPrivKey("")
	r := response.Result
	assert.Equal(t, "Ad8S24trcuchyLfEbJWqRP7BUScUT4t2pw", r.Address)
	assert.Equal(t, true, r.HasKey)
}

func TestRpcClient_InvokeFunction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
				"state": "HALT",
				"gas_consumed": "0.126",
				"stack": [
					{
						"type": "ByteArray",
						"value": "516c696e6b20546f6b656e"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.InvokeFunction("", "", "")
	r := response.Result
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", r.Script)
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "516c696e6b20546f6b656e", r.Stack[0].Value)
}

func TestRpcClient_InvokeScript(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
				"state": "HALT",
				"gas_consumed": "0.126",
				"stack": [
					{
						"type": "ByteArray",
						"value": "516c696e6b20546f6b656e"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.InvokeScript("", "")
	r := response.Result
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", r.Script)
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "516c696e6b20546f6b656e", r.Stack[0].Value)
}

func TestRpcClient_InvokeScript2(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
				"state": "HALT",
				"gas_consumed": "0.229",
				"stack": [
					{
						"type": "Array",
						"value": [
							{
								"type": "ByteArray",
								"value": "90aaf35d"
							},
							{
								"type": "Integer",
								"value": "1607668937"
							}
						]
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.InvokeScript("", "")
	r := response.Result
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", r.Script)
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "Array", r.Stack[0].Type)
	r.Stack[0].Convert()
	var a = r.Stack[0].Value.([]models.InvokeStack)
	assert.Equal(t, "90aaf35d", a[0].Value)
}

func TestRpcClient_ListAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"address": "ASL3KCvJasA7QzpYGePp25pWuQCj4dd9Sy",
					"haskey": true,
					"label": null,
					"watchonly": false
				},
				{
					"address": "AV2Ai7PXcNbjTSeKgWqsDEjLaEAJZpytru",
					"haskey": true,
					"label": null,
					"watchonly": false
				}
			]
		}`))),
	}, nil)

	response := rpc.ListAddress()
	r := response.Result
	a := r[0]
	assert.Equal(t, "ASL3KCvJasA7QzpYGePp25pWuQCj4dd9Sy", a.Address)
	assert.Equal(t, true, a.HasKey)
	assert.Equal(t, false, a.WatchOnly)
}

func TestRpcClient_ListPlugins(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"name": "ApplicationLogs",
					"version": "2.10.3.0",
					"interfaces": [
						"IRpcPlugin",
						"IPersistencePlugin"
					]
				},
				{
					"name": "ImportBlocks",
					"version": "2.10.3.0",
					"interfaces": []
				},
				{
					"name": "RpcNep5Tracker",
					"version": "2.10.3.0",
					"interfaces": [
						"IPersistencePlugin",
						"IRpcPlugin"
					]
				},
				{
					"name": "RpcSecurity",
					"version": "2.10.3.0",
					"interfaces": [
						"IRpcPlugin"
					]
				},
				{
					"name": "RpcSystemAssetTrackerPlugin",
					"version": "2.10.3.0",
					"interfaces": [
						"IPersistencePlugin",
						"IRpcPlugin"
					]
				},
				{
					"name": "RpcWallet",
					"version": "2.10.3.0",
					"interfaces": [
						"IRpcPlugin"
					]
				},
				{
					"name": "SimplePolicyPlugin",
					"version": "2.10.3.0",
					"interfaces": [
						"ILogPlugin",
						"IPolicyPlugin"
					]
				}
			]
		}`))),
	}, nil)

	response := rpc.ListPlugins()
	r := response.Result
	p := r[0]
	assert.Equal(t, "ApplicationLogs", p.Name)
	assert.Equal(t, "2.10.3.0", p.Version)
}

func TestRpcClient_SendFrom(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
		  "jsonrpc": "2.0",
		  "id": 1,
		  "result": {
			"txid": "0x60170ad03627ce45c7dd56ececbf33b26eab0845aa8b2cbbeecaefc5771b9eb1",
			"size": 262,
			"type": "ContractTransaction",
			"version": 0,
			"attributes": [],
			"vin": [
			  {
				"txid": "0xd2188c1bd454ac883d79826e5c677deedb91cc61ec6d819df48ff4a963873adb",
				"vout": 1
			  }
			],
			"vout": [
			  {
				"n": 0,
				"asset": "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
				"value": "1",
				"address": "AWg3L6W68bFfSS13Tf4rt8CRdG2ktaAjGb"
			  },
			  {
				"n": 1,
				"asset": "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
				"value": "17.4798197",
				"address": "AWg3L6W68bFfSS13Tf4rt8CRdG2ktaAjGb"
			  }
			],
			"sys_fee": "0",
			"net_fee": "0",
			"scripts": [
			  {
				"invocation": "40a8d40e1652d7ad0c7bb59ef8217237037824af54ee5e46f2fd096c44dd46ef27fa7255010e2a8a2166af8a904e13b96bd3ac82e791633685824c35e7f2731e79",
				"verification": "2102883118351f8f47107c83ab634dc7e4ffe29d274e7d3dcf70159c8935ff769bebac"
			  }
			]
		  }
		}`))),
	}, nil)

	response := rpc.SendFrom("", "", "", 0, 0, "")
	r := response.Result
	assert.Equal(t, "0x60170ad03627ce45c7dd56ececbf33b26eab0845aa8b2cbbeecaefc5771b9eb1", r.Txid)
	assert.Equal(t, 262, r.Size)
	assert.Equal(t, "ContractTransaction", r.Type)
}

func TestRpcClient_SendRawTransaction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": true
		}`))),
	}, nil)

	response := rpc.SendRawTransaction("")
	r := response.Result
	assert.Equal(t, true, r)
}

func TestRpcClient_SendToAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"txid": "0x06de043b9b914f04633c580ab02d89ba55556f775118a292adb6803208857c91",
				"size": 262,
				"type": "ContractTransaction",
				"version": 0,
				"attributes": [],
				"vin": [
					{
						"txid": "0x9c20c13f6b05691efbfd7e420b0edf470f8a5ae467e1e7ca7e11243c9b9fc333",
						"vout": 2
					}
				],
				"vout": [
					{
						"n": 0,
						"asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
						"value": "1",
						"address": "AK4if54jXjSiJBs6jkfZjxAastauJtjjse"
					},
					{
						"n": 1,
						"asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
						"value": "497",
						"address": "AK5q8peiC4QKwuZHWX5Dkqhmar1TAGvZBS"
					}
				],
				"sys_fee": "0",
				"net_fee": "0",
				"scripts": [
					{
						"invocation": "4059e40a2040fe43bf8a40230e1f136dcfe7b3ca37d492ac8d6439615f7b88601c8d9b8077cd0e4f8c9f402d10a2782945bfa50e0ed3f57f7cceebd2f792453eb0",
						"verification": "2103cf5ba6a9135f8eaeda771658564a855c1328af6b6808635496a4f51e3d29ac3eac"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.SendToAddress("", "", 0, 0, "")
	r := response.Result
	assert.Equal(t, "0x06de043b9b914f04633c580ab02d89ba55556f775118a292adb6803208857c91", r.Txid)
	assert.Equal(t, 262, r.Size)
	assert.Equal(t, "ContractTransaction", r.Type)
}

func TestRpcClient_SubmitBlock(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": true
		}`))),
	}, nil)

	response := rpc.SubmitBlock("")
	r := response.Result
	assert.Equal(t, true, r)
}

func TestRpcClient_ValidateAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "AQVh2pG732YvtNaxEGkQUei3YA4cvo7d2i",
				"isvalid": true
			}
		}`))),
	}, nil)

	response := rpc.ValidateAddress("")
	r := response.Result
	assert.Equal(t, "AQVh2pG732YvtNaxEGkQUei3YA4cvo7d2i", r.Address)
	assert.Equal(t, true, r.IsValid)
}

//var assets = []string{"8d51e5d75ec9adf8080213e0f310e78b0063f50d",
//	"b6cb731f90cefebbd4f9cedd0cf56bd1e21967f4",
//	"9a9db8a30a80951ec792effb9731af79781177c2",
//	"7ba002bc1dbc918d555f1d466acda1d540332e28",
//	"378b147c06a7737a6a712d7bdcfa0d5bc3ca4d53",
//	"5e529a73fd7dad3b1fed587f8874c0855cd634c5"}
//
//func TestRedis(t *testing.T) {
//	sb := sc.NewScriptBuilder()
//	scriptHash := helper.HexToBytes("dad9fbc914203b99c1dfbaad0d98738d4e19924d") //
//	address := "AGrgyoJR4FKeWCNGRvduSakaDzhZ9qikr9"
//	lyh, _ := helper.AddressToScriptHash(address)
//	log.Printf("lyh: %s", helper.BytesToHex(lyh.Bytes()))
//
//	for _, asset := range assets {
//		cp1 := sc.ContractParameter{
//			Type:  sc.ByteArray,
//			Value: lyh.Bytes(),
//		}
//
//		cp2 := sc.ContractParameter{
//			Type:  sc.ByteArray,
//			Value: helper.HexToBytes(asset), //
//		}
//
//		args := []sc.ContractParameter{cp1, cp2}
//		sb.MakeInvocationScript(scriptHash, "getStakingAmount", args)
//	}
//
//	for _, asset := range assets {
//		cp2 := sc.ContractParameter{
//			Type:  sc.ByteArray,
//			Value: helper.HexToBytes(asset), //
//		}
//		args := []sc.ContractParameter{cp2}
//		sb.MakeInvocationScript(scriptHash, "getCurrentTotalAmount", args)
//	}
//
//	script := sb.ToArray()
//
//	log.Printf("script: %s", helper.BytesToHex(script))
//
//	//client := NewClient("http://seed10.ngd.network:11332")
//	client := NewClient("https://wallet.ngd.network:10331")
//	checkWitnessHashes := "0000000000000000000000000000000000000000"
//	response := client.InvokeScript(helper.BytesToHex(script), checkWitnessHashes)
//	if response.HasError() || response.Result.State == "FAULT" {
//		log.Printf("invoke script error: %s", response.Error.Message)
//	}
//
//	if len(response.Result.Stack) > 0 {
//		var amount *big.Int
//		var success bool
//		for i, stack := range response.Result.Stack {
//			if stack.Type == "ByteArray" {
//				amount = helper.BigIntFromNeoBytes(helper.HexToBytes(stack.Value.(string)))
//			} else {
//				amount, success = new(big.Int).SetString(stack.Value.(string), 10)
//				assert.Equal(t, true, success)
//			}
//			if i < 6 {
//				log.Printf("staking amount: %s", amount.String())
//			} else {
//				log.Printf("total amount: %s", amount.String())
//			}
//		}
//	} else {
//		log.Printf("stack empty")
//	}
//}
//
//func Test2(t *testing.T) {
//	client := NewClient("https://wallet.ngd.network:10331")
//	response := client.GetNep5Balances("Abj465Y7SYEWRg6sgUsN8hH8SWLxZmuHqZ")
//	if response.HasError() {
//		log.Printf("invoke script error: %s", response.Error.Message)
//	}
//
//	for _, b := range response.Result.Balances {
//		log.Printf("balance: %s", b.Amount)
//	}
//}
//
//func Test3(t *testing.T) {
//	sb := sc.NewScriptBuilder()
//	scriptHash := helper.HexToBytes("9a9db8a30a80951ec792effb9731af79781177c2") // pONT
//	address := "AQ8JYMgWASbHFw1YvUjoBvGtUGuTQb6sDk"
//	lyh, _ := helper.AddressToScriptHash(address)
//	log.Printf("lyh: %s", helper.BytesToHex(lyh.Bytes()))
//
//	cp1 := sc.ContractParameter{
//		Type:  sc.ByteArray,
//		Value: lyh.Bytes(),
//	}
//
//	args := []sc.ContractParameter{cp1}
//	sb.MakeInvocationScript(scriptHash, "balanceOf", args)
//	script := sb.ToArray()
//
//	log.Printf("script: %s", helper.BytesToHex(script))
//
//	//client := NewClient("http://seed10.ngd.network:11332")
//	client := NewClient("https://wallet.ngd.network:10331")
//
//	checkWitnessHashes := "0000000000000000000000000000000000000000"
//	response := client.InvokeScript(helper.BytesToHex(script), checkWitnessHashes)
//	if response.HasError() || response.Result.State == "FAULT" {
//		log.Printf("invoke script error: %s", response.Error.Message)
//	}
//
//	stack := response.Result.Stack[0]
//	var amount *big.Int
//	var success bool
//	if stack.Type == "ByteArray" {
//		amount = helper.BigIntFromNeoBytes(helper.HexToBytes(stack.Value.(string)))
//	} else {
//		amount, success = new(big.Int).SetString(stack.Value.(string), 10)
//		assert.Equal(t, true, success)
//	}
//
//	log.Printf("balance: %s", amount.String())
//}
