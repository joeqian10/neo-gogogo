package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO
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