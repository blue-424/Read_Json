package model

type Payload struct {
	Function   string                 `json:"function"`
	Parameters map[string]interface{} `json:"parameters"`
}
