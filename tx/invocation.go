package transaction

import (
	"../helper/io"
	"github.com/joeqian10/neo-gogogo/helper"
)

// InvocationTX represents a invocation transaction and is used to
// deploy smart contract to the NEO blockchain.
type InvocationTX struct {
	// Script output of the smart contract.
	Script []byte

	// Gas cost of the smart contract.
	Gas     helper.Fixed8
	Version uint8
}

// NewInvocationTX returns a new invocation transaction.
func NewInvocationTX(script []byte) *Transaction {
	return &Transaction{
		Type:    InvocationType,
		Version: 1,
		Data: &InvocationTX{
			Script:  script,
			Version: 1,
		},
		Attributes: []*Attribute{},
		Inputs:     []*Input{},
		Outputs:    []*Output{},
		Scripts:    []*Witness{},
	}
}

// DecodeBinary implements Serializable interface.
func (tx *InvocationTX) DecodeBinary(br *io.BinReader) {
	tx.Script = br.ReadBytes()
	if tx.Version >= 1 {
		br.ReadLE(&tx.Gas)
	} else {
		tx.Gas = helper.Fixed8FromInt64(0)
	}
}

// EncodeBinary implements Serializable interface.
func (tx *InvocationTX) EncodeBinary(bw *io.BinWriter) {
	bw.WriteBytes(tx.Script)
	if tx.Version >= 1 {
		bw.WriteLE(tx.Gas)
	}
}
