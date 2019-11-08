package models

type RpcVersion struct {
	Port      uint32 `json:"port"`
	Nonce     uint32 `json:"nonce"`
	Useragent string `json:"useragent"`
}
