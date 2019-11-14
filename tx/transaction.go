package tx

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"unsafe"
)

const (
	TransactionVersion uint8 = 1 // neo-2.x
	MaxTransactionSize = 102400
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

func (t *Transaction) Size() int {
	size := unsafe.Sizeof(t.Type) + unsafe.Sizeof(t.Version) + unsafe.Sizeof(t.Attributes) + unsafe.Sizeof(t.Inputs) + unsafe.Sizeof(t.Outputs) + unsafe.Sizeof(t.Witnesses)
	return int(size)
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
