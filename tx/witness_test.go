package tx

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
)

func TestWitness_Deserialize(t *testing.T) {
	s := "41" + "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58" +
		"23" + "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	w := Witness{}
	w.Deserialize(br)
	assert.Equal(t, "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58", helper.BytesToHex(w.InvocationScript))
	assert.Equal(t, "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(w.VerificationScript))
}

func TestWitness_GetScriptHash(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexToBytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	scriptHash := w.GetScriptHash()
	assert.Equal(t, "71cb588c8291c18fa87fa07ce16c3fd92ab5aa30", scriptHash.String())
}

func TestWitness_Serialize(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexToBytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	bbw := io.NewBufBinaryWriter()
	w.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "41"+"40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"+
		"23"+"2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(b))
}

func TestWitness_Serialize2(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("520131"),
		VerificationScript: helper.HexToBytes(""),
	}
	bbw := io.NewBufBinaryWriter()
	w.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	log.Printf(":%s", helper.BytesToHex(b))
}

func TestCreateSignatureWitness(t *testing.T) {
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	msg := ctx.UnsignedRawTransaction()
	pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[0].Wif)
	witness, err := CreateSignatureWitness(msg, pair)
	assert.Nil(t, err)
	assert.Equal(t, 65, len(witness.InvocationScript))
	assert.Equal(t, 35, len(witness.VerificationScript))
}

func TestCreateMultiSignatureWitness(t *testing.T) {
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	msg := ctx.UnsignedRawTransaction()

	pairs := make([]*keys.KeyPair, 4)
	pubKeys := make([]*keys.PublicKey, 4)
	for i := 0; i < 4; i++ {
		pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[i].Wif)
		pairs[i] = pair
		pubKeys[i] = pair.PublicKey
	}

	witness, err := CreateMultiSignatureWitness(msg, pairs[:3], 3, pubKeys)
	assert.Nil(t, err)
	assert.Equal(t, 65*3, len(witness.InvocationScript))
	assert.Equal(t, 1+34*4+1+1, len(witness.VerificationScript))
}

func TestVerifySignatureWitness(t *testing.T) {
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	msg := ctx.UnsignedRawTransaction()
	pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[0].Wif)
	witness, err := CreateSignatureWitness(msg, pair)
	b := VerifySignatureWitness(msg, witness)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestVerifyMultiSignatureWitness(t *testing.T) {
	rawTx := "80000001888da99f8f497fd65c4325786a09511159c279af4e7eb532e9edd628c87cc1ee0000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50082167010000000a8666b4830229d6a1a9b80f6088059191c122d2b0141409e79e132290c82916a88f1a3db5cf9f3248b780cfece938ab0f0812d0e188f3a489c7d1a23def86bd69d863ae67de753b2c2392e9497eadc8eb9fc43aa52c645232103e2f6a334e05002624cf616f01a62cff2844c34a3b08ca16048c259097e315078ac"
	ctx := &ContractTransaction{NewTransaction()}
	ctx, err := ctx.FromHexString(rawTx)
	msg := ctx.UnsignedRawTransaction()
	pairs := make([]*keys.KeyPair, 4)
	pubKeys := make([]*keys.PublicKey, 4)
	for i := 0; i < 4; i++ {
		pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[i].Wif)
		pairs[i] = pair
		pubKeys[i] = pair.PublicKey
	}
	witness, err := CreateMultiSignatureWitness(msg, pairs[:3], 3, pubKeys)
	b := VerifyMultiSignatureWitness(msg, witness)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}
