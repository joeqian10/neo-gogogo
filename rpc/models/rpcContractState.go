package models

type ContractState struct {
	Version     int      `json:"version"`
	Hash        string   `json:"hash"`
	Script      string   `json:"script"`
	Parameters  []string `json:"parameters"`
	ReturnType  string   `json:"returntype"`
	Name        string   `json:"name"`
	CodeVersion string   `json:"code_version"`
	Author      string   `json:"author"`
	Email       string   `json:"email"`
	Description string   `json:"description"`
	Properties  struct {
		Storage       bool `json:"storage"`
		DynamicInvoke bool `json:"dynamic_invoke"`
	} `json:"properties"`
}
