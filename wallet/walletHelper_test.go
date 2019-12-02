package wallet

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewWalletHelper(t *testing.T) {
	txBuilder := tx.NewTransactionBuilder("http://seed1.ngd.network:20332")
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.NotNil(t, txBuilder)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	walletHelper := NewWalletHelper(txBuilder, account)
	assert.NotNil(t, walletHelper)
	assert.Equal(t, "http://seed1.ngd.network:20332", walletHelper.TxBuilder.EndPoint)
	assert.Equal(t, "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619", walletHelper.Account.KeyPair.PublicKey.String())
}

func TestWalletHelper_ClaimGas(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = &tx.TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.Nil(t, err)
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
	clientMock.On("SendRawTransaction", mock.Anything).Return(rpc.SendRawTransactionResponse{
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
		Result: true,
	})

	walletHelper := NewWalletHelper(tb, account)
	b, err := walletHelper.ClaimGas("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestWalletHelper_Transfer(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = &tx.TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.Nil(t, err)

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

	clientMock.On("SendRawTransaction", mock.Anything).Return(rpc.SendRawTransactionResponse{
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
		Result: true,
	})

	walletHelper := NewWalletHelper(tb, account)
	b, err := walletHelper.Transfer(tx.NeoToken, "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt", "AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV", 80000)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestWalletHelper_TransferNep5(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = &tx.TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.Nil(t, err)
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
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeScriptResponse{
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
		Result: models.InvokeResult{
			Script:"00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.126",
			Stack: []models.InvokeStackResult{
				{
					Type:  "Boolean",
					Value: "True",
				},
			},
		},
	})

	clientMock.On("SendRawTransaction", mock.Anything).Return(rpc.SendRawTransactionResponse{
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
		Result: true,
	})

	walletHelper := NewWalletHelper(tb, account)
	scriptHash, err := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	assert.Nil(t, err)
	b, err := walletHelper.TransferNep5(scriptHash, "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt", "AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV", 80000)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}