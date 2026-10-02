package storage

import (
	"errors"
	"testing"
	"vault-cluster-replication/internal/pkg/storage/mocks"

	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguredVaultClient_Success(t *testing.T) {
	vaultAddr := "http://localhost:8200"
	client, err := NewConfiguredVaultClient(vaultAddr)
	assert.NotNil(t, client)
	require.NoError(t, err)
}

func TestAuthenticateWithAppRole_Success(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	secretID := &auth.SecretID{FromString: "appSecretID"}
	appRoleAuth, err := auth.NewAppRoleAuth("appRoleID", secretID)
	if err != nil {
		t.Fatal(err)
	}

	mockAuth.On("Login", mock.Anything, appRoleAuth).Once().Return(nil, nil)
	_, err = AuthenticateWithAppRole(client, "appRoleID", "appSecretID")
	require.NoError(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}

func TestAuthenticateWithAppRole_FailureInvalidAppRoleAuth(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	secretID := &auth.SecretID{FromString: "appSecretID"}
	appRoleAuth, err := auth.NewAppRoleAuth("invalidAppRoleID", secretID)
	if err != nil {
		t.Fatal(err)
	}

	mockAuth.On("Login", mock.Anything, appRoleAuth).Once().Return(nil, errors.New("invalid appRole"))
	_, err = AuthenticateWithAppRole(client, "invalidAppRoleID", "appSecretID")
	require.Error(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}

func TestAuthenticateWithAppRole_FailureInvalidSecretID(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	_, err := AuthenticateWithAppRole(client, "", "appSecretID")
	require.Error(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}

func TestAuthenticateWithAppRole_FailureInvalidAppRoleID(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	_, err := AuthenticateWithAppRole(client, "invalidAppRoleID", "")
	require.Error(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}
