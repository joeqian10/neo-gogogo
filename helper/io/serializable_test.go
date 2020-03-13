package io

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSerialzable struct {
	Flag  bool
	Value []byte
}

func (t *TestSerialzable) Serialize(writer *BinaryWriter) {
	writer.WriteLE(t.Flag)
	writer.WriteVarBytes(t.Value)
}

func (t *TestSerialzable) Deserialize(reader *BinaryReader) {
	reader.ReadLE(&t.Flag)
	t.Value = reader.ReadVarBytes()
}

func TestToArray(t *testing.T) {
	ts := &TestSerialzable{
		Flag:  true,
		Value: helper.HexToBytes("abcd"),
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
