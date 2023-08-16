package storage

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
	"vault-cluster-replication/internal/pkg/logs"
)

//go:generate mockery --name Syncer
type Syncer interface {
	PullSnapshot() (string, error)
	PushSnapshot(fileName string) error
}

// RaftSnapshotManager is the interface for backup and restore

//go:generate mockery --name SYS
type SYS interface {
	RaftSnapshot(snapWriter io.Writer) error
	RaftSnapshotRestore(snapReader io.Reader, force bool) error
}

// System is the struct that contains the client for the secret backend
type System struct {
	Sys SYS
}

func NewSystemClient(sys SYS) *System {
	return &System{Sys: sys}
}

func (s System) PullSnapshot() (string, error) {
	backupFile := fmt.Sprintf("%d", time.Now().UnixNano())
	// Create io.Writer to write the snapshot to a file
	var snapshot bytes.Buffer
	err := s.Sys.RaftSnapshot(&snapshot)
	if err != nil {
		return "", fmt.Errorf("unable to generate snapshot, %s", err.Error())
	}
	logs.Logger.Info("writing snapshot locally in " + backupFile)
	err = os.WriteFile(backupFile, snapshot.Bytes(), 0600)
	if err != nil {
		return "", fmt.Errorf("unable to write snapshot to file, %s", err.Error())
	}
	logs.Logger.Info("snapshot file created: " + backupFile)
	return backupFile, nil
}

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
