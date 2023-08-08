package application

import (
	"reflect"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	filePath := "../../tests/test-config.yaml" // Create a test YAML file with sample content for testing

	expectedConfig := Config{
		Replication: []ReplicationConfig{
			{
				Active:   "cluster1",
				SyncedTo: []string{"cluster2"},
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
