package trace

import (
	"go_crud/models"
	"io/ioutil"
	"log"
	"path"
)

func Files(config models.Config, selectedDir string) []string {
	files, err := ioutil.ReadDir(path.Join(config.DataDir, selectedDir))
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	for _, f := range files {
		if !f.IsDir() {
			arr = append(arr, f.Name())
		}
	}

	return arr
}
