package mpt

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

//StateRoot truct of StateRoot message
type StateRoot struct {
	Version   uint32 `json:"version"`
	Index     uint32 `json:"index"`
	PreHash   string `json:"prehash"`
	StateRoot string `json:"stateroot"`
	Witness   []struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	} `json:"witness"`
}

func (sr *StateRoot) Deserialize(br *io.BinaryReader) {
	sr.DeserializeUnsigned(br)
	l := br.ReadVarUint()
	if l != 1 {
		return
	}
	sr.Witness = make([]struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	}, 1)
	sr.Witness[0].InvocationScript = helper.BytesToHex(br.ReadVarBytes())
	sr.Witness[0].VerificationScript = helper.BytesToHex(br.ReadVarBytes())
}

func (sr *StateRoot) Serialize(bw *io.BinaryWriter) {
	sr.SerializeUnsigned(bw)
	bw.WriteVarUint(1)
	bw.WriteVarBytes(helper.HexToBytes(sr.Witness[0].InvocationScript))
	bw.WriteVarBytes(helper.HexToBytes(sr.Witness[0].VerificationScript))
}

func (sr *StateRoot) DeserializeUnsigned(br *io.BinaryReader) {
	br.ReadLE(&sr.Version)
	br.ReadLE(&sr.Index)
	var preHash, stateRoot helper.UInt256
	br.ReadLE(&preHash)
	br.ReadLE(&stateRoot)
	sr.PreHash = preHash.String()
	sr.StateRoot = stateRoot.String()
}

func (sr *StateRoot) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(sr.Version)
	bw.WriteLE(sr.Index)
	preHash, _ := helper.UInt256FromString(sr.PreHash)
	bw.WriteLE(preHash)
	stateRoot, _ := helper.UInt256FromString(sr.StateRoot)
	bw.WriteLE(stateRoot)
}
