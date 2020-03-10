package mpt

type StateRoot struct {
	Version   int    `json:"version"`
	Index     uint32 `json:"index"`
	PreHash   string `json:"prehash"`
	StateRoot string `json:"stateroot"`
	Witness   struct {
		InvocationScript   string `json:"invocation"`
		VerificationScript string `json:"verification"`
	} `json:"witness"`
}
