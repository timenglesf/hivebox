package sensebox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	SenseboxService   struct{}
	SenseboxInterface interface {
		GetSenseboxData(id string) (SenseboxJSON, error)
	}
)

type SenseboxJSON struct {
	ID              string          `json:"_id,omitempty"`
	CreatedAt       time.Time       `json:"createdAt,omitempty"`
	UpdatedAt       time.Time       `json:"updatedAt,omitempty"`
	Name            string          `json:"name,omitempty"`
	CurrentLocation CurrentLocation `json:"currentLocation,omitempty"`
	Exposure        string          `json:"exposure,omitempty"`
	Sensors         []Sensors       `json:"sensors,omitempty"`
	Model           string          `json:"model,omitempty"`
	Grouptag        []any           `json:"grouptag,omitempty"`
	LastMeasurement time.Time       `json:"lastMeasurementAt,omitempty"`
	Location        []Location      `json:"loc,omitempty"`
}
type CurrentLocation struct {
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}
type LastMeasurement struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Value     string    `json:"value,omitempty"`
}
type Sensors struct {
	Title           string          `json:"title,omitempty"`
	Unit            string          `json:"unit,omitempty"`
	SensorType      string          `json:"sensorType,omitempty"`
	Icon            string          `json:"icon,omitempty"`
	ID              string          `json:"_id,omitempty"`
	LastMeasurement LastMeasurement `json:"lastMeasurement,omitempty"`
}
type Geometry struct {
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}
type Location struct {
	Geometry Geometry `json:"geometry,omitempty"`
	Type     string   `json:"type,omitempty"`
}

func (sbs *SenseboxService) GetSenseboxData(id string) (SenseboxJSON, error) {
	// Create structs to decode JSON data
	var data SenseboxJSON
	var apiError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	// Create URL and Client
	url := fmt.Sprintf("https://api.opensensemap.org/boxes/%s?format=json", id)

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Request data from the API
	res, err := client.Do(req)
	if err != nil {
		return data, err
	}

	// Decode error message if response status is not OK
	if res.StatusCode != http.StatusOK {
		if err = json.NewDecoder(res.Body).Decode(&apiError); err != nil {
			return data, err
		}
		return data, fmt.Errorf("api error: %s", apiError.Message)
	}

	// Decode SenseboxJSON data
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
