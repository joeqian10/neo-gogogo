package block

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/rpc/models"
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
}

func NewBlockHeaderFromRPC(header *models.RpcBlockHeader) (*BlockHeader, error) {
	version := uint32(header.Version)
	prevHash, err := helper.UInt256FromString(header.PreviousBlockHash)
	if err != nil {
		return nil, err
	}
	merkleRoot, err := helper.UInt256FromString(header.MerkleRoot)
	if err != nil {
		return nil, err
	}
	timeStamp := uint32(header.Time)
	index := uint32(header.Index)
	consensusData := binary.BigEndian.Uint64(helper.HexToBytes(header.Nonce)) // Nonce is in big endian
	nextConsensus, err := helper.AddressToScriptHash(header.NextConsensus)
	if err != nil {
		return nil, err
	}
	witness := &tx.Witness{
		InvocationScript:   helper.HexToBytes(header.Witness.InvocationScript),
		VerificationScript: helper.HexToBytes(header.Witness.VerificationScript),
	}
	hash, err := helper.UInt256FromString(header.Hash)
	if err != nil {
		return nil, err
	}
	bh := BlockHeader{
		Version:       version,
		PrevHash:      prevHash,
		MerkleRoot:    merkleRoot,
		Timestamp:     timeStamp,
		Index:         index,
		ConsensusData: consensusData,
		NextConsensus: nextConsensus,
		Witness:       witness,
		_hash:         hash,
	}
	return &bh, nil
}

func (bh *BlockHeader) Deserialize(br *io.BinaryReader) {
	bh.DeserializeUnsigned(br)
	var b byte
	br.ReadLE(&b)
	if b != byte(1) {
		br.Err = fmt.Errorf("format error: padding must equal 1 got %d", b)
	}
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

func (bh *BlockHeader) HashString() string {
	hash := crypto.Hash256(bh.GetHashData())
	bh._hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}
