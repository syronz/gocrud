package trace

import (
	"gocrud/models"
	"io/ioutil"
	"log"
	"path"
)

// Files get a directory and returns and array consist of all files
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
