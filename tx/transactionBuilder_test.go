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