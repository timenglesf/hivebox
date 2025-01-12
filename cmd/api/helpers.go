package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/timenglesf/hivebox/internal/sensebox"
)

// //////////////////////
// Application Struct //
// //////////////////////

type application struct {
	cfg      *config
	logger   *slog.Logger
	sensebox sensebox.SenseboxInterface
}

// Create an application struct to hold the logger and configuration settings
func createApplication(
	logger *slog.Logger,
	cfg *config,
) (*application, error) {
	app := &application{
		logger:   logger,
		cfg:      cfg,
		sensebox: &sensebox.SenseboxService{},
	}
	return app, nil
}

// Initialize the HTTP server with configuration settings
func (app *application) intializeServer() *http.Server {
	return &http.Server{
		Addr:         ":" + app.cfg.port,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}
}

// createLogger initializes and returns a new slog.Logger instance.
func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

//////////////////
// JSON Helpers //
//////////////////

// envelope is a type alias for a map with string keys and values of any type.
type envelope map[string]any

// writeJSON sends a JSON response with the specified status code and data.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(js); err != nil {
		return err
	}

	return nil
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is empty, it returns the defaultValue.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
