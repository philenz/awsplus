package config

import (
	"os"
	"log"
	"encoding/json"
)

type Configuration struct {
	Port   int
	Static string
	Region string
	DBName string
	DBUser string
	DBPass string
}

var Config Configuration

func init() {
	file, err := os.Open("runtime/etc/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}