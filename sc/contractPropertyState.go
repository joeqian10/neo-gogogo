package sc

type ContractPropertyState byte

const (
	NoProperty       ContractPropertyState = 0
	HasStorage       ContractPropertyState = 1
	HasDynamicInvoke ContractPropertyState = 2
	Payable          ContractPropertyState = 4
)
