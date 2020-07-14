package models

type RpcApplicationLog struct {
	TxId       string         `json:"txid"`
	Executions []RpcExecution `json:"executions"`
}

type RpcExecution struct {
	Trigger       string                 `json:"trigger"`
	Contract      string                 `json:"contract"`
	VMState       string                 `json:"vmstate"`
	GasConsumed   string                 `json:"gas_consumed"`
	Stack         []RpcContractParameter `json:"stack"`
	Notifications []RpcNotification      `json:"notifications"`
}

type RpcContractParameter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RpcState struct {
	Type  string                 `json:"type"`
	Value []RpcContractParameter `json:"value"`
}

type RpcNotification struct {
	Contract string   `json:"contract"`
	State    RpcState `json:"state"`
}
