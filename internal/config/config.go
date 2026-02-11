package config


import (
	"os"
	"encoding/json"
)


const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL		string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
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

func write(cfg Config) error {
	cfgPath, err := getConfigFilePath()
	if err != nil { 
		return err
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(cfgPath, b, 0600)
	return nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

