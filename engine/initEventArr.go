package engine

// InitEnvArr is used for generate first time list of folders, DEPRECATED!
func (e *Engine) InitEnvArr(startPage int) {
	for i := range e.Env {
		e.EnvArr = append(e.EnvArr, i)
	}
	e.SelectedEnv = e.EnvArr[0]

	e.Page = startPage
	e.PrePage = startPage
}
