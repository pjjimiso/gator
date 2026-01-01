package config


import (
	"os"
	"encoding/json"
)


const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil { 
		return "", err
	}
	
	cfgPath := homeDir + "/" + configFileName
	return cfgPath, nil
}

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil { 
		return Config{}, err
	}

	data, err := os.ReadFile(cfgPath)
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

/*
func (c Config) SetUser() error {
	
}
*/	
