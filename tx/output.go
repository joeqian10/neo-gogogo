package transaction

import (
	"../helper"
	"../helper/io"
)

// Output represents a Transaction output.
type Output struct {
	// The NEO asset id used in the transaction.
	AssetID helper.UInt256

	// Amount of AssetType send or received.
	Amount helper.Fixed8

	// The address of the recipient.
	ScriptHash helper.UInt160

	// The position of the Output in slice []Output. This is actually set in NewTransactionOutputRaw
	// and used for displaying purposes.
	Position int
}

// NewOutput returns a new transaction output.
func NewOutput(assetID helper.UInt256, amount helper.Fixed8, scriptHash helper.UInt160) *Output {
	return &Output{
		AssetID:    assetID,
		Amount:     amount,
		ScriptHash: scriptHash,
	}
}

// DecodeBinary implements Serializable interface.
func (out *Output) DecodeBinary(br *io.BinReader) {
	br.ReadLE(&out.AssetID)
	br.ReadLE(&out.Amount)
	br.ReadLE(&out.ScriptHash)
}

// EncodeBinary implements Serializable interface.
func (out *Output) EncodeBinary(bw *io.BinWriter) {
	bw.WriteLE(out.AssetID)
	bw.WriteLE(out.Amount)
	bw.WriteLE(out.ScriptHash)
}

// MarshalJSON implements the Marshaler interface.
//func (out *Output) MarshalJSON() ([]byte, error) {
//	return json.Marshal(map[string]interface{}{
//		"asset":   out.AssetID,
//		"value":   out.Amount,
//		"address": AddressFromUint160(out.ScriptHash),
//		"n":       out.Position,
//	})
//}
