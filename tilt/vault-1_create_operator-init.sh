#!/bin/bash



# Check if "sealed" field is false using jq
sealed=$(vault status -format json | jq -r '.sealed')

if [ "$sealed" = "true" ]; then
  # Init Vault cluster
  vault operator init -key-shares=1 -key-threshold=1 -format=json > vault-1_unseal_keys.json
fi



# Unseal Vault cluster
vault operator unseal $(jq -r '.unseal_keys_b64[]' vault-1_unseal_keys.json )

export VAULT_TOKEN=$(jq -r '.root_token' vault-1_unseal_keys.json)

# Enable AppRole Auth Method
vault auth enable approle

# Create Vault policy
vault policy write vaultclusterreplication-policy ./vaultclusterreplication-policy.hcl

# Create Vault Role

vault write auth/approle/role/vaultClusterReplication policies=vaultclusterreplication-policy
vault write auth/approle/role/vaultClusterReplication/role-id role_id=vaultClusterReplication secret_id_ttl=0 token_num_uses=0 token_ttl=900 token_max_ttl=0 token_policies=vaultclusterreplication-policy
vault read auth/approle/role/vaultClusterReplication/role-id
vault write auth/approle/role/vaultClusterReplication/custom-secret-id secret_id=root
