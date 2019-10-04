package models

type Request struct {
	Method        string      `json:"method"`
	Url           string      `json:"url"`
	Authorization string      `json:"authorization"`
	Payload       interface{} `json:"payload"`
}
