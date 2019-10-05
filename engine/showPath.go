package engine

import (
	"fmt"
)

func (e *Engine) ShowCurrentPath() {
	fmt.Printf("Path: %v/%v\n", e.SelectedDir, e.SelectedFile)
	translatedContent := e.TranslatedContent()
	fmt.Printf("Content:\n %v\n", translatedContent)

}
