package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	VERSION_ENV = "HIVEBOX_VERSION"
	PORT_ENV    = "HIVEBOX_PORT"
)

type application struct {
	cfg    *config
	logger *slog.Logger
}

type config struct {
	port    string
	version string
}

func main() {
	logger := createLogger()
	cfg, err := createConfig()
	if err != nil {
		panic(err)
	}
	app, err := createApplication(logger, cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on port", app.cfg.port)
	svr := app.intializeServer()

	err = svr.ListenAndServe()
	if err != nil {
		app.logger.Error("error starting server: %v", err)
	}
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

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func createApplication(
	logger *slog.Logger,
	cfg *config,
) (*application, error) {
	app := &application{
		logger: logger,
		cfg:    cfg,
	}
	return app, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
