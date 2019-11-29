package nep5

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewNep5Helper(t *testing.T) {
	nep5helper := NewNep5Helper("http://seed1.ngd.network:20332")
	if nep5helper == nil {
		t.Fail()
	}
	assert.Equal(t, "seed1.ngd.network:20332", nep5helper.EndPoint)
}

func TestNep5Helper_BalanceOf(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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
			Script:"148f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc51c10962616c616e63654f666763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.383",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "004eaca7902c",
				},
			},
		},
	})

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
	u, e := nh.BalanceOf(scriptHash, address)
	assert.Nil(t, e)
	assert.Equal(t, uint64(48999800000000), u)
}

func TestNep5Helper_Decimals(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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
			Script:"00c108646563696d616c736763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.246",
			Stack: []models.InvokeStackResult{
				{
					Type:  "Integer",
					Value: "8",
				},
			},
		},
	})

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	d, err := nh.Decimals(scriptHash)
	assert.Nil(t, err)
	assert.Equal(t, uint8(8), d)
}

func TestNep5Helper_Name(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	name, err := nh.Name(scriptHash)
	//name := string(helper.HexTobytes("516c696e6b20546f6b656e"))
	assert.Nil(t, err)
	assert.Equal(t, "Qlink Token", name)
}

func TestNep5Helper_Symbol(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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
			Script:"00c10673796d626f6c6763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.141",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "514c43",
				},
			},
		},
	})

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	name, err := nh.Name(scriptHash)
	assert.Nil(t, err)
	assert.Equal(t, "QLC", name)
}

func TestNep5Helper_TotalSupply(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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
			Script:"00c10b746f74616c537570706c796763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.223",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "00803dafbe50d300",
				},
			},
		},
	})

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	s, e := nh.TotalSupply(scriptHash)
	assert.Nil(t, e)
	assert.Equal(t, uint64(59480000000000000), s)
}

func TestNep5Helper_Transfer(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
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
			Script:"5114dcdf11cccb2efd9bbfa8449e57b60c9ce85b6c8f14dcdf11cccb2efd9bbfa8449e57b60c9ce85b6c8f53c1087472616e736665726763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"0.494",
			Stack: []models.InvokeStackResult{
				{
					Type:  "Boolean",
					Value: "True",
				},
			},
		},
	})

	scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
	b, e := nh.Transfer(scriptHash, address, address, 1)
	assert.Nil(t, e)
	assert.Equal(t, true, b)
}