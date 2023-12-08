package application

import (
	"testing"
	"vault-cluster-replication/internal/pkg/storage/mocks"

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
