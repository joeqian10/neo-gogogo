package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// InvocationTransaction inherits Transaction
type InvocationTransaction struct {
	*Transaction
	Script []byte
	Gas    helper.Fixed8
}

// NewInvocationTransaction creates an InvocationTransaction
func NewInvocationTransaction(script []byte) *InvocationTransaction {
	itx := &InvocationTransaction{
		Transaction: NewTransaction(),
		Script:      script,
	}
	itx.Type = Invocation_Transaction
	itx.Version = 1
	return itx
}

func (tx *InvocationTransaction) Size() int {
	return len(tx.RawTransaction())
}

// implement ITransaction interface
func (tx *InvocationTransaction) GetTransaction() *Transaction {
	return tx.Transaction
}

// HashString returns the transaction hash string
func (tx *InvocationTransaction) HashString() string {
	hash := crypto.Hash256(tx.UnsignedRawTransaction()) // twice sha-256
	tx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (tx *InvocationTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *InvocationTransaction) RawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *InvocationTransaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

// FromHexString parses a hex string to get an InvocationTransaction
func (tx *InvocationTransaction) FromHexString(rawTx string) (*InvocationTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinaryReaderFromBuf(b)
	tx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return tx, nil
}

// Deserialize implements Serializable interface.
func (tx *InvocationTransaction) Deserialize(br *io.BinaryReader) {
	tx.DeserializeUnsigned(br)
	tx.DeserializeWitnesses(br)
}

func (tx *InvocationTransaction) DeserializeUnsigned(br *io.BinaryReader) {
	tx.DeserializeUnsigned1(br)
	tx.DeserializeExclusiveData(br)
	tx.DeserializeUnsigned2(br)
}

func (tx *InvocationTransaction) DeserializeExclusiveData(br *io.BinaryReader) {
	tx.Script = br.ReadBytes()
	if tx.Version >= 1 {
		br.ReadLE(&tx.Gas)
	} else {
		tx.Gas = helper.Fixed8FromInt64(0)
	}
}

// Serialize implements Serializable interface.
func (tx *InvocationTransaction) Serialize(bw *io.BinaryWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *InvocationTransaction) SerializeUnsigned(bw *io.BinaryWriter) {
	tx.SerializeUnsigned1(bw)
	tx.SerializeExclusiveData(bw)
	tx.SerializeUnsigned2(bw)
}

func (tx *InvocationTransaction) SerializeExclusiveData(bw *io.BinaryWriter) {
	bw.WriteBytes(tx.Script)
	if tx.Version >= 1 {
		bw.WriteLE(tx.Gas)
	}
}
