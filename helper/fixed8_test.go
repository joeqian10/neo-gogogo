package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFixed8(t *testing.T) {
	f := NewFixed8(100000000)
	assert.Equal(t, int64(100000000), f.Value)
}

func TestFixed8FromInt64(t *testing.T) {
	f := Fixed8FromInt64(100000000)
	assert.Equal(t, int64(100000000*D), f.Value)
}

func TestFixed8FromFloat64(t *testing.T) {
	f := Fixed8FromFloat64(0.12345678)
	assert.Equal(t, int64(12345678), f.Value)
}

func TestFixed8FromString(t *testing.T) {
	f, err := Fixed8FromString("1234.5678")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456780000), f.Value)
}

func TestFixed8ToInt64(t *testing.T) {
	n := Fixed8ToInt64(NewFixed8(100000000))
	assert.Equal(t, int64(1), n)
}

func TestFixed8ToFloat64(t *testing.T) {
	c := Fixed8ToFloat64(NewFixed8(12345678))
	assert.Equal(t, float64(0.12345678), c)
}

func TestFixed8ToString(t *testing.T) {
	s := Fixed8ToString(NewFixed8(100000000))
	assert.Equal(t, "1", s)
}

func TestFixed8_Abs(t *testing.T) {
	f := NewFixed8(-100000000).Abs()
	assert.Equal(t, int64(100000000), f.Value)
}

func TestFixed8_Add(t *testing.T) {
	f := NewFixed8(100000000).Add(NewFixed8(100000000))
	assert.Equal(t, int64(200000000), f.Value)
}

func TestFixed8_Sub(t *testing.T) {
	f := NewFixed8(200000000).Sub(NewFixed8(100000000))
	assert.Equal(t, int64(100000000), f.Value)
}

func TestFixed8_Mul(t *testing.T) {
	f, err := NewFixed8(200000000).Mul(NewFixed8(100000000))
	assert.Nil(t, err)
	assert.Equal(t, int64(200000000), f.Value)
}

func TestFixed8_Div(t *testing.T) {
	f := NewFixed8(200000000).Div(NewFixed8(100000000))
	assert.Equal(t, int64(200000000), f.Value)
}

func TestFixed8_Ceiling(t *testing.T) {
	f := NewFixed8(12345678).Ceiling()
	assert.Equal(t, int64(100000000), f.Value)
}

func TestFixed8_GreaterThan(t *testing.T) {
	b := NewFixed8(200000000).GreaterThan(NewFixed8(100000000))
	assert.Equal(t, true, b)
}

func TestFixed8_Equal(t *testing.T) {
	b := NewFixed8(200000000).Equal(NewFixed8(200000000))
	assert.Equal(t, true, b)
}

func TestFixed8_LessThan(t *testing.T) {
	b := NewFixed8(200000000).LessThan(NewFixed8(300000000))
	assert.Equal(t, true, b)
}

func TestFixed8_String(t *testing.T) {
	s := NewFixed8(100000000).String()
	assert.Equal(t, "1", s)
}
