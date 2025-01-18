package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/timenglesf/hivebox/internal/sensebox"
)

type tempStatus string

var (
	tempStatusCold tempStatus = "Too Cold"
	tempStatusGood tempStatus = "Good"
	tempStatusHot  tempStatus = "Too Hot"
)

// GetTemperatureHandler handles the HTTP request for getting temperature data.
func (app *application) GetTemperatureHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Fetch Sensebox data
	data, err := app.GetSenseboxData()
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Parse temperature data
	var temps []float64
	for _, d := range data {
		t, err := app.parseTemperaturData(d)
		if err != nil {
			app.logger.Error(err.Error())
			continue
		}
		temps = append(temps, t)
	}

	averageTemp := average(temps)
	tempStatus := getTempStatus(averageTemp)

	js := envelope{"temperature": averageTemp, "status": tempStatus}

	// Write JSON response
	err = app.writeJSON(w, http.StatusOK, js, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			"The server encountered a problem and could not process your request",
			http.StatusInternalServerError,
		)
	}
}

// GetSenseboxData retrieves data from multiple Sensebox devices concurrently.
func (app *application) GetSenseboxData() ([]sensebox.SenseboxJSON, error) {
	// Create a slice to hold the Sensebox data
	var data []sensebox.SenseboxJSON
	// Create channels to receive data and errors
	dataChan := make(chan sensebox.SenseboxJSON, len(app.cfg.sensebox.ids))
	errChan := make(chan error, len(app.cfg.sensebox.ids))

	// Create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(len(app.cfg.sensebox.ids))

	for _, id := range app.cfg.sensebox.ids {
		// Create an anonymous goroutine to fetch data from each Sensebox device
		go func(id string) {
			defer wg.Done()
			senseboxData, err := app.sensebox.GetSenseboxData(id)
			if err != nil {
				errChan <- err
				return
			}
			dataChan <- senseboxData
		}(id)
	}

	wg.Wait()
	close(dataChan)
	close(errChan)

	for d := range dataChan {
		data = append(data, d)
	}

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return data, nil
}

// parseTemperaturData parses the temperature data from a SenseboxJSON object.
func (app *application) parseTemperaturData(data sensebox.SenseboxJSON) (float64, error) {
	var t float64
	var err error
	for _, s := range data.Sensors {
		if strings.ToLower(s.Title) == "temperatur" || strings.ToLower(s.Title) == "temperature" {
			if time.Since(s.LastMeasurement.CreatedAt) <= time.Hour {
				t, err = strconv.ParseFloat(s.LastMeasurement.Value, 64)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return t, nil
}

// average calculates the average of a slice of float64 numbers.
func average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	avg := sum / float64(len(numbers))
	return math.Round(avg*100) / 100
}

func getTempStatus(temp float64) tempStatus {
	switch {
	case temp < 11:
		return tempStatusCold
	case temp > 36:
		return tempStatusHot
	default:
		return tempStatusGood
	}
}
