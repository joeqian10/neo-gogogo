package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// ClaimTransaction inherits Transaction
type ClaimTransaction struct {
	*Transaction
	Claims []*CoinReference
}

// NewClaimTransaction creates an ClaimTransaction
func NewClaimTransaction(claims []*CoinReference) *ClaimTransaction {
	ctx := &ClaimTransaction{
		Transaction: NewTransaction(),
		Claims:      claims,
	}
	ctx.Type = Claim_Transaction
	return ctx
}

func (tx *ClaimTransaction) Size() int {
	return len(tx.RawTransaction())
}

// implement ITransaction interface
func (tx *ClaimTransaction) GetTransaction() *Transaction {
	return tx.Transaction
}

// HashString returns the transaction Id string
func (tx *ClaimTransaction) HashString() string {
	hash := crypto.Hash256(tx.UnsignedRawTransaction())
	tx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (tx *ClaimTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	tx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *ClaimTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	tx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (tx *ClaimTransaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

// FromHexString parses a hex string
func (tx *ClaimTransaction) FromHexString(rawTx string) (*ClaimTransaction, error) {
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
func (tx *ClaimTransaction) Deserialize(br *io.BinReader) {
	tx.DeserializeUnsigned(br)
	tx.Transaction.DeserializeWitnesses(br)
}

func (tx *ClaimTransaction) DeserializeUnsigned(br *io.BinReader) {
	tx.Transaction.DeserializeUnsigned1(br)
	tx.DeserializeExclusiveData(br)
	tx.Transaction.DeserializeUnsigned2(br)
}

func (tx *ClaimTransaction) DeserializeExclusiveData(br *io.BinReader) {
	lenClaims := br.ReadVarUint()
	tx.Claims = make([]*CoinReference, lenClaims)
	for i := 0; i < int(lenClaims); i++ {
		tx.Claims[i] = &CoinReference{}
		tx.Claims[i].Deserialize(br)
	}
}

// Serialize implements Serializable interface.
func (tx *ClaimTransaction) Serialize(bw *io.BinWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *ClaimTransaction) SerializeUnsigned(bw *io.BinWriter) {
	tx.Transaction.SerializeUnsigned1(bw)
	tx.SerializeExclusiveData(bw)
	tx.SerializeUnsigned2(bw)
}

func (tx *ClaimTransaction) SerializeExclusiveData(bw *io.BinWriter) {
	bw.WriteVarUint(uint64(len(tx.Claims)))
	for _, claim := range tx.Claims {
		claim.Serialize(bw)
	}
}
