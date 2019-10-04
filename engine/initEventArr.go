package engine

func (e *Engine) InitEnvArr(startPage int) {
	for i, _ := range e.Env {
		e.EnvArr = append(e.EnvArr, i)
	}
	e.SelectedEnv = e.EnvArr[0]

	e.Page = startPage
	e.PrePage = startPage
}
