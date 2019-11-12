package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// TransactionInput alias CoinReference, UTXO
type CoinReference struct {
	PrevHash  helper.UInt256
	PrevIndex uint16
}

// Deserialize implements Serializable interface.
func (in *CoinReference) Deserialize(br *io.BinReader) {
	br.ReadLE(&in.PrevHash)
	br.ReadLE(&in.PrevIndex)
}

// Serialize implements Serializable interface.
func (in *CoinReference) Serialize(bw *io.BinWriter) {
	bw.WriteLE(in.PrevHash)
	bw.WriteLE(in.PrevIndex)
}