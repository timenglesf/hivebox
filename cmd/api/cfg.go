package main

import (
	"fmt"
	"os"
)

const (
	DEFAULT_ID_1 = "6442621e0331e20009b859bd"
	DEFAULT_ID_2 = "62d12626e72c70001bcb61bd"
	DEFAULT_ID_3 = "6442648b0331e20009b9d50f"
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
		DEFAULT_ID_1,
		DEFAULT_ID_2,
		DEFAULT_ID_3,
	}

	ids := make([]string, len(defaultSenseboxIds))
	for i := range ids {
		envName := fmt.Sprintf("%s%d", SENSEBOX_ID_ENV, i+1)
		id := os.Getenv(envName)

		if id == "" {
			ids[i] = defaultSenseboxIds[i]
		} else {
			ids[i] = id
		}

	}

	return ids
}
