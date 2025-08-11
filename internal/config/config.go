// Package config is reponsible for reading and writing to the .gatorconfig.json file
package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	dbURL           url.URL
	currentUserName string
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}
	fmt.Println(configFilePath)
	return Config{}, nil
}

func (*Config) SetUser() {
	fmt.Println("Function config.SetUser() is not yet implemented")
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	fmt.Println("Function write is not yet implemented")
	return nil
}
