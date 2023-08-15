package application

import (
	"fmt"
	"vault-cluster-replication/internal/pkg/storage"
)

type ClusterCredential struct {
	ClusterName string
	Connection  storage.Syncer
}

// NewClusterCredential creates a new ClusterCredential instance
func NewClusterCredential(clusterName string, connection storage.Syncer) ClusterCredential {
	return ClusterCredential{ClusterName: clusterName, Connection: connection}
}

type ClusterCredentials []ClusterCredential

func Run(config Config) error {

	credentials, err := getClustersCredentials(config, storage.Client{})
	if err != nil {
		return err
	}

	fmt.Println(credentials)

	// Print the parsed data
	//fmt.Println("Replication:")
	/*	for _, rep := range config.Replication {
		fmt.Printf("Active: %s\n", rep.Active)
		fmt.Println("Synced To:")

		for _, synced := range rep.SyncTo {
			fmt.Printf("  - %s\n", synced)
		}
		fmt.Println()
	}*/

	return nil

}

func getClustersCredentials(config Config, client storage.Client) (ClusterCredentials, error) {
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
