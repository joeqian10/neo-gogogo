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
		Transaction:NewTransaction(),
	}
	ctx.Type = Contract_Transaction
	return ctx
}

// HashString returns the transaction Id string
func (ctx *ContractTransaction) HashString() string {
	hash := crypto.Hash256(ctx.UnsignedRawTransaction())
	ctx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (ctx *ContractTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	ctx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (ctx *ContractTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	ctx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (ctx *ContractTransaction) RawTransactionString() string {
	return hex.EncodeToString(ctx.RawTransaction())
}

// FromHexString parses a hex string
func (ctx *ContractTransaction) FromHexString(rawTx string) (*ContractTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinReaderFromBuf(b)
	ctx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return ctx, nil
}

// Deserialize implements Serializable interface.
func (ctx *ContractTransaction) Deserialize(br *io.BinReader) {
	ctx.DeserializeUnsigned(br)
	ctx.Transaction.DeserializeWitnesses(br)
}

func (ctx *ContractTransaction) DeserializeUnsigned(br *io.BinReader) {
	ctx.Transaction.DeserializeUnsigned1(br)
	ctx.DeserializeExclusiveData(br)
	ctx.Transaction.DeserializeUnsigned2(br)
}

func (ctx *ContractTransaction) DeserializeExclusiveData(br *io.BinReader) {
}

// Serialize implements Serializable interface.
func (ctx *ContractTransaction) Serialize(bw *io.BinWriter) {
	ctx.SerializeUnsigned(bw)
	ctx.SerializeWitnesses(bw)
}

func (ctx *ContractTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	ctx.Transaction.SerializeUnsigned1(bw)
	ctx.SerializeExclusiveData(bw)
	ctx.SerializeUnsigned2(bw)
}

func (ctx *ContractTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
}

