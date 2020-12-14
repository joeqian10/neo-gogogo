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

func TestWalletHelper_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = &tx.TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.Nil(t, err)
	clientMock.On("GetAccountState", mock.Anything).Return(rpc.GetAccountStateResponse{
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
		Result: models.AccountState{
			Version:    0,
			ScriptHash: "0x1179716da2e9523d153a35fb3ad10c561b1e5b1a",
			Frozen:     false,
			Votes:      make([]interface{}, 0),
			Balances: []models.AccountStateBalance{
				{
					Asset: "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
					Value: "100",
				},
				{
					Asset: "0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7",
					Value: "90.12345678",
				},
			},
		},
	})

	walletHelper := NewWalletHelper(tb, account)
	neoBalance, gasBalance, err := walletHelper.GetBalance("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	assert.Nil(t, err)
	assert.Equal(t, 100, neoBalance)
	assert.Equal(t, 90.12345678, gasBalance)
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
		Result: models.RpcClaimable{
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
			Address:   "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
			Unclaimed: 750.032,
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
	assert.Equal(t, "00de0ab4da0475f018fcc751ee9979ab21003a63e360908e6933e07423e25ae1", b)
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
					Unspents: []models.Unspent{
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
	assert.Equal(t, "e073d3e3524c16fb2996ba1dc81d1cb51364731143b0afebff6463ec17b94c32", b)
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
					Unspents: []models.Unspent{
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
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeScriptResponse{
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
		Result: models.InvokeResult{
			Script:      "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "0.126",
			Stack: []models.InvokeStack{
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
	s, err := walletHelper.TransferNep5(scriptHash, "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt", "AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV", 80000)
	assert.Nil(t, err)
	assert.Equal(t, 64, len(s))
}

//func TestWallet_Transfer(t *testing.T) {
//	txBuilder := tx.NewTransactionBuilder("http://seed2.ngd.network:20332")
//	account, _ := NewAccountFromWIF("L2LGkrwiNmUAnWYb1XGd5mv7v2eDf6P4F3gHyXSrNJJR4ArmBp7Q")
//	address := "AKeLhhHm4hEUfLWVBCYRNjio9xhGJAom5G"
//	api := NewWalletHelper(txBuilder, account)
//	neoBalance, gasBalace, _ := api.GetBalance(address)
//
//	assert.Equal(t, 800, neoBalance)
//	assert.Equal(t, 500.12345678, gasBalace)
//
//	result, err := api.Transfer(tx.NeoToken, address, "AR2uSMBjLv1RppjW9dYn4PHTnuPyBKtGta", 200)
//	assert.Nil(t, err)
//	assert.True(t, result)
//
//	claimable := txBuilder.Client.GetClaimable(address)
//	assert.True(t, claimable.Result.Unclaimed > 0)
//
//	res, _ := api.ClaimGas(address)
//	assert.True(t, res)
//}
//
//func TestWallet_NEP5(t *testing.T) {
//	txBuilder := tx.NewTransactionBuilder("http://localhost:30333")
//	account, _ := NewAccountFromWIF("L2LGkrwiNmUAnWYb1XGd5mv7v2eDf6P4F3gHyXSrNJJR4ArmBp7Q")
//	address := "AKeLhhHm4hEUfLWVBCYRNjio9xhGJAom5G"
//	api := NewWalletHelper(txBuilder, account)
//
//	tokenHash, _ := helper.UInt160FromString("0x43bb08d7c03ac66582079b57059108565f91ece5")
//	addressHash, _ := helper.AddressToScriptHash(address)
//	nep5Api := nep5.NewNep5Helper("http://localhost:30333")
//	tokenBalance, _ := nep5Api.BalanceOf(tokenHash, addressHash)
//
//	result, err := api.TransferNep5(tokenHash, "AKeLhhHm4hEUfLWVBCYRNjio9xhGJAom5G", "AdmyedL3jdw2TLvBzoUD2yU443NeKrP5t5", 200)
//	assert.Nil(t, err)
//	assert.True(t, result)
//
//	tokenBalanceAfter, _ := nep5Api.BalanceOf(tokenHash, addressHash)
//	assert.Equal(t, uint64(0), tokenBalanceAfter-tokenBalance)
//}
