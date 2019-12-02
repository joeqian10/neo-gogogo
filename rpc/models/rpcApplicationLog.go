package models

type RpcApplicationLog struct {
	TxId       string `json:"txid"`
	Executions []struct {
		Trigger     string `json:"trigger"`
		Contract    string `json:"contract"`
		VMState     string `json:"vmstate"`
		GasConsumed string `json:"gas_consumed"`
		Stack       []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"stack"`
		Notifications []struct {
			Contract string `json:"contract"`
			State    struct {
				Type  string        `json:"type"`
				Value []interface{} `json:"value"`
				//Value []struct {
				//	Type  string `json:"type"`
				//	Value string `json:"value"`
				//}
			} `json:"state"`
		} `json:"notifications"`
	} `json:"executions"`
}
