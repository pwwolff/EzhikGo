package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ready        bool
	DataBaseAddr string `json:"mongo_address"`
	DataBaseName string `json:"mongo_db_name"`
	EzhikAddress string `json:"ezhik_address"`
	Username     string `json:"mongo_username"`
	Password     string `json:"mongo_password"`
	Port         string `json:"port"`
}

var config Configuration

func GetConfig() Configuration {
	if config.ready {
		return config
	}

	file, err := ioutil.ReadFile("config/configuration.json")
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}
