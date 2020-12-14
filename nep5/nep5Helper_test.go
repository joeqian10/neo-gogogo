package nep5

import (
	"testing"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewNep5Helper(t *testing.T) {
	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	nep5helper := NewNep5Helper(scriptHash, "http://seed1.ngd.network:20332")
	assert.NotNil(t, nep5helper)
}

func TestNep5Helper_BalanceOf(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	var nh = Nep5Helper{
		scriptHash: scriptHash,
		Client:     clientMock,
	}
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
			Script:      "148f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc51c10962616c616e63654f666763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "0.383",
			Stack: []models.InvokeStack{
				{
					Type:  "ByteArray",
					Value: "004eaca7902c",
				},
			},
		},
	})

	address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
	u, e := nh.BalanceOf(address)
	assert.Nil(t, e)
	assert.Equal(t, uint64(48999800000000), u)
}

func TestNep5Helper_Decimals(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	var nh = Nep5Helper{
		scriptHash: scriptHash,
		Client:     clientMock,
	}
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
			Script:      "00c108646563696d616c736763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "0.246",
			Stack: []models.InvokeStack{
				{
					Type:  "Integer",
					Value: "8",
				},
			},
		},
	})

	d, err := nh.Decimals()
	assert.Nil(t, err)
	assert.Equal(t, uint8(8), d)
}

func TestNep5Helper_Name(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	var nh = Nep5Helper{
		scriptHash: scriptHash,
		Client:     clientMock,
	}
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
					Type:  "ByteArray",
					Value: "516c696e6b20546f6b656e",
				},
			},
		},
	})

	name, err := nh.Name()
	//name := string(helper.HexToBytes("516c696e6b20546f6b656e"))
	assert.Nil(t, err)
	assert.Equal(t, "Qlink Token", name)
}

func TestNep5Helper_Symbol(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	var nh = Nep5Helper{
		scriptHash: scriptHash,
		Client:     clientMock,
	}
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
			Script:      "00c10673796d626f6c6763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "0.141",
			Stack: []models.InvokeStack{
				{
					Type:  "ByteArray",
					Value: "514c43",
				},
			},
		},
	})

	name, err := nh.Name()
	assert.Nil(t, err)
	assert.Equal(t, "QLC", name)
}

func TestNep5Helper_TotalSupply(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		Client: clientMock,
	}
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
			Script:      "00c10b746f74616c537570706c796763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "0.223",
			Stack: []models.InvokeStack{
				{
					Type:  "ByteArray",
					Value: "00803dafbe50d300",
				},
			},
		},
	})

	s, e := nh.TotalSupply()
	assert.Nil(t, e)
	assert.Equal(t, uint64(59480000000000000), s)
}

//func TestNep5Helper_Transfer(t *testing.T) {
//	var clientMock = new(rpc.RpcClientMock)
//	var nh = Nep5Helper{
//		Client: clientMock,
//	}
//	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeScriptResponse{
//		RpcResponse: rpc.RpcResponse{
//			JsonRpc: "2.0",
//			ID:      1,
//		},
//		ErrorResponse: rpc.ErrorResponse{
//			Error: rpc.RpcError{
//				Code:    0,
//				Message: "",
//			},
//		},
//		Result: models.InvokeResult{
//			Script:      "5114dcdf11cccb2efd9bbfa8449e57b60c9ce85b6c8f14dcdf11cccb2efd9bbfa8449e57b60c9ce85b6c8f53c1087472616e736665726763d26113bac4208254d98a3eebaee66230ead7b9",
//			State:       "HALT",
//			GasConsumed: "0.494",
//			Stack: []models.InvokeStackResult{
//				{
//					Type:  "Boolean",
//					Value: "True",
//				},
//			},
//		},
//	})
//
//	address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
//	b, _, e := nh.Transfer(address, address, helper.Fixed8FromInt64(1))
//	assert.Nil(t, e)
//	assert.Equal(t, true, b)
//}
