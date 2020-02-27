package block

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/tx"
)

type BlockHeader struct {
	Version       uint32
	PrevHash      helper.UInt256
	MerkleRoot    helper.UInt256
	Timestamp     uint32
	Index         uint32
	ConsensusData uint64
	NextConsensus helper.UInt160
	Witness       *tx.Witness

	_hash helper.UInt256
	_size int
	// cross chain support, todo
	//CrossStatesRoot string
}

func (bh *BlockHeader) Deserialize(br *io.BinaryReader) {
	bh.DeserializeUnsigned(br)
	var b byte
	br.ReadLE(&b)
	if b != byte(1) {br.Err = fmt.Errorf("format error: padding must equal 1 got %d", b)}
	bh.Witness.Deserialize(br)
}

//DeserializeUnsigned deserialize blockheader without witness
func (bh *BlockHeader) DeserializeUnsigned(br *io.BinaryReader) {
	br.ReadLE(&bh.Version)
	br.ReadLE(&bh.PrevHash)
	br.ReadLE(&bh.MerkleRoot)
	br.ReadLE(&bh.Timestamp)
	br.ReadLE(&bh.Index)
	br.ReadLE(&bh.ConsensusData)
	br.ReadLE(&bh.NextConsensus)
}


func (bh *BlockHeader) Serialize(bw *io.BinaryWriter) {
	bh.SerializeUnsigned(bw)
	bw.WriteLE(byte(1))
	bh.Witness.Serialize(bw)
}

//SerializeUnsigned serialize blockheader without witness
func (bh *BlockHeader) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(bh.Version)
	bw.WriteLE(bh.PrevHash)
	bw.WriteLE(bh.MerkleRoot)
	bw.WriteLE(bh.Timestamp)
	bw.WriteLE(bh.Index)
	bw.WriteLE(bh.ConsensusData)
	bw.WriteLE(bh.NextConsensus)
}

func (bh *BlockHeader) GetHashData() []byte {
	buf := io.NewBufBinaryWriter()
	bh.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}