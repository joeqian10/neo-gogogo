package blockchain

import "github.com/joeqian10/neo-gogogo/helper/io"

//StorageItem value to store on blockchain
type StorageItem struct {
	Version   byte
	Value     []byte
	IsContant bool
}

//Deserialize deserialize from byte array
func (si *StorageItem) Deserialize(reader *io.BinaryReader) (StorageItem, error) {
	reader.ReadLE(&si.Version)
	si.Value = reader.ReadVarBytes()
	reader.ReadLE(&si.IsContant)
}

//Serialize serialize to byte array
func (si *StorageItem) Serialize(writer *io.BinaryWriter) []byte {
	writer.WriteLE(si.Version)
	writer.WriteVarBytes(si.Value)
	writer.WriteLE(si.IsContant)
}
