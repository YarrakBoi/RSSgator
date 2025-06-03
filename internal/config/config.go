package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

type State struct {
	ptr *Config
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	
	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err 
	}

	configFilePath := filepath.Join(home, configFileName)
	
	data, err := os.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func SetConfig(current_name string) error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err 
	}

	configFilePath := filepath.Join(home, configFileName)

	currConfig, err := ReadConfig() 

	if err != nil {
		return err 
	}
	currConfig.Current_user_name = current_name

	configJson, err := json.Marshal(currConfig)

	if err != nil {
		return err 
	}

	err = os.WriteFile(configFilePath, configJson, 0666)
	if err != nil {
		return err
	}

	return nil
}

