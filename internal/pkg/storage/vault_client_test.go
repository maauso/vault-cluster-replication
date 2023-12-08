package storage

import (
	"context"
	"errors"
	"testing"
	"vault-cluster-replication/internal/pkg/storage/mocks"

	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewConfiguredVaultClient_Success(t *testing.T) {
	vaultAddr := "http://localhost:8200"
	client, err := NewConfiguredVaultClient(vaultAddr)
	assert.NotNil(t, client)
	assert.Nil(t, err)
}

func TestAuthenticateWithAppRole_Success(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	secretID := &auth.SecretID{FromString: "appSecretID"}
	appRoleAuth, err := auth.NewAppRoleAuth("appRoleID", secretID)
	if err != nil {
		t.Fatal(err)
	}

	mockAuth.On("Login", context.TODO(), appRoleAuth).Once().Return(nil, nil)
	_, err = AuthenticateWithAppRole(client, "appRoleID", "appSecretID")
	assert.Equal(t, nil, err)
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

	mockAuth.On("Login", context.TODO(), appRoleAuth).Once().Return(nil, errors.New("invalid appRole"))
	_, err = AuthenticateWithAppRole(client, "invalidAppRoleID", "appSecretID")
	assert.NotNil(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}

func TestAuthenticateWithAppRole_FailureInvalidSecretID(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	_, err := AuthenticateWithAppRole(client, "", "appSecretID")
	assert.NotNil(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}

func TestAuthenticateWithAppRole_FailureInvalidAppRoleID(t *testing.T) {
	mockAuth := mocks.NewAuth(t)
	client := NewClient(nil, nil, mockAuth)
	_, err := AuthenticateWithAppRole(client, "invalidAppRoleID", "")
	assert.NotNil(t, err)
	mockAuth.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, mockAuth)
}
