package transaction

import (
	"encoding/hex"
	"encoding/json"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"../helper/io"
)

// Witness contains 2 scripts.
type Witness struct {
	InvocationScript   []byte
	VerificationScript []byte
}

// DecodeBinary implements Serializable interface.
func (w *Witness) DecodeBinary(br *io.BinReader) {
	w.InvocationScript = br.ReadBytes()
	w.VerificationScript = br.ReadBytes()
}

// EncodeBinary implements Serializable interface.
func (w *Witness) EncodeBinary(bw *io.BinWriter) {
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
func (w Witness) ScriptHash() helper.UInt160 {
	value,_:= helper.NewUInt160(crypto.Hash160(w.VerificationScript));
	return value
}
