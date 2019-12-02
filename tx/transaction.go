package tx

import (
	"bytes"
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

func (t *Transaction) DeserializeUnsigned1(br *io.BinaryReader) {
	br.ReadLE(&t.Type)
	br.ReadLE(&t.Version)
}

func (t *Transaction) DeserializeUnsigned2(br *io.BinaryReader) {
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

func (t *Transaction) DeserializeWitnesses(br *io.BinaryReader) {
	lenWitnesses := br.ReadVarUint()
	t.Witnesses = make([]*Witness, lenWitnesses)
	for i := 0; i < int(lenWitnesses); i++ {
		t.Witnesses[i] = &Witness{}
		t.Witnesses[i].Deserialize(br)
	}
}

func (t *Transaction) SerializeUnsigned1(bw *io.BinaryWriter) {
	bw.WriteLE(t.Type)
	bw.WriteLE(t.Version)
}

func (t *Transaction) SerializeUnsigned2(bw *io.BinaryWriter) {
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

func (t *Transaction) SerializeWitnesses(bw *io.BinaryWriter) {
	bw.WriteVarUint(uint64(len(t.Witnesses)))
	for _, s := range t.Witnesses {
		s.Serialize(bw)
	}
}

// add Attribute if the script hash is not in transaction attributes
// if you want to sign a transaction with a scriptHash which not in inputs, you need to add the scriptHash into transaction attributes
func (t *Transaction) AddScriptHashToAttribute(scriptHash helper.UInt160) {
	i := 0
	for ; i < len(t.Attributes); i++ {
		attr := t.Attributes[i]
		if attr.Usage == Script && bytes.Equal(attr.Data, scriptHash.Bytes()) {
			break
		}
	}
	if i == len(t.Attributes) {
		t.Attributes = append(t.Attributes, &TransactionAttribute{Usage: Script, Data: scriptHash.Bytes()})
	}
}

type ITransaction interface {
	GetTransaction() *Transaction
	UnsignedRawTransaction() []byte
}

// add signature for ITransaction
func AddSignature(transaction ITransaction, key *keys.KeyPair) error {
	scriptHash := key.PublicKey.ScriptHash()
	tx := transaction.GetTransaction()
	for _, witness := range tx.Witnesses {
		// the transaction has been signed with this KeyPair
		if witness.scriptHash == scriptHash {
			return nil
		}
	}

	// if the transaction is not signed before, just add script hash to transaction attributes
	if len(tx.Witnesses) == 0 {
		tx.AddScriptHashToAttribute(scriptHash)
	}

	// create witness
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
	if err != nil {
		return err
	}

	scriptHash, err := helper.UInt160FromBytes(crypto.Hash160(script))
	if err != nil {
		return err
	}

	for _, witness := range tx.Witnesses {
		// the transaction has been signed with this KeyPair
		if witness.scriptHash == scriptHash {
			return nil
		}
	}

	// if the transaction is not signed before, just add script hash to transaction attributes
	if len(tx.Witnesses) == 0 {
		tx.AddScriptHashToAttribute(scriptHash)
	}

	// create witness
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
