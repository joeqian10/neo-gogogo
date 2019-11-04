package keys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	for _, testCase := range KeyCases {
		p, err := NewPublicKeyFromString(testCase.PublicKey)
		assert.Nil(t, err)

		address := p.Address()
		assert.Equal(t, testCase.Address, address)
	}
}
