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
			name: "EnsureUniqueActiveNodes: Valid",
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
			name: "EnsureUniqueActiveNodes: Error",
			fields: fields{
				Replication: []ReplicationConfig{
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2"},
					},
					{
						Active: "cluster1",
						SyncTo: []string{"cluster2", "cluster5"},
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
			if err := c.EnsureUniqueActiveNodes(); (err != nil) != tt.wantErr {
				t.Errorf("EnsureUniqueActiveNodes() error = %v, wantErr %v", err, tt.wantErr)
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
			name: "EnsureActiveNodeIsNotPassive: Valid",
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
			name: "EnsureActiveNodeIsNotPassive: Error",
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
			if err := c.EnsureActiveNodeIsNotPassive(); (err != nil) != tt.wantErr {
				t.Errorf("EnsureActiveNodeIsNotPassive() error = %v, wantErr %v", err, tt.wantErr)
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
			name: "EnsureNotEmptySyncConfig: Valid",
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
			name: "EnsureNotEmptySyncConfig: Error",
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
			name: "EnsureNotEmptySyncConfig: SyncTo is empty",
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
			name: "EnsureNotEmptySyncConfig: Active is empty",
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
			name: "EnsureNotEmptySyncConfig: Replication is empty",
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
			if err := c.EnsureNotEmptySyncConfig(); (err != nil) != tt.wantErr {
				t.Errorf("EnsureNotEmptySyncConfig() error = %v, wantErr %v", err, tt.wantErr)
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
			name: "EnsureNotEmptyCredentials: Valid",
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
			name: "EnsureNotEmptyCredentials: Error",
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
			name: "EnsureNotEmptyCredentials: ClusterName Empty",
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
			name: "EnsureNotEmptyCredentials: AppRole Empty",
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
			name: "EnsureNotEmptyCredentials: SecretID Empty",
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
			if err := c.EnsureNotEmptyCredentials(); (err != nil) != tt.wantErr {
				t.Errorf("EnsureNotEmptyCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
