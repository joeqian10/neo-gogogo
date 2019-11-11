package models

type AccountState struct {
	Version    int                `json:"version"`
	ScriptHash string                `json:"script_hash"`
	Frozen     bool                  `json:"frozen"`
	Votes      []interface{}         `json:"votes"`
	Balances   []AccountStateBalance `json:"balances"`
}

type AccountStateBalance struct {
	Asset string `json:"asset"`
	Value string `json:"value"`
}
