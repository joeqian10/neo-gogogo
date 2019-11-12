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
		Transaction:NewTransaction(),
		Claims:claims,
	}
	ctx.Type = Claim_Transaction
	return ctx
}

// HashString returns the transaction Id string
func (ctx *ClaimTransaction) HashString() string {
	hash := crypto.Hash256(ctx.UnsignedRawTransaction())
	ctx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (ctx *ClaimTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	ctx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (ctx *ClaimTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	ctx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (ctx *ClaimTransaction) RawTransactionString() string {
	return hex.EncodeToString(ctx.RawTransaction())
}

// FromHexString parses a hex string
func (ctx *ClaimTransaction) FromHexString(rawTx string) (*ClaimTransaction, error) {
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
func (ctx *ClaimTransaction) Deserialize(br *io.BinReader) {
	ctx.DeserializeUnsigned(br)
	ctx.Transaction.DeserializeWitnesses(br)
}

func (ctx *ClaimTransaction) DeserializeUnsigned(br *io.BinReader) {
	ctx.Transaction.DeserializeUnsigned1(br)
	ctx.DeserializeExclusiveData(br)
	ctx.Transaction.DeserializeUnsigned2(br)
}

func (ctx *ClaimTransaction) DeserializeExclusiveData(br *io.BinReader) {
	lenClaims := br.ReadVarUint()
	ctx.Claims = make([]*CoinReference, lenClaims)
	for i := 0; i < int(lenClaims); i++ {
		ctx.Claims[i] = &CoinReference{}
		ctx.Claims[i].Deserialize(br)
	}
}

// Serialize implements Serializable interface.
func (ctx *ClaimTransaction) Serialize(bw *io.BinWriter) {
	ctx.SerializeUnsigned(bw)
	ctx.SerializeWitnesses(bw)
}

func (ctx *ClaimTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	ctx.Transaction.SerializeUnsigned1(bw)
	ctx.SerializeExclusiveData(bw)
	ctx.SerializeUnsigned2(bw)
}

func (ctx *ClaimTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
	bw.WriteVarUint(uint64(len(ctx.Claims)))
	for _, claim := range ctx.Claims {
		claim.Serialize(bw)
	}
}

