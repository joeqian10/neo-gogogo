package block

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SetupBlockHeaderWithValues() *BlockHeader {
	mr, _ := helper.UInt256FromBytes([]byte{214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251})
	//ts := time.Date(1968, 06, 01, 0, 0, 0, 0, time.UTC)

	bh := BlockHeader{
		Version:       0,
		PrevHash:      helper.UInt256{},
		MerkleRoot:    mr,
		Timestamp:     4244941696,
		Index:         0,
		ConsensusData: 30,
		NextConsensus: helper.UInt160{},
		Witness: &tx.Witness{
			InvocationScript:   []byte{},
			VerificationScript: []byte{byte(sc.PUSHT)},
		},
	}
	return &bh
}

func TestBlockHeader_Deserialize(t *testing.T) {
	rawBlock := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		30, 0, 0, 0, 0, 0, 0, 0, // ConsensusData
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 81} // Witness

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := &BlockHeader{
		Witness: &tx.Witness{},
	}
	bh.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, uint32(0), bh.Version)
	assert.Equal(t, uint32(4244941696), bh.Timestamp)
	assert.Equal(t, uint64(30), bh.ConsensusData)
	assert.Equal(t, byte(sc.PUSHT), bh.Witness.VerificationScript[0])
}

func TestBlockHeader_DeserializeUnsigned(t *testing.T) {
	rawBlock := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		30, 0, 0, 0, 0, 0, 0, 0, // ConsensusData
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 81} // Witness

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := &BlockHeader{
		Witness: &tx.Witness{},
	}
	bh.DeserializeUnsigned(br)
	assert.Equal(t, uint32(0), bh.Version)
	assert.Equal(t, uint32(4244941696), bh.Timestamp)
	assert.Equal(t, uint64(30), bh.ConsensusData)
}

func TestBlockHeader_GetHashData(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	bs := bh.GetHashData()
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		30, 0, 0, 0, 0, 0, 0, 0, // ConsensusData
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // NextConsensus

	assert.Equal(t, requiredData, bs)
}

func TestBlockHeader_Serialize(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	buf := io.NewBufBinaryWriter()
	bh.Serialize(buf.BinaryWriter)
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		30, 0, 0, 0, 0, 0, 0, 0, // ConsensusData
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 81} // Witness

	assert.Nil(t, buf.Err)
	assert.Equal(t, helper.BytesToHex(requiredData), helper.BytesToHex(buf.Bytes()))
}

func TestBlockHeader_SerializeUnsigned(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	buf := io.NewBufBinaryWriter()
	bh.SerializeUnsigned(buf.BinaryWriter)
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		30, 0, 0, 0, 0, 0, 0, 0, // ConsensusData
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // NextConsensus

	assert.Nil(t, buf.Err)
	assert.Equal(t, helper.BytesToHex(requiredData), helper.BytesToHex(buf.Bytes()))
}
