package models

type RpcApplicationLog struct {
	TxId       string `json:"txid"`
	Executions []struct {
		Trigger       string                 `json:"trigger"`
		Contract      string                 `json:"contract"`
		VMState       string                 `json:"vmstate"`
		GasConsumed   string                 `json:"gas_consumed"`
		Stack         []RpcContractParameter `json:"stack"`
		Notifications []struct {
			Contract string `json:"contract"`
			State    struct {
				Type  string                 `json:"type"`
				Value []RpcContractParameter `json:"value"`
			} `json:"state"`
		} `json:"notifications"`
	} `json:"executions"`
}

type RpcContractParameter struct {
	Type string 		`json:"type"`
	Value interface{} 	`json:"value"`
}