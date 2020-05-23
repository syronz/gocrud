package models

// Config is a main struct for saving target of important file, info from config.json will encoded here
type Config struct {
	DataDir  string `json:"data_dir"`
	BasePath string `json:"base_path"`
	Timeout  int    `json:"timeout"`
	Delay    int    `json:"delay"`
}
