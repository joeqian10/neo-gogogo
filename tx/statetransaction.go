package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// StateTransaction inherits Transaction
type StateTransaction struct {
	*Transaction
	Descriptors []*StateDescriptor
}

// NewStateTransaction creates an IssueTransaction
func NewStateTransaction(script []byte) *StateTransaction {
	itx := &StateTransaction{
		Transaction:NewTransaction(),
	}
	itx.Type = State_Transaction
	return itx
}

// HashString returns the transaction Id string
func (stx *StateTransaction) HashString() string {
	hash := crypto.Hash256(stx.UnsignedRawTransaction())
	stx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (stx *StateTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	stx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (stx *StateTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	stx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (stx *StateTransaction) RawTransactionString() string {
	return hex.EncodeToString(stx.RawTransaction())
}

// FromHexString parses a hex string
func (stx *StateTransaction) FromHexString(rawTx string) (*StateTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinReaderFromBuf(b)
	stx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return stx, nil
}

// Deserialize implements Serializable interface.
func (stx *StateTransaction) Deserialize(br *io.BinReader) {
	stx.DeserializeUnsigned(br)
	stx.Transaction.DeserializeWitnesses(br)
}

func (stx *StateTransaction) DeserializeUnsigned(br *io.BinReader) {
	stx.Transaction.DeserializeUnsigned1(br)
	stx.DeserializeExclusiveData(br)
	stx.Transaction.DeserializeUnsigned2(br)
}

func (stx *StateTransaction) DeserializeExclusiveData(br *io.BinReader) {
	lenDesc := br.ReadVarUint()
	stx.Descriptors = make([]*StateDescriptor, lenDesc)
	for i := 0; i < int(lenDesc); i++ {
		stx.Descriptors[i] = &StateDescriptor{}
		stx.Descriptors[i].Deserialize(br)
	}
}

// Serialize implements Serializable interface.
func (stx *StateTransaction) Serialize(bw *io.BinWriter) {
	stx.SerializeUnsigned(bw)
	stx.SerializeWitnesses(bw)
}

func (stx *StateTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	stx.Transaction.SerializeUnsigned1(bw)
	stx.SerializeExclusiveData(bw)
	stx.SerializeUnsigned2(bw)
}

func (stx *StateTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
	bw.WriteVarUint(uint64(len(stx.Descriptors)))
	for _, desc := range stx.Descriptors {
		desc.Serialize(bw)
	}
}

