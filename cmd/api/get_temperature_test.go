package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureHandler(t *testing.T) {
	// os.Setenv(VERSION_ENV, "1.0.0")
	app, err := newTestApplication(t, false)
	if err != nil {
		t.Fatal(err)
	}

	svr := newTestServer(t, app.routes())

	var data map[string]interface{}
	svr.get(t, "/temperature", &data)

	assert.NotNil(t, data, nil)
}
