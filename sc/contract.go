package sc

import (
	"github.com/joeqian10/neo-gogogo/helper"
)

type Contract struct {
	Script        []byte
	ParameterList []ContractParameterType
	ScriptHash    helper.UInt160
	Address       string
}
