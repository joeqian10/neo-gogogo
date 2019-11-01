package tx

import (
	"bytes"
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
)

type TransactionAttribute struct {
	Usage TransactionAttributeUsage
	Data  []byte
}

// Transaction attribute usages
type TransactionAttributeUsage byte

const (
	ContractHash TransactionAttributeUsage = 0x00

	ECDH02 TransactionAttributeUsage = 0x02
	ECDH03 TransactionAttributeUsage = 0x03

	Script TransactionAttributeUsage = 0x20

	Vote TransactionAttributeUsage = 0x30

	DescriptionUrl TransactionAttributeUsage = 0x81
	Description    TransactionAttributeUsage = 0x90

	Hash1  TransactionAttributeUsage = 0xa1
	Hash2  TransactionAttributeUsage = 0xa2
	Hash3  TransactionAttributeUsage = 0xa3
	Hash4  TransactionAttributeUsage = 0xa4
	Hash5  TransactionAttributeUsage = 0xa5
	Hash6  TransactionAttributeUsage = 0xa6
	Hash7  TransactionAttributeUsage = 0xa7
	Hash8  TransactionAttributeUsage = 0xa8
	Hash9  TransactionAttributeUsage = 0xa9
	Hash10 TransactionAttributeUsage = 0xaa
	Hash11 TransactionAttributeUsage = 0xab
	Hash12 TransactionAttributeUsage = 0xac
	Hash13 TransactionAttributeUsage = 0xad
	Hash14 TransactionAttributeUsage = 0xae
	Hash15 TransactionAttributeUsage = 0xaf

	Remark   TransactionAttributeUsage = 0xf0
	Remark1  TransactionAttributeUsage = 0xf1
	Remark2  TransactionAttributeUsage = 0xf2
	Remark3  TransactionAttributeUsage = 0xf3
	Remark4  TransactionAttributeUsage = 0xf4
	Remark5  TransactionAttributeUsage = 0xf5
	Remark6  TransactionAttributeUsage = 0xf6
	Remark7  TransactionAttributeUsage = 0xf7
	Remark8  TransactionAttributeUsage = 0xf8
	Remark9  TransactionAttributeUsage = 0xf9
	Remark10 TransactionAttributeUsage = 0xfa
	Remark11 TransactionAttributeUsage = 0xfb
	Remark12 TransactionAttributeUsage = 0xfc
	Remark13 TransactionAttributeUsage = 0xfd
	Remark14 TransactionAttributeUsage = 0xfe
	Remark15 TransactionAttributeUsage = 0xff
)

// Transaction types
type TransactionType byte

const (
	Miner_Transaction      TransactionType = 0x00
	Issue_Transaction      TransactionType = 0x01
	Claim_Transaction      TransactionType = 0x02
	Enrollment_Transaction TransactionType = 0x20
	Register_Transaction   TransactionType = 0x40
	Contract_Transaction   TransactionType = 0x80
	State_Transaction      TransactionType = 0x90
	/// <summary>
	/// Publish scripts to the blockchain for being invoked later.
	/// </summary>
	Publish_Transaction    TransactionType = 0xd0
	Invocation_Transaction TransactionType = 0xd1
)

const (
	TransactionVersion byte = 1 // neo-2.x
)

// TransactionInput alias CoinReference
type CoinReference struct {
	PrevHash  helper.UInt256
	PrevIndex uint16
}

// TransactionOutput
type TransactionOutput struct {
	AssetId    helper.UInt256
	Value      helper.Fixed8
	ScriptHash helper.UInt160
}

// Witness
type Witness struct {
	InvocationScript   []byte // signature
	VerificationScript []byte // pub key
}

type Transaction struct {
	Type       TransactionType
	Version    byte
	Hash       helper.UInt256
	Attributes []TransactionAttribute
	Inputs     []CoinReference
	Outputs    []TransactionOutput
	Witnesses  []Witness
}

// CreateContractTransaction creates a contract transaction
// TODO add contract transaction type
func CreateContractTransaction() *Transaction {
	tx := &Transaction{
		Type:    Contract_Transaction,
		Version: TransactionVersion,
	}
	return tx
}

// UnsignedRawTransaction ...
func (tx *Transaction) UnsignedRawTransaction() []byte {
	buff := new(bytes.Buffer)
	buff.Write(tx.UnsignedRawTransactionPart1())
	buff.Write(tx.UnsignedRawTransactionPart2())
	return buff.Bytes()
}

func (tx *Transaction) UnsignedRawTransactionPart1() []byte {
	buff := new(bytes.Buffer)
	buff.WriteByte(byte(tx.Type))
	buff.WriteByte(tx.Version)
	return buff.Bytes()
}

func (tx *Transaction) UnsignedRawTransactionPart2() []byte {
	buff := new(bytes.Buffer)
	buff.Write(tx.SerializeAttributes())
	buff.Write(tx.SerializeInputs())

	return buff.Bytes()
}

func (tx *Transaction) SerializeAttributes() []byte {
	buff := new(bytes.Buffer)
	attributeCount := helper.VarIntFromInt(len(tx.Attributes))
	buff.Write(attributeCount.Bytes())
	for _, attr := range tx.Attributes {
		buff.WriteByte(byte(attr.Usage))
		if attr.Usage == DescriptionUrl {
			buff.WriteByte(byte(len(attr.Data)))
		} else if attr.Usage == Description || attr.Usage >= Remark {
			var dataLength helper.VarInt
			dataLength.Value = uint64(len(attr.Data))
			buff.Write(dataLength.Bytes())
		}
		if attr.Usage == ECDH02 || attr.Usage == ECDH03 {
			data := attr.Data[1:33]
			buff.Write(data)
		}
		buff.Write(attr.Data)
	}
	return buff.Bytes()
}

func (tx *Transaction) SerializeInputs() []byte {
	buff := new(bytes.Buffer)
	inputCount := helper.VarIntFromInt(len(tx.Inputs))
	buff.Write(inputCount.Bytes())
	for i := 0; i < len(tx.Inputs); i++ {
		buff.Write(tx.Inputs[i].PrevHash.Data)
		buff.Write(helper.VarIntFromUInt64(uint64(tx.Inputs[i].PrevIndex)).Bytes())
	}
	return buff.Bytes()
}

func (tx *Transaction) SerializeOutputs() []byte {
	buff := new(bytes.Buffer)
	outputCount := helper.VarIntFromInt(len(tx.Outputs))
	buff.Write(outputCount.Bytes())
	for i := 0; i < len(tx.Outputs); i++ {
		buff.Write(tx.Outputs[i].AssetId.Data)
		buff.Write(helper.VarIntFromUInt64(uint64(tx.Outputs[i].Value.Value)).Bytes())
		buff.Write(tx.Outputs[i].ScriptHash.Data)
	}
	return buff.Bytes()
}

// SerializeWitnesses returns serialized witness data
func (tx *Transaction) SerializeWitnesses() []byte {
	buff := new(bytes.Buffer)
	var witnessCount helper.VarInt
	witnessCount.Value = uint64(len(tx.Witnesses))
	buff.Write(witnessCount.Bytes())
	var invocationCount helper.VarInt
	var verificationCount helper.VarInt
	for i := 0; i < len(tx.Witnesses); i++ {
		invocationCount = helper.VarIntFromInt(len(tx.Witnesses[i].InvocationScript))
		buff.Write(invocationCount.Bytes())
		buff.Write(tx.Witnesses[i].InvocationScript)
		verificationCount = helper.VarIntFromInt(len(tx.Witnesses[i].VerificationScript))
		buff.Write(verificationCount.Bytes())
		buff.Write(tx.Witnesses[i].VerificationScript)
	}
	return buff.Bytes()
}

// RawTransaction 返回签名后的完整二进制交易
func (tx *Transaction) RawTransaction() []byte {
	return append(tx.UnsignedRawTransaction(), tx.SerializeWitnesses()...)
}

// RawTransactionString 返回完整交易的二进制字符串
func (tx *Transaction) RawTransactionString() string {
	return hex.EncodeToString(tx.RawTransaction())
}

func (tx *Transaction) TXID() string {
	txid := crypto.Sha256(tx.UnsignedRawTransaction())
	return hex.EncodeToString(helper.ReverseBytes(txid))
}
