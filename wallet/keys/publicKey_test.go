package keys

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	for _, testCase := range KeyCases {
		p, err := NewPublicKeyFromString(testCase.PublicKey)
		assert.Nil(t, err)

		address := p.Address()
		assert.Equal(t, testCase.Address, address)

		scripthash := p.ScriptHash()
		assert.Equal(t, testCase.ScriptHash, scripthash.String())
	}
}

func TestCreateMultiSigRedeemScript(t *testing.T) {
	privateKey1, _ := hex.DecodeString(KeyCases[0].PrivateKey)
	privateKey2, _ := hex.DecodeString(KeyCases[1].PrivateKey)
	privateKey3, _ := hex.DecodeString(KeyCases[2].PrivateKey)
	privateKey4, _ := hex.DecodeString(KeyCases[3].PrivateKey)


	keyPair1, _ := NewKeyPair(privateKey1)
	keyPair2, _ := NewKeyPair(privateKey2)
	keyPair3, _ := NewKeyPair(privateKey3)
	keyPair4, _ := NewKeyPair(privateKey4)

	multiSignature, _ := CreateMultiSigRedeemScript(3, keyPair1.PublicKey, keyPair2.PublicKey, keyPair3.PublicKey, keyPair4.PublicKey)

	assert.Equal(t, "5321027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b21038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a12103b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe554444266192103d08d6f766b54e35745bc99d643c939ec6f3d37004f2a59006be0e53610f0be2554ae", hex.EncodeToString(multiSignature))
}
