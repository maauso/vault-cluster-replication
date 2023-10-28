package storage

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

const timeoutSeconds = 3600

type Client struct {
	Logical Logical
	Sys     Sys
	Auth    Auth
}

func NewClient(logical Logical, sys Sys, auth Auth) Client {
	return Client{Logical: logical, Sys: sys, Auth: auth}
}

// NewConfiguredVaultClient new Vault client, authenticates with the AppRole, and returns a new Vault storage client.
func NewConfiguredVaultClient(vaultAddr string) (*api.Client, error) {
	httpClient := &http.Client{
		Timeout: timeoutSeconds * time.Second,
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

	return client, nil
}

// AuthenticateWithAppRole authenticates the client using the AppRole authentication method.
// It takes a client, an appRoleID string and an appSecretID string as input parameters.
// It returns the authenticated client and an error if any.
func AuthenticateWithAppRole(client Client, appRoleID string, appSecretID string) (Client, error) {
	secretID := &auth.SecretID{FromString: appSecretID}
	appRoleAuth, err := auth.NewAppRoleAuth(appRoleID, secretID)
	if err != nil {
		return Client{}, err
	}
	_, err = client.Auth.Login(context.TODO(), appRoleAuth)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}
