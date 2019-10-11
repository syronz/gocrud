package engine

import (
	"gocrud/helper"
)

// TranslatedContent prepare request by swap keywords with variables, which defined in basic.json and env part
func (e *Engine) TranslatedContent() string {
	contentInJSON := helper.EnvReplace(e.Content, e.GetEnv())

	return contentInJSON
}
