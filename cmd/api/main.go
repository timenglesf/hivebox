package main

import (
	"fmt"
	"log"
)

const (
	VERSION_ENV = "HIVEBOX_VERSION"
	PORT_ENV    = "HIVEBOX_PORT"
)

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

	log.Fatal(svr.ListenAndServe())
}
