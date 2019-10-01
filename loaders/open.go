package loaders

import (
	//"encoding/json"
	"fmt"
	"go_crud/models"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// load config.json
func Open(config models.Config, selectedDir, selectedFile string) string {
	fmt.Println("FILE: ", path.Join(config.DataDir, selectedDir, selectedFile))
	f, err := os.Open(path.Join(config.DataDir, selectedDir, selectedFile))
	if err != nil {
		log.Fatal("can't open json file: ", err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	return string(b)
}
