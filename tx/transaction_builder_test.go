package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeContractTransaction(t *testing.T) {
	from, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	to, _ := helper.AddressToScriptHash("AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV")
	assetId := NeoToken
	amount := helper.Fixed8FromInt64(50000000)
	ctx, _ := MakeContractTransaction(from, to, assetId, amount, nil, helper.UInt160{}, helper.Fixed8FromInt64(0))
	outputs := ctx.Outputs
	output := outputs[0]
	assert.Equal(t, amount.Value, output.Value.Value)
}

func TestMakeInvocationTransaction(t *testing.T) {
	from, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	scriptHash := helper.ReverseBytes(helper.HexTobytes("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26"))
	operation := "name"
	itx, _ := MakeInvocationTransaction(scriptHash, operation, nil, from, nil, helper.UInt160{}, helper.Fixed8FromInt64(0))
	assert.Equal(t, helper.Fixed8FromInt64(0).Value, itx.Gas.Value)
}

// you need to use your private net for testing
func TestMakeClaimTransaction(t *testing.T) {
	from, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	ctx, _ := MakeClaimTransaction(from, helper.UInt160{}, nil)
	outputs := ctx.Outputs
	output := outputs[0]
	assert.Equal(t, helper.Fixed8FromInt64(14248).Value, output.Value.Value)
}