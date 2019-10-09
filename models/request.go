package models

// Request is used for encoding payload files
type Request struct {
	Method        string      `json:"method"`
	Url           string      `json:"url"`
	Authorization string      `json:"authorization"`
	Payload       interface{} `json:"payload"`
}
