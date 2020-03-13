package helper

import (
	"testing"

	nio "github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
)

type TestSerialzable struct {
	Flag  bool
	Value []byte
}

func (t *TestSerialzable) Serialize(writer *nio.BinaryWriter) {
	writer.WriteLE(t.Flag)
	writer.WriteVarBytes(t.Value)
}

func (t *TestSerialzable) Deserialize(reader *nio.BinaryReader) {
	reader.ReadLE(&t.Flag)
	t.Value = reader.ReadVarBytes()
}

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

func TestToArray(t *testing.T) {
	ts := &TestSerialzable{
		Flag:  true,
		Value: HexToBytes("abcd"),
	}
	data, _ := ToArray(ts)
	assert.Equal(t, data, []byte{0x01, 0x02, 0xab, 0xcd})
}

func TestAsSerializable(t *testing.T) {
	data := []byte{0x01, 0x02, 0xab, 0xcd}
	ts := &TestSerialzable{}
	AsSerializable(ts, data)
	assert.Equal(t, true, ts.Flag)
	assert.Equal(t, []byte{0xab, 0xcd}, ts.Value)
}
