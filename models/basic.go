package models

// Header is part of Basic
type Header struct {
	ContentType  string `json:"Content-Type"`
	UserAgent    string `json:"User-Agent"`
	Accept       string `json:"Accept"`
	CacheControl string `json:"Cache-Control"`
	Host         string `json:"Host"`
	Connection   string `json:"Connection"`
}

// Basic is used for holding environments which is used for translate variables inside the payloads and etc
type Basic struct {
	Env    []interface{} `json:"env"`
	Header Header        `json:"header"`
}
