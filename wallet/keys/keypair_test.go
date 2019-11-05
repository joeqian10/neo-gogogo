package keys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
