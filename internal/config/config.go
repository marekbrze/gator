// Package config is reponsible for reading and writing to the .gatorconfig.json file
package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, nil
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, nil
	}
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return Config{}, nil
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	if err := write(*c); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = os.WriteFile(configFilePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
