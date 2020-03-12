package models

type RpcBlockHeader struct {
	Hash              string `json:"hash"`
	Size              int    `json:"size"`
	Version           int    `json:"version"`
	PreviousBlockHash string `json:"previousblockhash"`
	MerkleRoot        string `json:"merkleroot"`
	Time              int    `json:"time"`
	Index             int    `json:"index"`
	Nonce             string `json:"nonce"`         //ulong = uint64
	NextConsensus     string `json:"nextconsensus"` //address
	Witness           struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	} `json:"witness"`
	Confirmations int    `json:"confirmations"`
	NextBlockHash string `json:"nextblockhash"`
	// cross chain support
	CrossStatesRoot string `json:"crossstatesroot"`
}

type RpcBlock struct {
	RpcBlockHeader
	Tx []RpcTransaction `json:"tx"`
}

//func (bh *RpcBlockHeader) Deserialize(br *io.BinaryReader) {
//	bh.DeserializeUnsigned(br)
//	bh.DeserializeWitness(br)
//}
//
////DeserializeUnsigned deserialize blockheader without witness
//func (bh *RpcBlockHeader) DeserializeUnsigned(br *io.BinaryReader) {
//	var h, ph, mr, cr helper.UInt256
//	br.ReadLE(&h)
//	bh.Hash = h.String()
//
//	br.ReadLE(&bh.Version)
//
//	br.ReadLE(&ph)
//	bh.PreviousBlockHash = ph.String()
//
//	br.ReadLE(&mr)
//	bh.MerkleRoot = mr.String()
//	br.ReadLE(&bh.Time)
//	br.ReadLE(&bh.Index)
//	bh.Nonce = helper.BytesToHex(helper.ReverseBytes(br.ReadUnit64Bytes())) // hex string
//
//	var nextConsensus helper.UInt160
//	br.ReadLE(&nextConsensus)
//	bh.NextBlockHash = helper.ScriptHashToAddress(nextConsensus)
//
//	bh.ChainID = helper.BytesToHex(helper.ReverseBytes(br.ReadUnit64Bytes()))
//	br.ReadLE(&cr)
//	bh.CrossStatesRoot = cr.String()
//}
//
////DeserializeWitness deserialize witness
//func (bh *RpcBlockHeader) DeserializeWitness(br *io.BinaryReader) {
//	var padding uint8
//	br.ReadLE(&padding)
//	if padding != 1 {
//		br.Err = fmt.Errorf("format error: padding must equal 1 got %d", padding)
//		return
//	}
//	bh.Script.InvocationScript = helper.BytesToHex(br.ReadVarBytes())
//	bh.Script.VerificationScript = helper.BytesToHex(br.ReadVarBytes())
//}
//
//
//
//func (bh *RpcBlockHeader) Serialize(bw *io.BufBinaryWriter) {
//	bh.SerializeUnsigned(bw)
//	bh.SerializeWitness(bw)
//}
//
////SerializeUnsigned serialize blockheader without witness
//func (bh *RpcBlockHeader) SerializeUnsigned(bw *io.BufBinaryWriter) {
//	var h, ph, mr, cr helper.UInt256
//	h, _ = helper.UInt256FromString(bh.Hash)
//	bw.WriteLE(h)
//	bw.WriteLE(bh.Version)
//	ph, _ = helper.UInt256FromString(bh.PreviousBlockHash)
//	bw.WriteLE(ph)
//	mr, _ = helper.UInt256FromString(bh.MerkleRoot)
//	bw.WriteLE(mr)
//	bw.WriteLE(bh.Time)
//	bw.WriteLE(bh.Index)
//	bw.WriteLE(helper.ReverseBytes(helper.HexToBytes(bh.Nonce)))
//	var nc helper.UInt160
//	nc, _ = helper.AddressToScriptHash(bh.NextConsensus)
//	bw.WriteLE(nc)
//	bw.WriteLE(helper.ReverseBytes(helper.HexToBytes(bh.ChainID)))
//	cr,_ = helper.UInt256FromString(bh.CrossStatesRoot)
//	bw.WriteLE(cr)
//}
//
////SerializeWitness serialize witness
//func (bh *RpcBlockHeader) SerializeWitness(bw *io.BufBinaryWriter) {
//	bw.WriteLE(uint8(1))
//	bw.WriteVarBytes(helper.HexToBytes(bh.Script.InvocationScript))
//	bw.WriteVarBytes(helper.HexToBytes(bh.Script.VerificationScript))
//}
