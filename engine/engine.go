package engine

import (
	"go_crud/models"
)

// Engine is must important struct for saving almost anything and pass important variables to various functions
// also holds key methods
type Engine struct {
	Page    int
	PrePage int
	Dirs    []string
	Files   []string
	Env     map[string]models.Env
	EnvArr  []string

	SelectedDir  string
	SelectedFile string
	SelectedEnv  string
	Content      string
	Config       models.Config
	Header       models.Header
}
