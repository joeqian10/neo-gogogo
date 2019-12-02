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

// NewIssueTransaction creates an IssueTransaction
func NewIssueTransaction(script []byte) *IssueTransaction {
	tx := &IssueTransaction{
		Transaction:NewTransaction(),
	}
	tx.Type = Issue_Transaction
	return tx
}

// HashString returns the transaction Id string
func (tx *IssueTransaction) HashString() string {
	hash := crypto.Hash256(tx.UnsignedRawTransaction())
	tx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (tx *IssueTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	tx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *IssueTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	tx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *IssueTransaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

// FromHexString parses a hex string to get an IssueTransaction
func (tx *IssueTransaction) FromHexString(rawTx string) (*IssueTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinReaderFromBuf(b)
	tx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return tx, nil
}

// Deserialize implements Serializable interface.
func (tx *IssueTransaction) Deserialize(br *io.BinReader) {
	tx.DeserializeUnsigned(br)
	tx.Transaction.DeserializeWitnesses(br)
}

func (tx *IssueTransaction) DeserializeUnsigned(br *io.BinReader) {
	tx.Transaction.DeserializeUnsigned1(br)
	tx.DeserializeExclusiveData(br)
	tx.Transaction.DeserializeUnsigned2(br)
}

func (tx *IssueTransaction) DeserializeExclusiveData(br *io.BinReader) {
}

// Serialize implements Serializable interface.
func (tx *IssueTransaction) Serialize(bw *io.BinWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *IssueTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	tx.Transaction.SerializeUnsigned1(bw)
	tx.SerializeExclusiveData(bw)
	tx.SerializeUnsigned2(bw)
}

func (tx *IssueTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
}

