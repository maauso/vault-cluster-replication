package storage

import (
	"context"
	"io"

	"github.com/hashicorp/vault/api"
)

// Sys is an interface that provides methods for creating and restoring snapshots.
//go:generate mockery --name Sys
type Sys interface {
	RaftSnapshot(snapWriter io.Writer) error
	RaftSnapshotRestore(snapReader io.Reader, force bool) error
}

//go:generate mockery --name Auth
type Auth interface {
	Login(ctx context.Context, authMethod api.AuthMethod) (*api.Secret, error)
}

//go:generate mockery --name Logical
type Logical interface {
	List(path string) (*api.Secret, error)
	Delete(path string) (*api.Secret, error)
	Write(path string, data map[string]interface{}) (*api.Secret, error)
	Read(path string) (*api.Secret, error)
}
