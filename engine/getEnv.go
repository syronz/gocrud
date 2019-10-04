package engine

import (
	"go_crud/models"
)

func (e *Engine) GetEnv() models.Env {
	return e.Env[e.SelectedEnv]
}
