package keys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportWIF(t *testing.T) {
	for _, testCase := range KeyCases {
		privKey, err := NewKeyPairFromWIF(testCase.Wif)
		privateKey := privKey.String()
		publicKey := privKey.PublicKey.String()
		wif := privKey.ExportWIF()
		nep2, err := privKey.ExportNep2(testCase.Passphrase)

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
		assert.Equal(t, testCase.Wif, wif)
		assert.Equal(t, testCase.Nep2key, nep2)
	}
}

func TestExportNEP2(t *testing.T) {
	for _, testCase := range KeyCases {
		privKey, err := NewKeyPairFromNEP2(testCase.Nep2key, testCase.Passphrase)
		privateKey := privKey.String()
		publicKey := privKey.PublicKey.String()
		wif := privKey.ExportWIF()
		nep2, err := privKey.ExportNep2(testCase.Passphrase)

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
		assert.Equal(t, testCase.Wif, wif)
		assert.Equal(t, testCase.Nep2key, nep2)
	}
}

func TestPubKeyVerify(t *testing.T) {
	var data = []byte("sample")
	privKey, err := GenerateKeyPair()
	assert.Nil(t, err)
	signedData, err := privKey.Sign(data)
	assert.Nil(t, err)
	pubKey := privKey.PublicKey
	result := VerifySignature(data, signedData, pubKey)
	expected := true
	assert.Equal(t, expected, result)
}

func TestWrongPubKey(t *testing.T) {
	privKey, _ := GenerateKeyPair()
	sample := []byte("sample")
	signedData, _ := privKey.Sign(sample)

	secondPrivKey, _ := GenerateKeyPair()
	wrongPubKey := secondPrivKey.PublicKey

	actual := VerifySignature(sample, signedData, wrongPubKey)
	expcted := false
	assert.Equal(t, expcted, actual)
}
