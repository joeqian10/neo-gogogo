package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"unsafe"
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
		Transaction:NewTransaction(),
		Script:script,
		}
	itx.Type = Invocation_Transaction
	return itx
}

func (itx *InvocationTransaction) Size() int {
	size := unsafe.Sizeof(itx.Script) + unsafe.Sizeof(itx.Gas)
	return itx.Transaction.Size() + int(size)
}

// HashString returns the transaction hash string
func (itx *InvocationTransaction) HashString() string {
	hash := crypto.Hash256(itx.UnsignedRawTransaction()) // twice sha-256
	itx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (itx *InvocationTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	itx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (itx *InvocationTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	itx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (itx *InvocationTransaction) RawTransactionString() string {
	return hex.EncodeToString(itx.RawTransaction())
}

// FromHexString parses a hex string to get an InvocationTransaction
func (itx *InvocationTransaction) FromHexString(rawTx string) (*InvocationTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinReaderFromBuf(b)
	itx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return itx, nil
}

// Deserialize implements Serializable interface.
func (itx *InvocationTransaction) Deserialize(br *io.BinReader) {
	itx.DeserializeUnsigned(br)
	itx.DeserializeWitnesses(br)
}

func (itx *InvocationTransaction) DeserializeUnsigned(br *io.BinReader) {
	itx.DeserializeUnsigned1(br)
	itx.DeserializeExclusiveData(br)
	itx.DeserializeUnsigned2(br)
}

func (itx *InvocationTransaction) DeserializeExclusiveData(br *io.BinReader) {
	itx.Script = br.ReadBytes()
	if itx.Version >= 1 {
		br.ReadLE(&itx.Gas)
	} else {
		itx.Gas = helper.Fixed8FromInt64(0)
	}
}

// Serialize implements Serializable interface.
func (itx *InvocationTransaction) Serialize(bw *io.BinWriter) {
	itx.SerializeUnsigned(bw)
	itx.SerializeWitnesses(bw)
}

func (itx *InvocationTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	itx.SerializeUnsigned1(bw)
	itx.SerializeExclusiveData(bw)
	itx.SerializeUnsigned2(bw)
}

func (itx *InvocationTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
	bw.WriteBytes(itx.Script)
	if itx.Version >= 1 {
		bw.WriteLE(itx.Gas)
	}
}

