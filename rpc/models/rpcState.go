package models

import (
	"github.com/joeqian10/neo-gogogo/mpt"
)

type MPTProof struct {
	Success bool   `json:"success"`
	Proof   string `json:"proof"`
}

type StateHeight struct {
	BlockHeight uint32 `json:"blockheight"`
	StateHeight uint32 `json:"stateheight"`
}

type StateRootState struct {
	Flag      string        `json:"flag"`
	StateRoot mpt.StateRoot `json:"stateroot"`
}
