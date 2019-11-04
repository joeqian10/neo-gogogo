package keys

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNEP2Encrypt(t *testing.T) {
	for _, testCase := range KeyCases {
		b, _ := hex.DecodeString(testCase.PrivateKey)

		privKey, err := NewKeyPair(b)
		assert.Nil(t, err)

		encryptedWif, err := NEP2Encrypt(privKey, testCase.Passphrase)
		assert.Nil(t, err)

		assert.Equal(t, testCase.Nep2key, encryptedWif)
	}
}

func TestNEP2Decrypt(t *testing.T) {
	for _, testCase := range KeyCases {

		privKey, err := NEP2Decrypt(testCase.Nep2key, testCase.Passphrase)
		assert.Nil(t, err)

		assert.Equal(t, testCase.PrivateKey, privKey.String())

		wif := privKey.ExportWIF()
		assert.Equal(t, testCase.Wif, wif)

		address := privKey.PublicKey.Address()
		assert.Equal(t, testCase.Address, address)
	}
}
