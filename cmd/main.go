package main

import (
	"os"
	"time"
	"vault-cluster-replication/internal/application"
	"vault-cluster-replication/internal/pkg/logs"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	ConfigFilePath             string        `envconfig:"default=configs/default-conf.yaml"`
	ScheduledExecutionInterval time.Duration `envconfig:"default=5m"`
}

func main() {
	cnf, err := parseConfig()
	if err != nil {
		logs.Logger.Error(err, "error parsing config")
		os.Exit(1)
	}
	applicationConfig, err := application.ParseConfigFile(cnf.ConfigFilePath)
	if err != nil {
		logs.Logger.Error(err, "error parsing config file")
		os.Exit(1)
	}

	err = application.ConfigValidator(applicationConfig)
	if err != nil {
		logs.Logger.Error(err, "error validating config")
		os.Exit(1)
	}
	for {
		if err := application.Run(applicationConfig); err != nil {
			logs.Logger.Error(err, "error running application")
			os.Exit(1)
		}
		time.Sleep(DurationToSeconds(cnf.ScheduledExecutionInterval))
	}
}
func parseConfig() (Config, error) {
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func DurationToSeconds(duration time.Duration) time.Duration {
	return duration / time.Second * time.Second
}
