package models

type Basic struct {
	Tokens  []string    `json:"tokens"`
	Payload interface{} `json:"payload"`
}
