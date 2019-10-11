package trace

import (
	"gocrud/models"
	"io/ioutil"
	"log"
)

// Directory check the folder and return an array according to the list of all directories
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
