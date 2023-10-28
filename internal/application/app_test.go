package application

import (
	"fmt"
	"reflect"
	"testing"
	"vault-cluster-replication/internal/pkg/storage"
	"vault-cluster-replication/internal/pkg/storage/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReplicator_ok(t *testing.T) {
	mockSyncer := mocks.NewSyncer(t)
	config := Config{
		Replication: []ReplicationConfig{
			{
				Active: "cluster1",
				SyncTo: []string{"cluster2"},
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

	credentials := ClusterCredentials{
		{
			ClusterName: "cluster1",
			Connection:  mockSyncer,
		},
		{
			ClusterName: "cluster2",
			Connection:  mockSyncer,
		},
	}

	mockSyncer.On("PullSnapshot").Once().Return("backupDone", nil)
	mockSyncer.On("PushSnapshot", "backupDone").Once().Return(nil)
	err := performReplication(config, credentials)
	assert.Equal(t, nil, err)
	mockSyncer.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSyncer)
}

func TestReplicator_PullSnapshot_Error(t *testing.T) {
	mockSyncer := mocks.NewSyncer(t)
	config := Config{
		Replication: []ReplicationConfig{
			{
				Active: "cluster1",
				SyncTo: []string{"cluster2"},
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

	credentials := ClusterCredentials{
		{
			ClusterName: "cluster1",
			Connection:  mockSyncer,
		},
		{
			ClusterName: "cluster2",
			Connection:  mockSyncer,
		},
	}

	mockSyncer.On("PullSnapshot").Once().Return("", fmt.Errorf("error PullSnapshot"))
	err := performReplication(config, credentials)
	assert.EqualError(t, err, "error PullSnapshot")
	mockSyncer.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSyncer)
}

func TestReplicator_PushSnapshot_Error(t *testing.T) {
	mockSyncer := mocks.NewSyncer(t)
	config := Config{
		Replication: []ReplicationConfig{
			{
				Active: "cluster1",
				SyncTo: []string{"cluster2"},
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

	credentials := ClusterCredentials{
		{
			ClusterName: "cluster1",
			Connection:  mockSyncer,
		},
		{
			ClusterName: "cluster2",
			Connection:  mockSyncer,
		},
	}

	mockSyncer.On("PullSnapshot").Once().Return("backupDone", nil)
	mockSyncer.On("PushSnapshot", "backupDone").Once().Return(fmt.Errorf("error PushSnapshot"))
	err := performReplication(config, credentials)
	assert.EqualError(t, err, "error PushSnapshot")
	mockSyncer.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockSyncer)
}

func Test_retrieveClusterCredentials(t *testing.T) {
	mockSyncer := mocks.NewSyncer(t)
	type args struct {
		clusterURL  string
		credentials ClusterCredentials
	}
	tests := []struct {
		name string
		args args
		want storage.Syncer
	}{
		{
			name: "Cluster credentials found",
			args: args{
				clusterURL: "cluster1",
				credentials: ClusterCredentials{
					{
						ClusterName: "cluster1",
						Connection:  mockSyncer,
					},
				},
			},
			want: storage.Syncer(mockSyncer),
		},
		{
			name: "Cluster credentials not found",
			args: args{
				clusterURL: "cluster3",
				credentials: ClusterCredentials{
					{
						ClusterName: "cluster1",
						Connection:  mockSyncer,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retrieveClusterCredentials(tt.args.clusterURL, tt.args.credentials); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveClusterCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	// Backup the original functions
	originalReplicator := replicateFunc
	originalGenerateConfigs := generateClusterConfigsFunc
	defer func() {
		replicateFunc = originalReplicator
		generateClusterConfigsFunc = originalGenerateConfigs
	}()

	// Create a sample configuration
	config := Config{
		Replication: []ReplicationConfig{
			{
				Active: "cluster1",
				SyncTo: []string{"cluster2", "cluster3"},
			},
		},
	}

	// Replace the performReplication function with a mock function
	replicateFunc = func(config Config, credentials ClusterCredentials) error {
		// Simulate the behavior of the mock performReplication
		// You can use assertions here if you need
		return nil
	}

	// Replace generateClustersConfigs with a mock function
	generateClusterConfigsFunc = func(config Config) (ClusterCredentials, error) {
		// Simulate the behavior of the mock generateClustersConfigs
		// You can use assertions here if you need
		return ClusterCredentials{}, nil
	}

	// Call the Run function
	err := Run(config)

	// Check if the Run function returns the expected error (nil in this case)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

}

func TestRun_generateConfigs_Error(t *testing.T) {
	// Backup the original functions

	originalGenerateConfigs := generateClusterConfigsFunc
	defer func() {
		generateClusterConfigsFunc = originalGenerateConfigs
	}()

	// Create a sample configuration
	config := Config{}

	// Replace generateClustersConfigs with a mock function
	generateClusterConfigsFunc = func(config Config) (ClusterCredentials, error) {
		// Simulate the behavior of the mock generateClustersConfigs
		// You can use assertions here if you need
		return nil, fmt.Errorf("error generateClustersConfigs")
	}

	// Call the Run function
	err := Run(config)

	// Check if the Run function returns the expected error (nil in this case)
	if err == nil {
		t.Errorf("Expected error, but got: %v", err)
	}
}

func TestRun_replicateFunc_Error(t *testing.T) {
	// Backup the original functions
	originalReplicator := replicateFunc
	originalGenerateConfigs := generateClusterConfigsFunc
	defer func() {
		replicateFunc = originalReplicator
		generateClusterConfigsFunc = originalGenerateConfigs
	}()

	// Create a sample configuration
	config := Config{}
	// Replace generateClustersConfigs with a mock function
	generateClusterConfigsFunc = func(config Config) (ClusterCredentials, error) {
		// Simulate the behavior of the mock generateClustersConfigs
		// You can use assertions here if you need
		return ClusterCredentials{}, nil
	}
	// Replace the performReplication function with a mock function
	replicateFunc = func(config Config, credentials ClusterCredentials) error {
		// Simulate the behavior of the mock performReplication
		// You can use assertions here if you need
		return fmt.Errorf("error performReplication")
	}

	// Call the Run function
	err := Run(config)

	// Check if the Run function returns the expected error (nil in this case)
	if err == nil {
		t.Errorf("Expected error, but got: %v", err)
	}
}
