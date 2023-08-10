package main

import (
	"fmt"
	"os"
	"vault-cluster-replication/internal/application"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	ConfigPath string `envconfig:"default=configs/default-conf.yaml"`
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

	err = application.Run(applicationConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseConfig() Config {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		fmt.Printf("err=%s\n", err)
	}
	return conf
}
