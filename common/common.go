package common

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var cnfFile = flag.String("config", "config.json", "config file")

type Config struct {
	DbUser     string `json:"db_user"`
	DbPass     string `json:"db_pass"`
	DbName     string `json:"db_name"`
	DbUrl      string `json:"db_url"`
	ServerPort string `json:"server_port"`
}

func ReadConfig() (config *Config, err error) {
	b, err := os.ReadFile(*cnfFile)
	if err != nil {
		log.Printf("error :%v", err)
		return
	}
	config = &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Printf("error :%v", err)
		return
	}
	return
}
