package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBytes(t *testing.T) {
	var b = make([]byte, 0)
	r := ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1}
	r = ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1, 2}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{2, 1}, r)

	b = []byte{1, 2, 3}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{1, 2, 3}, b)
	assert.Equal(t, []byte{3, 2, 1}, r)
}
