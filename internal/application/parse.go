package application

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ReplicationConfig struct {
	Active   string   `yaml:"active"`
	SyncedTo []string `yaml:"synced_to"`
}

type Credential struct {
	Name     string `yaml:"name"`
	AppRole  string `yaml:"appRole"`
	SecretID string `yaml:"secretID"`
}

type Config struct {
	Replication []ReplicationConfig `yaml:"replication"`
	Credentials []Credential        `yaml:"credentials"`
}

func ParseConfigFile(filePath string) (Config, error) {
	// Read YAML file
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading YAML file: %w", err)
	}

	// Create a Config struct to store the parsed data
	var config Config

	// Unmarshal the YAML data into the Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling YAML data: %w", err)
	}

	return config, nil
}
