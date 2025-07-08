package config

import (
	"encoding/json"
	"fmt"
	"gator/internal/utils"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error locating config file: %w", err)
	}

	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer utils.SafeClose(jsonFile)

	decoder := json.NewDecoder(jsonFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding JSON from config file %s: %w", configFilePath, err)
	}

	return cfg, nil

}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	return write(*c)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find home directory: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error locating config file: %w", err)
	}

	file, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening config file for writing: %w", err)
	}
	defer utils.SafeClose(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return fmt.Errorf("error encoding config to json: %w", err)
	}

	return nil
}
