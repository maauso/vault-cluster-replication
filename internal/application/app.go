package application

import (
	"time"
	"vault-cluster-replication/internal/pkg/logs"
	"vault-cluster-replication/internal/pkg/storage"
)

var (
	replicateFunc              = performReplication
	generateClusterConfigsFunc = generateClustersConfigs
)

func Run(config Config) error {
	currentTime := time.Now()
	logs.Logger.Info("Starting replication", "time", currentTime.Format(time.RFC3339))

	credentials, err := generateClusterConfigsFunc(config)
	if err != nil {
		return err
	}

	err = replicateFunc(config, credentials)
	if err != nil {
		return err
	}

	return nil
}

func performReplication(config Config, credentials ClusterCredentials) error {
	for _, replication := range config.Replication {
		sync := retrieveClusterCredentials(replication.Active, credentials)
		backup, err := sync.PullSnapshot()
		if err != nil {
			return err
		}
		for _, syncTo := range replication.SyncTo {
			syncToCluster := retrieveClusterCredentials(syncTo, credentials)
			err := syncToCluster.PushSnapshot(backup)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func retrieveClusterCredentials(clusterURL string, credentials ClusterCredentials) storage.Syncer {
	for _, credential := range credentials {
		if credential.ClusterName == clusterURL {
			connectionCredentials := credential.Connection

			return connectionCredentials
		}
	}

	return nil
}
