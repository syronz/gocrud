package engine

import (
	"go_crud/models"
)

// GetEnv return back the selected environment
func (e *Engine) GetEnv() models.Env {
	return e.Env[e.SelectedEnv]
}
