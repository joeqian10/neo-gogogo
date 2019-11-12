package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// IssueTransaction inherits Transaction
type IssueTransaction struct {
	*Transaction
}

// NewInvocationTransaction creates an IssueTransaction
func NewIssueTransaction(script []byte) *IssueTransaction {
	itx := &IssueTransaction{
		Transaction:NewTransaction(),
	}
	itx.Type = Issue_Transaction
	return itx
}

// HashString returns the transaction Id string
func (itx *IssueTransaction) HashString() string {
	hash := crypto.Hash256(itx.UnsignedRawTransaction())
	itx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (itx *IssueTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	itx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (itx *IssueTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	itx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (itx *IssueTransaction) RawTransactionString() string {
	return hex.EncodeToString(itx.RawTransaction())
}

// FromHexString parses a hex string to get an IssueTransaction
func (itx *IssueTransaction) FromHexString(rawTx string) (*IssueTransaction, error) {
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
func (itx *IssueTransaction) Deserialize(br *io.BinReader) {
	itx.DeserializeUnsigned(br)
	itx.Transaction.DeserializeWitnesses(br)
}

func (itx *IssueTransaction) DeserializeUnsigned(br *io.BinReader) {
	itx.Transaction.DeserializeUnsigned1(br)
	itx.DeserializeExclusiveData(br)
	itx.Transaction.DeserializeUnsigned2(br)
}

func (itx *IssueTransaction) DeserializeExclusiveData(br *io.BinReader) {
}

// Serialize implements Serializable interface.
func (itx *IssueTransaction) Serialize(bw *io.BinWriter) {
	itx.SerializeUnsigned(bw)
	itx.SerializeWitnesses(bw)
}

func (itx *IssueTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	itx.Transaction.SerializeUnsigned1(bw)
	itx.SerializeExclusiveData(bw)
	itx.SerializeUnsigned2(bw)
}

func (itx *IssueTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
}

