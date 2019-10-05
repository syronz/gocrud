package engine

import (
	"go_crud/models"
)

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
