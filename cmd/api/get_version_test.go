package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersionHandler(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{
			name:    "version 1.0.0",
			version: "1.0.0",
		},
		{
			name:    "version 1.0.1",
			version: "1.0.1",
		},
	}

	for _, tt := range tests {
		t.Run("Test GetVersionHandler", func(t *testing.T) {
			os.Setenv(PORT_ENV, "8080")
			defer os.Unsetenv(PORT_ENV)

			os.Setenv(VERSION_ENV, tt.version)
			defer os.Unsetenv(VERSION_ENV)

			app, err := newTestApplication(t, true)
			if err != nil {
				t.Fatal(err)
			}

			svr := newTestServer(t, app.routes())

			var data struct {
				Version string `json:"version"`
			}
			//	var data map[string]interface{}

			status, _, err := svr.get(t, "/version", &data)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 200, status)
			assert.Equal(t, tt.version, data.Version)
		})
	}
}
