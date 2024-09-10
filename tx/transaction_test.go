package tx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo-gogogo/wallet/keys"
)

func TestAddSignature(t *testing.T) {
	// mainNet transaction: bdf6cc3b9af12a7565bda80933a75ee8cef1bc771d0d58effc08e4c8b436da79
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	assert.Nil(t, err)
	assert.Equal(t, Contract_Transaction, ctx.Type)
	assert.IsType(t, ctx, &ContractTransaction{})

	key, _ := keys.GenerateKeyPair()
	ctx.Witnesses = make([]*Witness, 0)
	err = AddSignature(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, ctx.Attributes[0].Data, key.PublicKey.ScriptHash().Bytes())
	assert.True(t, keys.VerifySignature(ctx.UnsignedRawTransaction(), ctx.Witnesses[0].InvocationScript[1:], key.PublicKey))

	err = AddSignature(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ctx.Witnesses))
}

func TestAddMultiSignature(t *testing.T) {
	// mainNet transaction: bdf6cc3b9af12a7565bda80933a75ee8cef1bc771d0d58effc08e4c8b436da79
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	//ctx := NewContractTransaction()
	// Deserialize
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	assert.Nil(t, err)
	assert.Equal(t, Contract_Transaction, ctx.Type)
	assert.IsType(t, ctx, &ContractTransaction{})

	key, _ := keys.GenerateKeyPair()
	key2, _ := keys.GenerateKeyPair()
	key3, _ := keys.GenerateKeyPair()
	pairs := []*keys.KeyPair{key, key2, key3}
	sort.Sort(sort.Reverse(keys.KeyPairSlice(pairs)))

	ctx.Witnesses = make([]*Witness, 0)
	err = AddMultiSignature(ctx, pairs, 2, []*keys.PublicKey{key.PublicKey, key2.PublicKey, key3.PublicKey})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ctx.Attributes))
	assert.True(t, keys.VerifySignature(ctx.UnsignedRawTransaction(), ctx.Witnesses[0].InvocationScript[1:], pairs[0].PublicKey))

	err = AddMultiSignature(ctx, pairs, 2, []*keys.PublicKey{key.PublicKey, key2.PublicKey, key3.PublicKey})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ctx.Witnesses))

	err = AddSignature(ctx, key)
	err = AddSignature(ctx, key2)
	err = AddSignature(ctx, key3)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(ctx.Attributes))
	assert.Equal(t, 4, len(ctx.Witnesses))
	assert.True(t, ctx.Witnesses[0].scriptHash.Less(ctx.Witnesses[1].scriptHash))
	assert.True(t, ctx.Witnesses[1].scriptHash.Less(ctx.Witnesses[2].scriptHash))
	assert.True(t, ctx.Witnesses[2].scriptHash.Less(ctx.Witnesses[3].scriptHash))
}
