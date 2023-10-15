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

func getClustersConfigs(config Config) (ClusterCredentials, error) {
	credentials := ClusterCredentials{}
	for _, cred := range config.Credentials {
		clientConfig, err := getStorageClientConfig(cred.Name)
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

func getStorageClientConfig(storageAddr string) (*api.Client, error) {
	storageClientConfig, err := storage.ClientConfig(storageAddr)
	if err != nil {
		return nil, err
	}

	return storageClientConfig, nil
}

func getClusterToken(clientConfig *api.Client, appRoleID string, appSecretID string) (storage.Client, error) {
	vaultClient := storage.NewClient(clientConfig.Logical(), clientConfig.Sys(), clientConfig.Auth())
	client, err := storage.ClientLogin(vaultClient, appRoleID, appSecretID)
	if err != nil {
		return storage.Client{}, err
	}

	return client, nil
}

func getStorageClient(storageClient storage.Client) *storage.System {
	systemClient := storage.NewSystem(storageClient.Sys)

	return systemClient
}
