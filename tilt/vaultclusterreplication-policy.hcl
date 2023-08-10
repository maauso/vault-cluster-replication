path "sys/mounts" {
  capabilities = ["read", "list"]
}

path "sys/policies/acl/*" {
  capabilities = ["read"]
}

path "sys/storage/raft/snapshot*" {
  capabilities = ["create", "update", "read"]
}


path "sys/raft/snapshots/*/restore" {
  capabilities = ["update"]
}
