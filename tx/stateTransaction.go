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

// NewStateTransaction creates a StateTransaction
func NewStateTransaction(script []byte) *StateTransaction {
	tx := &StateTransaction{
		Transaction:NewTransaction(),
	}
	tx.Type = State_Transaction
	return tx
}

// HashString returns the transaction Id string
func (tx *StateTransaction) HashString() string {
	hash := crypto.Hash256(tx.UnsignedRawTransaction())
	tx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (tx *StateTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *StateTransaction) RawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *StateTransaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

// FromHexString parses a hex string
func (tx *StateTransaction) FromHexString(rawTx string) (*StateTransaction, error) {
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
func (tx *StateTransaction) Deserialize(br *io.BinaryReader) {
	tx.DeserializeUnsigned(br)
	tx.Transaction.DeserializeWitnesses(br)
}

func (tx *StateTransaction) DeserializeUnsigned(br *io.BinaryReader) {
	tx.Transaction.DeserializeUnsigned1(br)
	tx.DeserializeExclusiveData(br)
	tx.Transaction.DeserializeUnsigned2(br)
}

func (tx *StateTransaction) DeserializeExclusiveData(br *io.BinaryReader) {
	lenDesc := br.ReadVarUint()
	tx.Descriptors = make([]*StateDescriptor, lenDesc)
	for i := 0; i < int(lenDesc); i++ {
		tx.Descriptors[i] = &StateDescriptor{}
		tx.Descriptors[i].Deserialize(br)
	}
}

// Serialize implements Serializable interface.
func (tx *StateTransaction) Serialize(bw *io.BinaryWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *StateTransaction) SerializeUnsigned(bw *io.BinaryWriter)  {
	tx.Transaction.SerializeUnsigned1(bw)
	tx.SerializeExclusiveData(bw)
	tx.SerializeUnsigned2(bw)
}

func (tx *StateTransaction) SerializeExclusiveData(bw *io.BinaryWriter)  {
	bw.WriteVarUint(uint64(len(tx.Descriptors)))
	for _, desc := range tx.Descriptors {
		desc.Serialize(bw)
	}
}

