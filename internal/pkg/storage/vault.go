// Package storage provides an interface for backup and restore of a secret backend.
// It contains a syncerClient struct that contains the client for the secret backend.
// The syncerClient struct implements the Syncer interface, which provides methods for pulling and pushing snapshots.
// The syncerClient struct also contains the Sys interface, which provides methods for creating and restoring snapshots.
package storage

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"vault-cluster-replication/internal/pkg/logs"
)

const filePermission = 0o600

// Syncer is an interface that provides methods for pulling and pushing snapshots.
//
//go:generate mockery --name Syncer
type Syncer interface {
	PullSnapshot() (string, error)
	PushSnapshot(fileName string) error
}

// System is the struct that contains the client for the secret backend.
type System struct {
	Sys Sys
}

func NewSystem(sys Sys) *System {
	return &System{Sys: sys}
}

// PullSnapshot generates a snapshot and writes it to a file.
// It returns the name of the file where the snapshot was written.
func (s System) PullSnapshot() (string, error) {
	backupFile := fmt.Sprintf("%d", time.Now().UnixNano())
	// Create io.Writer to write the snapshot to a file
	var snapshot bytes.Buffer
	err := s.Sys.RaftSnapshot(&snapshot)
	if err != nil {
		return "", fmt.Errorf("unable to generate snapshot, %s", err.Error())
	}
	logs.Logger.Info("writing snapshot locally in " + backupFile)
	err = writerFile(backupFile, snapshot.Bytes())
	if err != nil {
		return "", fmt.Errorf("unable to write snapshot to file, %s", err.Error())
	}
	logs.Logger.Info("snapshot file created: " + backupFile)

	return backupFile, nil
}

// PushSnapshot reads a snapshot from a file and restores it.
func (s System) PushSnapshot(backupFileName string) error {
	snapshot, err := openFile(backupFileName)
	if err != nil {
		return fmt.Errorf("unable to open snapshot file, %s", err.Error())
	}
	logs.Logger.Info("restoring snapshot from file: " + backupFileName)
	err = s.Sys.RaftSnapshotRestore(snapshot, true)
	if err != nil {
		return fmt.Errorf("unable to restore snapshot, %s", err.Error())
	}
	logs.Logger.Info("snapshot restored successfully")

	return nil
}

func writerFile(backupFile string, snapshot []byte) error {
	err := os.WriteFile(backupFile, snapshot, filePermission)
	if err != nil {
		return err
	}

	return nil
}

func openFile(backupFileName string) (*os.File, error) {
	snapshot, err := os.Open(backupFileName)
	if err != nil {
		return nil, err
	}

	return snapshot, nil
}
