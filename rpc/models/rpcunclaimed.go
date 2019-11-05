package models

type UnclaimedGasInWallet struct {
	Available   string `json:"available"`
	Unavailable string `json:"unavailable"`
}

type UnclaimedGasInAddress struct {
	Available   float32 `json:"available"`
	Unavailable float32 `json:"unavailable"`
	Unclaimed   float32 `json:"unclaimed"`
}
