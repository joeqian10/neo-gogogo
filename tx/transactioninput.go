package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/rpc/models"
)

// TransactionInput alias CoinReference, UTXO
type CoinReference struct {
	PrevHash  helper.UInt256
	PrevIndex uint16
}

// Deserialize implements Serializable interface.
func (in *CoinReference) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&in.PrevHash)
	br.ReadLE(&in.PrevIndex)
}

// Serialize implements Serializable interface.
func (in *CoinReference) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(in.PrevHash)
	bw.WriteLE(in.PrevIndex)
}

func ToCoinReference(u models.Unspent) *CoinReference {
	h, _:= helper.UInt256FromString(u.Txid)
	return &CoinReference{
		PrevHash:  h,
		PrevIndex: uint16(u.N),
	}
}