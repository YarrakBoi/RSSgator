package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

type State struct {
	Cfg *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CmdNames map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.CmdNames[cmd.Name]
	if !ok {
		return fmt.Errorf("command does not exist, please register if needed")
	}
	
	fErr := f(s, cmd)

	if fErr != nil {
		return fErr
	}

	return nil


}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CmdNames[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username not supplied")
	}
	s.Cfg.Current_user_name = cmd.Args[0]
	err := SaveConfig(s.Cfg)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
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

func SaveConfig(cfg *Config) error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err 
	}

	configFilePath := filepath.Join(home, configFileName)

	configJson, err := json.Marshal(cfg)

	if err != nil {
		return err 
	}

	err = os.WriteFile(configFilePath, configJson, 0666)
	if err != nil {
		return err
	}

	return nil
}

