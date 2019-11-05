package models

type RpcAssetState struct {
	Version    uint32      `json:"version"`
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Name       []AssetName `json:"name"`
	Amount     string      `json:"amount"`
	Available  string      `json:"available"`
	Precision  uint32      `json:"precision"`
	Owner      string      `json:"owner"`
	Admin      string      `json:"admin"`
	Issuer     string      `json:"issuer"`
	Expiration uint32      `json:"expiration"`
	Frozen     bool        `json:"frozen"`
}

type AssetName struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
}

type AssetBalance struct {
	Balance   string `json:"balance"`
	Confirmed string `json:"confirmed"`
}
