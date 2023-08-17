package application

import (
	"vault-cluster-replication/internal/pkg/storage"
)

func Run(config Config) error {
	credentials, err := getClustersToken(config, storage.Client{})
	if err != nil {
		return err
	}

	err = replicator(config, credentials)
	if err != nil {
		return err
	}

	return nil
}

func replicator(config Config, credentials ClusterCredentials) error {
	for _, replication := range config.Replication {
		sync := getClusterCredentials(replication.Active, credentials)
		backup, err := sync.PullSnapshot()
		if err != nil {
			return err
		}
		for _, syncTo := range replication.SyncTo {
			syncToCluster := getClusterCredentials(syncTo, credentials)
			err := syncToCluster.PushSnapshot(backup)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getClusterCredentials(clusterURL string, credentials ClusterCredentials) storage.Syncer {
	for _, credential := range credentials {
		if credential.ClusterName == clusterURL {
			connectionCredentials := credential.Connection

			return connectionCredentials
		}
	}

	return nil
}
