package main

import (
	"os"

	"vault-cluster-replication/internal/application"
	"vault-cluster-replication/internal/pkg/logs"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	ConfigPath string `envconfig:"default=configs/default-conf.yaml"`
}

func main() {
	cnf := parseConfig()
	applicationConfig, err := application.ParseConfigFile(cnf.ConfigPath)
	if err != nil {
		logs.Logger.Error(err, "error parsing config file")
		os.Exit(1)
	}
	err = application.ConfigValidator(applicationConfig)
	if err != nil {
		logs.Logger.Error(err, "error validating config")
		os.Exit(1)
	}

	err = application.Run(applicationConfig)
	if err != nil {
		logs.Logger.Error(err, "error running application")
		os.Exit(1)
	}
}

func parseConfig() Config {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		logs.Logger.Error(err, "error parsing config")
	}

	return conf
}
