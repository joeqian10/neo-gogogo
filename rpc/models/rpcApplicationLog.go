package models

type RpcApplicationLog struct {
	TxId       string         `json:"txid"`
	Executions []RpcExecution `json:"executions"`
}

type RpcExecution struct {
	Trigger       string            `json:"trigger"`
	Contract      string            `json:"contract"`
	VMState       string            `json:"vmstate"`
	GasConsumed   string            `json:"gas_consumed"`
	Stack         []InvokeStack     `json:"stack"`
	Notifications []RpcNotification `json:"notifications"`
}

type RpcNotification struct {
	Contract string      `json:"contract"`
	State    InvokeStack `json:"state"`
}
