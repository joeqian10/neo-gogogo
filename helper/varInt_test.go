package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarIntFromUInt64(t *testing.T) {
	v := VarIntFromUInt64(100000000)
	assert.Equal(t, uint64(100000000), v.Value)
}

func TestVarInt_Length(t *testing.T) {
	v := VarIntFromUInt64(100000000)
	assert.Equal(t, 5, v.Length())
}

func TestVarInt_Bytes(t *testing.T) {
	v := VarIntFromUInt64(100000000)
	assert.Equal(t, "fe00e1f505", BytesToHex(v.Bytes()))
}
