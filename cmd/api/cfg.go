package main

import (
	"os"
)

func createConfig() (*config, error) {
	cfg := &config{}
	cfg.port = getEnv(PORT_ENV, "8282")

	cfg.version = os.Getenv(VERSION_ENV)
	if cfg.version == "" {
		return nil, ErrVersionEnvNotSet
	}
	return cfg, nil
}

func (cfg *config) GetVersionString() string {
	return cfg.version
}
