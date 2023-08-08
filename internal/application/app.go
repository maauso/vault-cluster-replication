package application

import "fmt"

func Run(config Config) {
	// Print the parsed data
	fmt.Println("Replication:")
	for _, rep := range config.Replication {
		fmt.Printf("Active: %s\n", rep.Active)
		fmt.Println("Synced To:")
		for _, synced := range rep.SyncTo {
			fmt.Printf("  - %s\n", synced)
		}
		fmt.Println()
	}

	fmt.Println("Credentials:")
	for _, cred := range config.Credentials {
		fmt.Printf("Name: %s\n", cred.Name)
		fmt.Printf("AppRole: %s\n", cred.AppRole)
		fmt.Printf("SecretID: %s\n\n", cred.SecretID)
	}
}
