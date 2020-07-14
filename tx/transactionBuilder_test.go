package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewTransactionBuilder(t *testing.T) {
	tb := NewTransactionBuilder("http://seed1.ngd.network:20332")
	if tb == nil {
		t.Fail()
	}
	assert.Equal(t, "http://seed1.ngd.network:20332", tb.EndPoint)
}

func TestTransactionBuilder_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("GetUnspents", mock.Anything).Return(rpc.GetUnspentsResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID: 1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.RpcUnspent{
			Balances: []models.UnspentBalance{
				{
					Unspents:    []models.Unspent{
						{
							Txid:  "4ee4af75d5aa60598fbae40ce86fb9a23ffec5a75dfa8b59d259d15f9e304319",
							N:     0,
							Value: 27844.821,
						},
					},
					AssetHash:   "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
					Asset:       "GAS",
					AssetSymbol: "GAS",
					Amount:      27844.821,
				},
				{
					Unspents:    []models.Unspent{
						{
							Txid:  "c3182952855314b3f4b1ecf01a03b891d4627d19426ce841275f6d4c186e729a",
							N:     0,
							Value: 800000,
						},
					},
					AssetHash:   "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
					Asset:       "NEO",
					AssetSymbol: "NEO",
					Amount:      800000,
				},
			},
			Address: "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
		},
	})

	u, a, e := tb.GetBalance(helper.UInt160{}, GasToken)
	assert.Nil(t, e)
	assert.Equal(t, GasTokenId, u.AssetHash)
	assert.Equal(t, helper.Fixed8FromFloat64(27844.821).Value, a.Value)
}

func TestTransactionBuilder_GetGasConsumed(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeScriptResponse{
		RpcResponse:   rpc.RpcResponse{
			JsonRpc: "2.0",
			ID: 1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result:        models.InvokeResult{
			Script:"00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.126",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "516c696e6b20546f6b656e",
				},
			},
		},
	})

	f, e := tb.GetGasConsumed([]byte{})
	assert.Nil(t, e)
	assert.Equal(t, helper.Fixed8FromFloat64(0).Value, f.Value) // 10 gas free limit
}

func TestTransactionBuilder_GetClaimables(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("GetClaimable", mock.Anything).Return(rpc.GetClaimableResponse{
		RpcResponse:   rpc.RpcResponse{
			JsonRpc: "2.0",
			ID: 1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result:        models.RpcClaimable{
			Claimables: []models.Claimable{
				{
					TxId:        "52ba70ef18e879785572c917795cd81422c3820b8cf44c24846a30ee7376fd77",
					N:           1,
					Value:       800000,
					StartHeight: 476496,
					EndHeight:   488154,
					Generated:   746.112,
					SysFee:      3.92,
					Unclaimed:   750.032,
				},
			},
			Address:"AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
			Unclaimed:750.032,
		},
	})

	crs, f, e := tb.GetClaimables(helper.UInt160{})
	assert.Nil(t, e)
	assert.NotNil(t, crs)
	assert.Equal(t, 1, len(crs))
	cr := crs[0]
	assert.Equal(t, "52ba70ef18e879785572c917795cd81422c3820b8cf44c24846a30ee7376fd77", cr.PrevHash.String())
	assert.Equal(t, helper.Fixed8FromFloat64(750.032).Value, f.Value)
}

func TestTransactionBuilder_GetTransactionInputs(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("GetUnspents", mock.Anything).Return(rpc.GetUnspentsResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.RpcUnspent{
			Balances: []models.UnspentBalance{
				{
					Unspents: []models.Unspent{
						{
							Txid:  "0a99ebd286931375c2ec828603e88392e3a40e9cecd4b228bd6be206fdb21005",
							N:     0,
							Value: 11250,
						},
						{
							Txid:  "1c2f4605fa4c5ba9ca2a8ae87ae083a241d407f59472e707fe34e52d277d2331",
							N:     1,
							Value: 81.96167,
						},
						{
							Txid:  "6d7de9f60c1b8c86f3a2a0d8001c051e9192d59fd2e922f02516770533e0cfc4",
							N:     0,
							Value: 0.03833,
						},
					},
					AssetHash:   "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
					Asset:       "GAS",
					AssetSymbol: "GAS",
					Amount:      11332,
				},
				{
					Unspents: []models.Unspent{
						{
							Txid:  "c724d26a3e2bb4417f6cebd56a7c5138987dc0b49b41fe1b5c632f5208c1e05f",
							N:     0,
							Value: 100000000,
						},
					},
					AssetHash:   "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
					Asset:       "NEO",
					AssetSymbol: "NEO",
					Amount:      100000000,
				},
			},
			Address: "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
		},
	})

	inputs, payTotal, err := tb.GetTransactionInputs(helper.UInt160{}, GasToken, helper.Fixed8FromFloat64(10000))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(inputs))
	assert.Equal(t, int64(1125000000000), payTotal.Value)
}

func TestTransactionBuilder_LoadScriptTransaction(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}

	script := []byte{ 0x01, 0x02, 0x03, 0x04 }
	paramList := "0710"
	returnType := "05"
	hasStorage := true
	hasDynamicInvoke := true
	isPayable := true
	var contractName = "test"
	var contractVersion = "1.0"
	var contractAuthor = "ngd"
	var contractEmail = "test@ngd.neo.org"
	var contractDescription = "cd"

	itx, scriptHash, err := tb.LoadScriptTransaction(script, paramList, returnType, hasStorage, hasDynamicInvoke, isPayable, contractName, contractVersion, contractAuthor, contractEmail, contractDescription)

	assert.Nil(t, err)
	assert.Equal(t, "706ea1768da7f0c489bf931b362c2d26d8cbd2ec", scriptHash.String())
	assert.Equal(t, "0263641074657374406e67642e6e656f2e6f7267036e676403312e300474657374575502071004010203046804f66ca56e", helper.BytesToHex(itx.Script))
}
