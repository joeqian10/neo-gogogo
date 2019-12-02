package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvocationTransaction(t *testing.T) {
	rawTx := "d101590400b33f7114839c33710da24cf8e7d536b8d244f3991cf565c8146063795d3b9b3cd55aef026eae992b91063db0db53c1087472616e7366657267c5cc1cb5392019e2cc4e6d6b5ea54c8d4b6d11acf166cb072961424c54f6000000000000000001206063795d3b9b3cd55aef026eae992b91063db0db0000014140c6a131c55ca38995402dff8e92ac55d89cbed4b98dfebbcb01acbc01bd78fa2ce2061be921b8999a9ab79c2958875bccfafe7ce1bbbaf1f56580815ea3a4feed232102d41ddce2c97be4c9aa571b8a32cbc305aa29afffbcae71b0ef568db0e93929aaac"
	itx := &InvocationTransaction{Transaction: NewTransaction()}
	// Deserialize
	itx, err := itx.FromHexString(rawTx)
	assert.Nil(t, err)
	assert.Equal(t, itx.Type, Invocation_Transaction)
	assert.IsType(t, itx, &InvocationTransaction{})

	script := "0400b33f7114839c33710da24cf8e7d536b8d244f3991cf565c8146063795d3b9b3cd55aef026eae992b91063db0db53c1087472616e7366657267c5cc1cb5392019e2cc4e6d6b5ea54c8d4b6d11acf166cb072961424c54f6"
	assert.Equal(t, script, hex.EncodeToString(itx.Script))
	assert.Equal(t, helper.Fixed8FromInt64(0), itx.Gas)

	assert.Equal(t, 1, len(itx.Attributes))
	assert.Equal(t, 0, len(itx.Inputs))
	assert.Equal(t, 0, len(itx.Outputs))
	invoc := "40c6a131c55ca38995402dff8e92ac55d89cbed4b98dfebbcb01acbc01bd78fa2ce2061be921b8999a9ab79c2958875bccfafe7ce1bbbaf1f56580815ea3a4feed"
	verif := "2102d41ddce2c97be4c9aa571b8a32cbc305aa29afffbcae71b0ef568db0e93929aaac"
	assert.Equal(t, 1, len(itx.Witnesses))
	assert.Equal(t, invoc, hex.EncodeToString(itx.Witnesses[0].InvocationScript))
	assert.Equal(t, verif, hex.EncodeToString(itx.Witnesses[0].VerificationScript))

	// test Size()
	assert.Equal(t, len(rawTx)/2, itx.Size())

	// Serialize
	buf := io.NewBufBinaryWriter()
	itx.Serialize(buf.BinaryWriter)
	assert.Nil(t, buf.Err)

	assert.Equal(t, rawTx, hex.EncodeToString(buf.Bytes()))
}
