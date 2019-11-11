package models

type RpcVersion struct {
	Port      int `json:"port"`
	Nonce     int `json:"nonce"`
	Useragent string `json:"useragent"`
}
