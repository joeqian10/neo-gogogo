package tx

import (
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
	"sort"
)

const (
	TransactionVersion uint8 = 1 // neo-2.x
	MaxTransactionSize       = 102400
)

// base class
type Transaction struct {
	Type       TransactionType
	Version    uint8
	Hash       helper.UInt256
	Attributes []*TransactionAttribute
	Inputs     []*CoinReference
	Outputs    []*TransactionOutput
	Witnesses  []*Witness
	//ExclusiveData io.Serializable
}

func NewTransaction() *Transaction {
	return &Transaction{
		Version:    TransactionVersion,
		Attributes: []*TransactionAttribute{},
		Inputs:     []*CoinReference{},
		Outputs:    []*TransactionOutput{},
		Witnesses:  []*Witness{},
	}
}

func (t *Transaction) DeserializeUnsigned1(br *io.BinReader) {
	br.ReadLE(&t.Type)
	br.ReadLE(&t.Version)
}

func (t *Transaction) DeserializeUnsigned2(br *io.BinReader) {
	// Attributes
	lenAttributes := br.ReadVarUint()
	t.Attributes = make([]*TransactionAttribute, lenAttributes)
	for i := 0; i < int(lenAttributes); i++ {
		t.Attributes[i] = &TransactionAttribute{}
		t.Attributes[i].Deserialize(br)
	}
	// Inputs
	lenInputs := br.ReadVarUint()
	t.Inputs = make([]*CoinReference, lenInputs)
	for i := 0; i < int(lenInputs); i++ {
		t.Inputs[i] = &CoinReference{}
		t.Inputs[i].Deserialize(br)
	}
	// Outputs
	lenOutputs := br.ReadVarUint()
	t.Outputs = make([]*TransactionOutput, lenOutputs)
	for i := 0; i < int(lenOutputs); i++ {
		t.Outputs[i] = &TransactionOutput{}
		t.Outputs[i].Deserialize(br)
	}
}

func (t *Transaction) DeserializeWitnesses(br *io.BinReader) {
	lenWitnesses := br.ReadVarUint()
	t.Witnesses = make([]*Witness, lenWitnesses)
	for i := 0; i < int(lenWitnesses); i++ {
		t.Witnesses[i] = &Witness{}
		t.Witnesses[i].Deserialize(br)
	}
}

func (t *Transaction) SerializeUnsigned1(bw *io.BinWriter) {
	bw.WriteLE(t.Type)
	bw.WriteLE(t.Version)
}

func (t *Transaction) SerializeUnsigned2(bw *io.BinWriter) {
	// Attributes
	bw.WriteVarUint(uint64(len(t.Attributes)))
	for _, attr := range t.Attributes {
		attr.Serialize(bw)
	}
	// Inputs
	bw.WriteVarUint(uint64(len(t.Inputs)))
	for _, in := range t.Inputs {
		in.Serialize(bw)
	}
	// Outputs
	bw.WriteVarUint(uint64(len(t.Outputs)))
	for _, out := range t.Outputs {
		out.Serialize(bw)
	}
}

func (t *Transaction) SerializeWitnesses(bw *io.BinWriter) {
	bw.WriteVarUint(uint64(len(t.Witnesses)))
	for _, s := range t.Witnesses {
		s.Serialize(bw)
	}
}

type ITransaction interface {
	GetTransaction() *Transaction
	UnsignedRawTransaction() []byte
}

// add signature for ITransaction
// TODO add attribute script hash if the KeyPair is not in input
func AddSignature(transaction ITransaction, key *keys.KeyPair) error {
	tx := transaction.GetTransaction()
	for _, witness := range tx.Witnesses {
		// the transaction has been signed with this KeyPair
		if witness.scriptHash == key.PublicKey.ScriptHash() {
			return nil
		}
	}

	witness, err := CreateSignatureWitness(transaction.UnsignedRawTransaction(), key)
	if err != nil {
		return err
	}
	tx.Witnesses = append(tx.Witnesses, witness)
	sort.Slice(tx.Witnesses, func(i, j int) bool {
		return tx.Witnesses[i].scriptHash.Less(tx.Witnesses[j].scriptHash)
	})
	return nil
}

// add multi-signature for ITransaction
func AddMultiSignature(transaction ITransaction, pairs []*keys.KeyPair, m int, publicKeys []*keys.PublicKey) error {
	tx := transaction.GetTransaction()
	script, err := keys.CreateMultiSigRedeemScript(m, publicKeys...)
	scriptHash, err := helper.UInt160FromBytes(crypto.Hash160(script))
	for _, witness := range tx.Witnesses {
		// the transaction has been signed with this KeyPair
		if witness.scriptHash == scriptHash {
			return nil
		}
	}

	witness, err := CreateMultiSignatureWitness(transaction.UnsignedRawTransaction(), pairs, m, publicKeys)
	if err != nil {
		return err
	}
	tx.Witnesses = append(tx.Witnesses, witness)
	sort.Slice(tx.Witnesses, func(i, j int) bool {
		return tx.Witnesses[i].scriptHash.Less(tx.Witnesses[j].scriptHash)
	})
	return nil
}
