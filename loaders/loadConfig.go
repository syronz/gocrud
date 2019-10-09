package loaders

import (
	"encoding/json"
	"go_crud/models"
	"log"
	"os"
)

// LoadConfig will encode the config.json and save it to the config struct
func LoadConfig(configPath string) models.Config {
	var config models.Config

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal("can't open config.json: ", err)
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	return config
}
