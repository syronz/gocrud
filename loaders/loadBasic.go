package loaders

import (
	"encoding/json"
	"go_crud/models"
	"log"
	"os"
)

// load basic.json
func LoadBasic(config models.Config) models.Basic {
	var basic models.Basic
	basicFile, err := os.Open(config.BasePath)
	if err != nil {
		log.Fatal("can't open basic.json: ", err)
	}
	defer basicFile.Close()
	decoder := json.NewDecoder(basicFile)
	err = decoder.Decode(&basic)
	if err != nil {
		log.Fatal("can't decode basic JSON: ", err)
	}

	return basic
}
