package application

import "fmt"

func ConfigValidator(cnf Config) error {
	err := cnf.EnsureActiveNodeIsNotPassive()
	if err != nil {
		return err
	}
	err = cnf.EnsureNotEmptySyncConfig()
	if err != nil {
		return err
	}
	err = cnf.EnsureNotEmptyCredentials()
	if err != nil {
		return err
	}
	err = cnf.EnsureNotEmptyCredentials()
	if err != nil {
		return err
	}

	return nil
}

// EnsureUniqueActiveNodes check if there are any duplicate Active nodes array
func (c *Config) EnsureUniqueActiveNodes() error {
	nameCount := make(map[string]int)
	for _, syncConfig := range c.Replication {
		nameCount[syncConfig.Active]++
		if nameCount[syncConfig.Active] > 1 {
			return fmt.Errorf("%s is duplicated in the SyncConfig array", syncConfig.Active)
		}
	}
	return nil
}

// EnsureActiveNodeIsNotPassive check if there is an Active node that is also a Passive node
func (c *Config) EnsureActiveNodeIsNotPassive() error {
	activeSet := make(map[string]bool)
	for _, syncConfig := range c.Replication {
		activeSet[syncConfig.Active] = true
	}

	for _, syncConfig := range c.Replication {
		for _, syncedTo := range syncConfig.SyncTo {
			if activeSet[syncedTo] {
				return fmt.Errorf("%s is an active node and cannot be a passive node", syncedTo)
			}
		}
	}
	return nil
}

// EnsureNotEmptySyncConfig check if the SyncConfig array is empty
func (c *Config) EnsureNotEmptySyncConfig() error {
	if len(c.Replication) == 0 {
		return fmt.Errorf("SyncConfig array is empty")
	}
	for _, syncConfig := range c.Replication {
		if syncConfig.Active == "" {
			return fmt.Errorf("active node is empty")
		}
		if len(syncConfig.SyncTo) == 0 {
			return fmt.Errorf("passive nodes array is empty")
		}
		for _, syncTo := range syncConfig.SyncTo {
			if syncTo == "" {
				return fmt.Errorf("passive node is empty")
			}
		}
	}
	return nil
}

// EnsureNotEmptyCredentials check if the credentials array is empty
func (c *Config) EnsureNotEmptyCredentials() error {
	if len(c.Credentials) == 0 {
		return fmt.Errorf("credentials array is empty")
	}
	for _, cluster := range c.Credentials {
		if cluster.Name == "" {
			return fmt.Errorf("cluster name is empty")
		}
		if cluster.AppRole == "" {
			return fmt.Errorf("cluster username is empty, cluster: %s", cluster.Name)
		}
		if cluster.SecretID == "" {
			return fmt.Errorf("cluster password is empty, cluster: %s", cluster.Name)
		}
	}
	return nil
}
