package mpt

import (
	"bytes"
	"errors"
	"io"

	nio "github.com/joeqian10/neo-gogogo/helper/io"

	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
)

//ProofDb a db to use for verify
type ProofDb struct {
	nodes map[string]([]byte)
}

//NewProofDb new instance of ProofDb from a string list
func NewProofDb(proofBytes []byte) *ProofDb {
	proof := bytesToArray(proofBytes)
	p := &ProofDb{}
	p.nodes = make(map[string]([]byte), len(proof))
	for _, v := range proof {
		data := v
		hashstr := helper.BytesToHex(crypto.Hash256(data))
		p.nodes[hashstr] = data
	}
	return p
}

//Get for TrieDb
func (pd *ProofDb) Get(key []byte) ([]byte, error) {
	keystr := helper.BytesToHex(key)
	if v, ok := pd.nodes[keystr]; ok {
		return v, nil
	}
	return nil, errors.New("cant find the value in ProofDb, key=" + keystr)
}

func bytesToArray(data []byte) [][]byte {
	buffer := bytes.NewBuffer(data)
	reader := nio.NewBinaryReaderFromIO(io.Reader(buffer))
	count := reader.ReadVarUint()
	result := make([][]byte, count)
	for i := uint64(0); i < count; i++ {
		result[i] = reader.ReadVarBytes()
	}
	return result
}
