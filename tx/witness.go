package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
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
	w.InvocationScript = br.ReadBytes()
	w.VerificationScript = br.ReadBytes()
}

// Serialize implements Serializable interface.
func (w *Witness) Serialize(bw *io.BinaryWriter) {
	bw.WriteBytes(w.InvocationScript)
	bw.WriteBytes(w.VerificationScript)
}

// MarshalJSON implements the json marshaller interface.
func (w *Witness) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"invocation":   hex.EncodeToString(w.InvocationScript),
		"verification": hex.EncodeToString(w.VerificationScript),
	}

	return json.Marshal(data)
}

// Create Witness with invocationScript and verificationScript
func CreateWitness(invocationScript []byte, verificationScript []byte) (witness *Witness, err error) {
	if len(verificationScript) == 0 {
		return nil, fmt.Errorf("verificationScript should not be empty")
	}
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: verificationScript}
	witness.scriptHash, err = helper.UInt160FromBytes(crypto.Hash160(witness.VerificationScript))
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
	invocationScript := builder.ToArray()

	// verificationScript: SignatureRedeemScript
	verificationScript := keys.CreateSignatureRedeemScript(pair.PublicKey)
	return CreateWitness(invocationScript, verificationScript)
}

// create multi-signature witness
func CreateMultiSignatureWitness(msg []byte, pairs []*keys.KeyPair, least int, publicKeys []*keys.PublicKey) (witness *Witness, err error) {
	// TODO ensure the pairs match with publicKeys
	if len(pairs) == least {
		return witness, fmt.Errorf("the multi-signature contract needs least %v signatures", least)
	}
	// invocationScript: push signature
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].PublicKey.Compare(pairs[j].PublicKey) == 1
	})
	builder := sc.NewScriptBuilder()
	for _, pair := range pairs {
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
