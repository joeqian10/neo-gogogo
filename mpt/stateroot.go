package mpt

import "github.com/joeqian10/neo-gogogo/helper/io"

//StateRoot truct of StateRoot message
type StateRoot struct {
	Version   uint32 `json:"version"`
	Index     uint32 `json:"index"`
	PreHash   string `json:"prehash"`
	StateRoot string `json:"stateroot"`
	Witness   struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	} `json:"witness"`
}

func (sr *StateRoot) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&sr.Version)
	br.ReadLE(&sr.Index)
	sr.PreHash = br.ReadVarString()
	sr.StateRoot = br.ReadVarString()
	sr.Witness.InvocationScript = br.ReadVarString()
	sr.Witness.VerificationScript = br.ReadVarString()
}

func (sr *StateRoot) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(sr.Version)
	bw.WriteLE(sr.Index)
	bw.WriteVarString(sr.PreHash)
	bw.WriteVarString(sr.StateRoot)
	bw.WriteVarString(sr.Witness.InvocationScript)
	bw.WriteVarString(sr.Witness.VerificationScript)
}
