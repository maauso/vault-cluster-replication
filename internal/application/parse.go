package application

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ReplicationConfig struct {
	Active string   `yaml:"active"`
	SyncTo []string `yaml:"sync_to"`
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

func NewConfig(replication []ReplicationConfig, credentials []Credential) Config {
	return Config{
		Replication: replication,
		Credentials: credentials,
	}
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
