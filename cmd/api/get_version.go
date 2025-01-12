package main

import (
	"net/http"
)

type VersionJSONRes struct {
	Version string `json:"version"`
}

func (app *application) GetVersionHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	data := VersionJSONRes{
		Version: app.cfg.version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			"The server encountered a problem and could not process your request",
			http.StatusInternalServerError,
		)
	}
}
