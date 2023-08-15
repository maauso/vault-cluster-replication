package storage

import (
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

type Login interface {
	NewSysClientConnection(vaultAddr, appRoleID, appSecretID string) (Syncer, error)
}

type Client struct {
	Login
}

// It creates a new Vault client, authenticates with the AppRole, and returns a new Vault storage client
func (c *Client) NewSysClientConnection(vaultAddr, appRoleID, appSecretID string) (Syncer, error) {
	httpClient := &http.Client{
		Timeout: 3600 * time.Second,
	}

	client, err := api.NewClient(
		&api.Config{
			Address:    vaultAddr,
			HttpClient: httpClient,
		},
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   appRoleID,
		"secret_id": appSecretID,
	})
	if err != nil {
		return nil, err
	}
	client.SetToken(resp.Auth.ClientToken)

	return NewSystemClient(client.Sys()), nil
}