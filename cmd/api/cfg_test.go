package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfig(t *testing.T) {
	tests := []struct {
		name            string
		setPort         string
		setVersion      string
		expectedPort    string
		expectedVersion string
		expectedError   error
	}{
		{
			name:            "port and version set",
			setPort:         "1000",
			setVersion:      "1.0.0",
			expectedPort:    "1000",
			expectedVersion: "1.0.0",
			expectedError:   nil,
		},
		{
			name:            "default port and version set",
			setPort:         "",
			setVersion:      "1.0.0",
			expectedPort:    "8282",
			expectedVersion: "1.0.0",
			expectedError:   nil,
		},
		{
			name:            "port and version not set",
			setPort:         "",
			setVersion:      "",
			expectedPort:    "8282",
			expectedVersion: "",
			expectedError:   ErrVersionEnvNotSet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the environment variables
			if tt.setPort != "" {
				_ = os.Setenv(PORT_ENV, tt.setPort)
			}
			defer os.Unsetenv(PORT_ENV)

			os.Setenv(VERSION_ENV, tt.setVersion)
			defer os.Unsetenv(VERSION_ENV)

			cfg, err := createConfig()
			if err != nil {
				if tt.expectedError == nil {
					t.Errorf("an unexpected error occurred: %v", err)
					return
				}
				if errors.Is(err, tt.expectedError) {
					assert.Equal(t, tt.expectedError, err, "error")
					return
				}
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}

			assert.Equal(t, tt.expectedPort, cfg.port, "port")
			assert.Equal(t, tt.expectedVersion, cfg.version, "version")
		})
	}
}

func TestGetVersionString(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		version string
	}{
		{
			name:    "version exists",
			want:    "1.0.0",
			version: "1.0.0",
		},
		{
			name:    "version does not exist",
			want:    "",
			version: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config{
				port:    "8282",
				version: tt.version,
			}

			if got := cfg.GetVersionString(); got != tt.want {
				t.Errorf("GetVersionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
