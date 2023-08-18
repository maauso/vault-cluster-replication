package application

import "testing"

func TestConfig_EnsureUniqueActiveNodes(t *testing.T) {
	type fields struct {
		Replication []ReplicationConfig
		Credentials []Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ensureUniqueActiveNodes: Valid",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster4", "cluster5"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ensureUniqueActiveNodes: Error",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster5",
						SyncTo: []string{"cluster6"},
					},
					{
						Active: "cluster5",
						SyncTo: []string{"cluster6", "cluster7"},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(tt.fields.Replication, tt.fields.Credentials)
			if err := c.ensureUniqueActiveNodes(); (err != nil) != tt.wantErr {
				t.Errorf("ensureUniqueActiveNodes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_EnsureActiveNodeIsNotPassive(t *testing.T) {
	type fields struct {
		Replication []ReplicationConfig
		Credentials []Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ensureActiveNodeIsNotPassive: Valid",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster4", "cluster5"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ensureActiveNodeIsNotPassive: Error",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster1", "cluster5"},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Replication: tt.fields.Replication,
				Credentials: tt.fields.Credentials,
			}
			if err := c.ensureActiveNodeIsNotPassive(); (err != nil) != tt.wantErr {
				t.Errorf("ensureActiveNodeIsNotPassive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_EnsureNotEmptySyncConfig(t *testing.T) {
	type fields struct {
		Replication []ReplicationConfig
		Credentials []Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ensureNotEmptySyncConfig: Valid",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster4", "cluster5"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ensureNotEmptySyncConfig: Error",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster1", "cluster5"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptySyncConfig: SyncTo is empty",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2", ""},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptySyncConfig: Active is empty",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "",
						SyncTo: []string{"cluster2"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptySyncConfig: Replication is empty",
			fields: fields{
				Replication: []ReplicationConfig{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Replication: tt.fields.Replication,
				Credentials: tt.fields.Credentials,
			}
			if err := c.ensureNotEmptySyncConfig(); (err != nil) != tt.wantErr {
				t.Errorf("ensureNotEmptySyncConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_EnsureNotEmptyCredentials(t *testing.T) {
	type fields struct {
		Replication []ReplicationConfig
		Credentials []Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ensureNotEmptyCredentials: Valid",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster4", "cluster5"},
					},
				},
				Credentials: []Credential{
					{
						Name:     "cluster1",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster2",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster3",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster4",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster5",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ensureNotEmptyCredentials: Error",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster3",
						SyncTo: []string{"cluster4", "cluster5"},
					},
				},
				Credentials: []Credential{
					{
						Name:     "cluster1",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster2",
						AppRole:  "vault-cluster-replication",
						SecretID: "",
					},
					{
						Name:     "cluster3",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster4",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
					{
						Name:     "cluster5",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptyCredentials: ClusterName Empty",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
				},
				Credentials: []Credential{
					{
						Name:     "",
						AppRole:  "vault-cluster-replication",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptyCredentials: AppRole Empty",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
				},
				Credentials: []Credential{
					{
						Name:     "cluster1",
						AppRole:  "",
						SecretID: "xxxx-xxxx-xxxx-xxxx",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ensureNotEmptyCredentials: SecretID Empty",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
				},
				Credentials: []Credential{
					{
						Name:     "cluster1",
						AppRole:  "vault-cluster-replication",
						SecretID: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Replication: tt.fields.Replication,
				Credentials: tt.fields.Credentials,
			}
			if err := c.ensureNotEmptyCredentials(); (err != nil) != tt.wantErr {
				t.Errorf("ensureNotEmptyCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
