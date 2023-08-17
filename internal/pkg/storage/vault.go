// Package storage provides an interface for backup and restore of a secret backend.
// It contains a System struct that contains the client for the secret backend.
// The System struct implements the Syncer interface, which provides methods for pulling and pushing snapshots.
// The System struct also contains the SYS interface, which provides methods for creating and restoring snapshots.
package storage

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
	"vault-cluster-replication/internal/pkg/logs"
)

const filePermission = 0600

// Syncer is an interface that provides methods for pulling and pushing snapshots.
type Syncer interface {
	PullSnapshot() (string, error)
	PushSnapshot(fileName string) error
}

// SYS is an interface that provides methods for creating and restoring snapshots.
//
//go:generate mockery --name Syncer
type SYS interface {
	RaftSnapshot(snapWriter io.Writer) error
	RaftSnapshotRestore(snapReader io.Reader, force bool) error
}

// System is the struct that contains the client for the secret backend.
//
//go:generate mockery --name SYS
type System struct {
	Sys SYS
}

// NewSystemClient returns a new System struct with the provided SYS interface.
func NewSystemClient(sys SYS) *System {
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
	err = os.WriteFile(backupFile, snapshot.Bytes(), filePermission)
	if err != nil {
		return "", fmt.Errorf("unable to write snapshot to file, %s", err.Error())
	}
	logs.Logger.Info("snapshot file created: " + backupFile)

	return backupFile, nil
}

// PushSnapshot reads a snapshot from a file and restores it.
func (s System) PushSnapshot(backupFileName string) error {
	snapshot, err := os.Open(backupFileName)
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
