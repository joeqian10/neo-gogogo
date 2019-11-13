package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// Witness
type Witness struct {
	InvocationScript   []byte         // signature
	VerificationScript []byte         // pub key
	scriptHash         helper.UInt160 // script hash
}

// Deserialize implements Serializable interface.
func (w *Witness) Deserialize(br *io.BinReader) {
	w.InvocationScript = br.ReadBytes()
	w.VerificationScript = br.ReadBytes()
}

// Serialize implements Serializable interface.
func (w *Witness) Serialize(bw *io.BinWriter) {
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

// ScriptHash returns the hash of the VerificationScript.
//func (w Witness) ScriptHash() helper.UInt160 {
//	value, _:= helper.UInt160FromBytes(crypto.Hash160(w.VerificationScript));
//	return value
//}

func CreateWitness(invocationScript []byte, verificationScript []byte) (witness *Witness, err error) {
	if len(verificationScript) == 0 {
		return nil, fmt.Errorf("verificationScript should not be empty")
	}
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: verificationScript}
	witness.scriptHash, err = helper.UInt160FromBytes(crypto.Hash160(witness.VerificationScript))
	return
}

func CreateWitnessWithScriptHash(scriptHash helper.UInt160, invocationScript []byte) (witness *Witness) {
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: []byte{}, scriptHash: scriptHash}
	return
}
