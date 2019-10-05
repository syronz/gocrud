package loaders

import (
	"encoding/json"
	"go_crud/models"
	"log"
	"os"
	//"os/user"
)

// load config.json
func LoadConfig(configPath string) models.Config {
	//	usr, err := user.Current()
	//if err != nil {
	//log.Fatal(err)
	//}

	var config models.Config
	//configFile, err := os.Open(usr.HomeDir + "/.go_crud/config.json")
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
