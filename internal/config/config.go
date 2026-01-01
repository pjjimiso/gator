package config


import (
	"os"
	"encoding/json"
	"fmt"
)


const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
}


func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil { 
		return Config{}, err
	}

	data, err := os.ReadFile(homeDir + "/" + configFileName)
	if err != nil { 
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

