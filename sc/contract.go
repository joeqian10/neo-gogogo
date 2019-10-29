package sc

import "github.com/joeqian10/neo-gogogo/crypto"

type Contract struct {
	Script        []byte
	ParameterList []ContractParameterType
	ScriptHash crypto.HASH160
	Address
}


