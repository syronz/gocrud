package trace

import (
	"go_crud/models"
	"io/ioutil"
	"log"
)

func Directory(config models.Config) []string {
	files, err := ioutil.ReadDir(config.DataDir)
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	for _, f := range files {
		if f.IsDir() {
			arr = append(arr, f.Name())
		}
	}

	return arr
}
