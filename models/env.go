package models

// Dics is part of Env
type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Env used for encoding env (inside basic.json) and used for translate payloads json files
type Env struct {
	Name string `json:"name"`
	Dict []Dict `json:"dict"`
}
