package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) GetVersionHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	data := map[string]string{
		"version": app.GetVersionString(),
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			"The server encountered a problem and could not process your request",
			http.StatusInternalServerError,
		)
		return
	}
	// Append new line to JSON response
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) GetVersionString() string {
	return app.cfg.version
}
