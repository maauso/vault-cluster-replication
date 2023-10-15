package application

import (
	"vault-cluster-replication/internal/pkg/storage"

	"github.com/hashicorp/vault/api"
)

type ClusterCredential struct {
	ClusterName string
	Connection  storage.Syncer
}

// NewClusterCredential creates a new ClusterCredential instance.
func NewClusterCredential(clusterName string, connection storage.Syncer) ClusterCredential {
	return ClusterCredential{ClusterName: clusterName, Connection: connection}
}

type ClusterCredentials []ClusterCredential

// generateClustersConfigs generates a list of ClusterCredentials based on the provided Config.
// It creates a Vault client configuration for each credential, retrieves a cluster token, and creates a new ClusterCredential
// with the retrieved token and credential name.
// Returns a list of ClusterCredentials and an error if any.
func generateClustersConfigs(config Config) (ClusterCredentials, error) {
	credentials := ClusterCredentials{}
	for _, cred := range config.Credentials {
		clientConfig, err := createVaultClientConfig(cred.Name)
		if err != nil {
			return nil, err
		}
		storageClient, err := getClusterToken(clientConfig, cred.AppRole, cred.SecretID)
		if err != nil {
			return nil, err
		}

		systemClient := getStorageClient(storageClient)

		if err != nil {
			return nil, err
		}

		clusterCredential := NewClusterCredential(cred.Name, systemClient)
		credentials = append(credentials, clusterCredential)
	}

	return credentials, nil
}

func createVaultClientConfig(storageAddr string) (*api.Client, error) {
	storageClientConfig, err := storage.ClientConfig(storageAddr)
	if err != nil {
		return nil, err
	}

	return storageClientConfig, nil
}

func getClusterToken(clientConfig *api.Client, appRoleID string, appSecretID string) (storage.Client, error) {
	vaultClient := storage.NewClient(clientConfig.Logical(), clientConfig.Sys(), clientConfig.Auth())
	client, err := storage.ClientAppRoleAuthentication(vaultClient, appRoleID, appSecretID)
	if err != nil {
		return storage.Client{}, err
	}

	return client, nil
}

func getStorageClient(storageClient storage.Client) *storage.System {
	systemClient := storage.NewSystem(storageClient.Sys)

	return systemClient
}
