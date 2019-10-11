package engine

import (
	"gocrud/models"
)

// GetEnv return back the selected environment
func (e *Engine) GetEnv() models.Env {
	return e.Env[e.SelectedEnv]
}
