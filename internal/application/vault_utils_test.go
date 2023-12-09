package application

import (
	"testing"
	"vault-cluster-replication/internal/pkg/storage"
	"vault-cluster-replication/internal/pkg/storage/mocks"

	"github.com/hashicorp/vault/api"

	"github.com/stretchr/testify/assert"
)

func TestNewClusterCredential_CreatesExpectedInstance(t *testing.T) {
	mockSyncer := mocks.NewSyncer(t)
	clusterName := "testCluster"
	clusterCredential := NewClusterCredential(clusterName, mockSyncer)

	assert.Equal(t, clusterName, clusterCredential.ClusterName)
	assert.Equal(t, mockSyncer, clusterCredential.Connection)
}

func TestCreateVaultClientConfig_Success(t *testing.T) {
	storageAddr := "http://localhost:8200"
	client, err := createVaultClientConfig(storageAddr)
	assert.NotNil(t, client)
	assert.Nil(t, err)
}

func TestGetStorageClient_ReturnsExpectedSystemClient(t *testing.T) {
	mockClient := storage.Client{Sys: &api.Sys{}}
	systemClient := getStorageClient(mockClient)

	assert.NotNil(t, systemClient)
}
