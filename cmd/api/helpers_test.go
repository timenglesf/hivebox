package main

import (
	"io"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateApplication(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	cfg := &config{}

	app, err := createApplication(logger, cfg)
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, logger, app.logger)
	assert.Equal(t, cfg, app.cfg)
}

func TestInitializeServer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	cfg := &config{port: "8080"}
	app := &application{logger: logger, cfg: cfg}

	server := app.intializeServer()

	assert.NotNil(t, server)
	assert.Equal(t, ":8080", server.Addr)
	assert.NotNil(t, server.Handler)
	assert.Equal(t, time.Minute, server.IdleTimeout)
	assert.Equal(t, 5*time.Second, server.ReadTimeout)
	assert.Equal(t, 10*time.Second, server.WriteTimeout)
	assert.NotNil(t, server.ErrorLog)
}

func TestCreateLogger(t *testing.T) {
	logger := createLogger()

	assert.NotNil(t, logger)
	_, ok := logger.Handler().(*slog.TextHandler)
	assert.True(t, ok)
}

func TestWriteJSON(t *testing.T) {
	t.Run("Test WriteJSON", func(t *testing.T) {
	})
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        string
		defaultValue string
		keyExists    bool
	}{
		{
			name:         "value for key is defined",
			key:          "PORT",
			value:        "8080",
			defaultValue: "0000",
			keyExists:    true,
		},
		{
			name:         "value for key is undefined",
			key:          "PORT",
			value:        "",
			defaultValue: "0000",
			keyExists:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.key, tt.value)
			defer os.Unsetenv(tt.key)

			v := getEnv(tt.key, tt.defaultValue)

			if !tt.keyExists {
				assert.Equal(t, tt.defaultValue, v)
				return
			}

			assert.Equal(t, tt.value, v)
		})
	}
}
