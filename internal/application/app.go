package application

import (
	"vault-cluster-replication/internal/pkg/storage"
)

func Run(config Config) error {
	_, err := getClustersCredentials(config, storage.Client{})
	if err != nil {
		return err
	}

	return nil

}
