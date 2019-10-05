package engine

import (
	"go_crud/helper"
)

func (e *Engine) TranslatedContent() string {
	contentInJSON := helper.EnvReplace(e.Content, e.GetEnv())

	return contentInJSON
}
