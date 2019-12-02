package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// ContractTransaction inherits Transaction
type ContractTransaction struct {
	*Transaction
}

// CreateContractTransaction creates a contract transaction
func NewContractTransaction() *ContractTransaction {
	ctx := &ContractTransaction{
		Transaction: NewTransaction(),
	}
	ctx.Type = Contract_Transaction
	return ctx
}

func (tx *ContractTransaction) Size() int {
	return len(tx.RawTransaction())
}

// implement ITransaction interface
func (tx *ContractTransaction) GetTransaction() *Transaction {
	return tx.Transaction
}

// HashString returns the transaction Id string
func (tx *ContractTransaction) HashString() string {
	hash := crypto.Hash256(tx.UnsignedRawTransaction())
	tx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (tx *ContractTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *ContractTransaction) RawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *ContractTransaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

// FromHexString parses a hex string
func (tx *ContractTransaction) FromHexString(rawTx string) (*ContractTransaction, error) {
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
func (tx *ContractTransaction) Deserialize(br *io.BinaryReader) {
	tx.DeserializeUnsigned(br)
	tx.Transaction.DeserializeWitnesses(br)
}

func (tx *ContractTransaction) DeserializeUnsigned(br *io.BinaryReader) {
	tx.Transaction.DeserializeUnsigned1(br)
	tx.DeserializeExclusiveData(br)
	tx.Transaction.DeserializeUnsigned2(br)
}

func (tx *ContractTransaction) DeserializeExclusiveData(br *io.BinaryReader) {
}

// Serialize implements Serializable interface.
func (tx *ContractTransaction) Serialize(bw *io.BinaryWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *ContractTransaction) SerializeUnsigned(bw *io.BinaryWriter)  {
	tx.SerializeUnsigned1(bw)
	tx.SerializeExclusiveData(bw)
	tx.SerializeUnsigned2(bw)
}

func (tx *ContractTransaction) SerializeExclusiveData(bw *io.BinaryWriter) {
}
