package tx

import (
	"bytes"
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
)

// InvocationTransaction inherits transaction
type InvocationTransaction struct {
	Transaction
	Script []byte
	Gas    helper.Fixed8
}

// CreateInvocationTransaction creates an InvocationTransaction
func CreateInvocationTransaction() *InvocationTransaction {
	itx := &InvocationTransaction{
		Transaction{
			Type:    Invocation_Transaction,
			Version: TransactionVersion,
		},
		*new([]byte),
		helper.NewFixed8(0),
	}
	return itx
}

func (itx *InvocationTransaction) UnsignedRawTransaction() []byte {
	buff := new(bytes.Buffer)
	buff.Write(itx.UnsignedRawTransactionPart1())
	scriptLength := helper.VarIntFromInt(len(itx.Script))
	buff.Write(scriptLength.Bytes())
	buff.Write(itx.Script)
	buff.Write(helper.VarIntFromUInt64(uint64(itx.Gas.Value)).Bytes())
	buff.Write(itx.UnsignedRawTransactionPart2())
	return buff.Bytes()
}

func (itx *InvocationTransaction) RawTransaction() []byte {
	return append(itx.UnsignedRawTransaction(), itx.SerializeWitnesses()...)
}

//
func (itx *InvocationTransaction) RawTransactionString() string {
	return hex.EncodeToString(itx.RawTransaction())
}

func (itx *InvocationTransaction) TXID() string {
	txid := crypto.Sha256(itx.UnsignedRawTransaction())
	return hex.EncodeToString(helper.ReverseBytes(txid))
}
