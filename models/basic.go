package models

type Header struct {
	ContentType  string `json:"Content-Type"`
	UserAgent    string `json:"User-Agent"`
	Accept       string `json:"Accept"`
	CacheControl string `json:"Cache-Control"`
	Host         string `json:"Host"`
	Connection   string `json:"Connection"`
}

type Basic struct {
	Env    []interface{} `json:"env"`
	Header Header        `json:"header"`
}
