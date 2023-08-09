package main

import (
	"fmt"
	"github.com/vrischmann/envconfig"
	"os"
	"vault-cluster-replication/internal/application"
)

type Config struct {
	ConfigPath string `envconfig:"default=tests/test-config.yaml"`
}

func main() {
	cnf := parseConfig()
	fmt.Println(cnf.ConfigPath)
	applicationConfig, err := application.ParseConfigFile(cnf.ConfigPath)
	if err != nil {
		err := fmt.Errorf("error parsing config file: %w", err)
		fmt.Println(err)
		os.Exit(1)
	}
	err = application.ConfigValidator(applicationConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	application.Run(applicationConfig)
}

func parseConfig() Config {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return conf
}
