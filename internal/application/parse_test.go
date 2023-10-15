package application

import (
	"reflect"
	"testing"
)

func TestParseConfigFile_OK(t *testing.T) {
	t.Parallel()
	filePath := "../../tests/test-config.yaml" // Create a test YAML file with sample content for testing

	expectedConfig := Config{
		Replication: []ReplicationConfig{
			{
				Active: "cluster1",
				SyncTo: []string{"cluster2"},
			},
			{
				Active: "cluster3",
				SyncTo: []string{"cluster4", "cluster5"},
			},
		},
		Credentials: []Credential{
			{
				Name:     "cluster1",
				AppRole:  "vault-cluster-replication",
				SecretID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			{
				Name:     "cluster2",
				AppRole:  "vault-cluster-replication",
				SecretID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			{
				Name:     "cluster3",
				AppRole:  "vault-cluster-replication",
				SecretID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			{
				Name:     "cluster4",
				AppRole:  "vault-cluster-replication",
				SecretID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
			{
				Name:     "cluster5",
				AppRole:  "vault-cluster-replication",
				SecretID: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			},
		},
	}

	config, err := ParseConfigFile(filePath)
	if err != nil {
		t.Errorf("Error parsing config file: %v", err)
	}

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("Parsed config does not match expected:\nExpected: %+v\nGot: %+v", expectedConfig, config)
	}
}

func TestParseConfigFile_Error_FileNotExist(t *testing.T) {
	t.Parallel()
	filePath := ""

	_, err := ParseConfigFile(filePath)
	if err == nil {
		t.Errorf("Expected error parsing config file, got nil")
	}
}

func TestParseConfigFile_Error_InvalidYAML(t *testing.T) {
	t.Parallel()
	filePath := "../../tests/test-config-invalid.yaml" // Create a test YAML file with sample content for testing

	_, err := ParseConfigFile(filePath)
	if err == nil {
		t.Errorf("Expected error parsing config file, got nil")
	}
}
