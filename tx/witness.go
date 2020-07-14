package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
	"sort"
)

// Witness
type Witness struct {
	InvocationScript   []byte         // signature
	VerificationScript []byte         // pub key
	scriptHash         helper.UInt160 // script hash
}

// Deserialize implements Serializable interface.
func (w *Witness) Deserialize(br *io.BinaryReader) {
	w.InvocationScript = br.ReadVarBytes()
	w.VerificationScript = br.ReadVarBytes()
}

// Serialize implements Serializable interface.
func (w *Witness) Serialize(bw *io.BinaryWriter) {
	bw.WriteVarBytes(w.InvocationScript)
	bw.WriteVarBytes(w.VerificationScript)
}

// MarshalJSON implements the json marshaller interface.
func (w *Witness) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"invocation":   hex.EncodeToString(w.InvocationScript),
		"verification": hex.EncodeToString(w.VerificationScript),
	}

	return json.Marshal(data)
}

// this method is a getter of scriptHash
func (w *Witness) GetScriptHash() helper.UInt160 {
	w.scriptHash, _ = helper.BytesToScriptHash(w.VerificationScript)
	return w.scriptHash
}

// Create Witness with invocationScript and verificationScript
func CreateWitness(invocationScript []byte, verificationScript []byte) (witness *Witness, err error) {
	if len(verificationScript) == 0 {
		return nil, fmt.Errorf("verificationScript should not be empty")
	}
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: verificationScript}
	witness.scriptHash, err = helper.BytesToScriptHash(witness.VerificationScript)
	return
}

// this is only used for empty VerificationScript, neo block chain will search the contract script with scriptHash
func CreateWitnessWithScriptHash(scriptHash helper.UInt160, invocationScript []byte) (witness *Witness) {
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: []byte{}, scriptHash: scriptHash}
	return
}

// create single signature witness
func CreateSignatureWitness(msg []byte, pair *keys.KeyPair) (witness *Witness, err error) {
	// 	invocationScript: push signature
	signature, err := pair.Sign(msg)
	if err != nil {
		return
	}
	builder := sc.NewScriptBuilder()
	_ = builder.EmitPushBytes(signature)
	invocationScript := builder.ToArray() // length 65

	// verificationScript: SignatureRedeemScript
	verificationScript := keys.CreateSignatureRedeemScript(pair.PublicKey)
	return CreateWitness(invocationScript, verificationScript)
}

// create multi-signature witness
func CreateMultiSignatureWitness(msg []byte, pairs []*keys.KeyPair, least int, publicKeys []*keys.PublicKey) (witness *Witness, err error) {
	if len(pairs) < least {
		return witness, fmt.Errorf("the multi-signature contract needs least %v signatures", least)
	}
	// invocationScript: push signature
	keyPairs := keys.KeyPairSlice(pairs)
	sort.Sort(keyPairs) // ascending

	builder := sc.NewScriptBuilder()
	for _, pair := range keyPairs {
		signature, err := pair.Sign(msg)
		if err != nil {
			return witness, err
		}
		err = builder.EmitPushBytes(signature)
		if err != nil {
			return witness, err
		}
	}
	invocationScript := builder.ToArray()

	// verificationScript: CreateMultiSigRedeemScript
	verificationScript, _ := keys.CreateMultiSigRedeemScript(least, publicKeys...)
	return CreateWitness(invocationScript, verificationScript)
}


func VerifySignatureWitness(msg []byte, witness *Witness) bool {
	invocationScript := witness.InvocationScript
	length := invocationScript[0]
	if int(length) != len(invocationScript[1:]) {
		return false
	}
	signature := invocationScript[1:]
	verificationScript := witness.VerificationScript
	if len(verificationScript) != 35 {
		return false
	}
	data := verificationScript[:34] // length 34
	publicKey, _ := keys.NewPublicKey(data[1:])
	return keys.VerifySignature(msg, signature, publicKey)
}


func VerifyMultiSignatureWitness(msg []byte, witness *Witness) bool {
	invocationScript := witness.InvocationScript
	lenInvoScript := len(invocationScript)
	if lenInvoScript%65 != 0 {
		return false
	}
	m := lenInvoScript / 65 // m signatures

	verificationScript := witness.VerificationScript
	least := verificationScript[0] - byte(sc.PUSH1) + 1 // least required signatures, usually 4

	if m < int(least) {
		return false
	} // not enough signatures
	var signatures = make([][]byte, m)
	for i := 0; i < m; i++ {
		signatures[i] = invocationScript[i*65+1 : i*65+65] // signature length is 64
	}

	lenVeriScript := len(verificationScript)
	n := verificationScript[lenVeriScript-2] - byte(sc.PUSH1) + 1 // public keys, usually 7
	if m > int(n) {
		return false
	} // too many signatures

	var pubKeys = make([]*keys.PublicKey, n)
	for i := 0; i < int(n); i++ {
		data := verificationScript[i*34+1 : i*34+35] // length 34
		publicKey, _ := keys.NewPublicKey(data[1:])
		pubKeys[i] = publicKey
	}
	return keys.VerifyMultiSig(msg, signatures, pubKeys)
}

type WitnessSlice []*Witness

func (ws WitnessSlice) Len() int           { return len(ws) }
func (ws WitnessSlice) Less(i, j int) bool { return ws[i].scriptHash.Less(ws[j].scriptHash) }
func (ws WitnessSlice) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }
