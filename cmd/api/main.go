package main

import (
	"log"
)

const (
	VERSION_ENV     = "HIVEBOX_VERSION"
	PORT_ENV        = "HIVEBOX_PORT"
	SENSEBOX_ID_ENV = "SENSEBOX_ID_"
)

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

	app.logger.Info("server starting", "port", app.cfg.port)
	svr := app.intializeServer()

	log.Fatal(svr.ListenAndServe())
}
