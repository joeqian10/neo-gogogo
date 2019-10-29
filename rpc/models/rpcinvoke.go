package models

type InvokeResult struct {
	Script      string              `json:"script"`
	State       string              `json:"state"`
	GasConsumed string              `json:"gas_consumed"`
	Stack       []InvokeStackResult `json:"stack"`
}

type InvokeStackResult struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
