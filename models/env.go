package models

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Env struct {
	Name string `json:"name"`
	Dict []Dict `json:"dict"`
}
