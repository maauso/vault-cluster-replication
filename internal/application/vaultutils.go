package application

import "vault-cluster-replication/internal/pkg/storage"

type ClusterCredential struct {
	ClusterName string
	Connection  storage.Syncer
}

// NewClusterCredential creates a new ClusterCredential instance.
func NewClusterCredential(clusterName string, connection storage.Syncer) ClusterCredential {
	return ClusterCredential{ClusterName: clusterName, Connection: connection}
}

type ClusterCredentials []ClusterCredential

func getClustersToken(config Config, client storage.Client) (ClusterCredentials, error) {
	credentials := ClusterCredentials{}
	for _, cred := range config.Credentials {
		clientConnection, err := client.NewSysClientConnection(cred.Name, cred.AppRole, cred.SecretID)
		if err != nil {
			return nil, err
		}

		clusterCredential := NewClusterCredential(cred.Name, clientConnection)
		credentials = append(credentials, clusterCredential)
	}

	return credentials, nil
}
