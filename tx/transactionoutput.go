package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// TransactionOutput
type TransactionOutput struct {
	AssetId    helper.UInt256
	Value      helper.Fixed8
	ScriptHash helper.UInt160
}

// NewTransactionOutput returns a new transaction output.
func NewTransactionOutput(assetID helper.UInt256, value helper.Fixed8, scriptHash helper.UInt160) *TransactionOutput {
	return &TransactionOutput{
		AssetId:    assetID,
		Value:     value,
		ScriptHash: scriptHash,
	}
}

// Deserialize implements Serializable interface.
func (out *TransactionOutput) Deserialize(br *io.BinReader) {
	br.ReadLE(&out.AssetId)
	br.ReadLE(&out.Value)
	br.ReadLE(&out.ScriptHash)
}

// Serialize implements Serializable interface.
func (out *TransactionOutput) Serialize(bw *io.BinWriter) {
	bw.WriteLE(out.AssetId)
	bw.WriteLE(out.Value)
	bw.WriteLE(out.ScriptHash)
}
