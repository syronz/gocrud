package models

type Basic struct {
	Tokens  []string      `json:"tokens"`
	Env     []interface{} `json:"env"`
	Payload interface{}   `json:"payload"`
}
