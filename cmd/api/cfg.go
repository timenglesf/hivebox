package main

import (
	"fmt"
	"os"
)

type config struct {
	port     string
	version  string
	sensebox struct {
		ids []string
	}
}

func createConfig() (*config, error) {
	cfg := &config{}
	cfg.port = getEnv(PORT_ENV, "8282")

	cfg.version = os.Getenv(VERSION_ENV)
	if cfg.version == "" {
		return nil, ErrVersionEnvNotSet
	}

	cfg.sensebox.ids = getSenseBoxIds()

	return cfg, nil
}

func (cfg *config) GetVersionString() string {
	return cfg.version
}

func getSenseBoxIds() []string {
	defaultSenseboxIds := []string{
		"6442621e0331e20009b859bd",
		"62d12626e72c70001bcb61bd",
		"6442648b0331e20009b9d50f",
	}

	ids := make([]string, len(defaultSenseboxIds))
	for i := range ids {
		id := os.Getenv(fmt.Sprintf("%s%d", SENSEBOX_ID_ENV, i+1))

		if id != "" {
			ids[i] = defaultSenseboxIds[i]
		}

		ids[i] = id
	}

	return ids
}
