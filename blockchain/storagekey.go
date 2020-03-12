package blockchain

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

//Storagekey key use to store StorageItem on blockchain
type Storagekey struct {
	ScriptHash helper.UInt160
	Key        []byte
}

//Deserialize deserialize from byte array
func (sk *Storagekey) Deserialize(reader *io.BinaryReader) (StorageItem, error) {
	sk.ScriptHash = helper.UInt160FromBytes(reader.ReadVarBytes())
	sk.Key = reader.ReadVarBytes()
}

//Serialize serialize to byte array
func (sk *Storagekey) Serialize(writer *io.BinaryWriter) []byte {
	writer.WriteVarBytes(sk.ScriptHash.Bytes())
	writer.WriteVarBytes(sk.Key)
}
