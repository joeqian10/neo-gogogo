package blockchain

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// StorageKey key use to store StorageItem on blockchain
type StorageKey struct {
	ScriptHash helper.UInt160
	Key        []byte
}

// Deserialize deserializes from byte array
func (sk *StorageKey) Deserialize(reader *io.BinaryReader) {
	reader.ReadLE(&sk.ScriptHash)
	sk.Key, _ = reader.ReadBytesWithGrouping()
}

// Serialize serializes to byte array
func (sk *StorageKey) Serialize(writer *io.BinaryWriter) {
	writer.WriteLE(sk.ScriptHash)
	writer.WriteBytesWithGrouping(sk.Key)
}
