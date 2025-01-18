package main

import (
	"fmt"
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

	fmt.Println(data)
	assert.NotNil(t, data, nil)
	assert.NotNil(t, data["temperature"], nil)
	assert.Equal(t, float64(12.5), data["temperature"], "temperature should be 12.5")
	assert.NotNil(t, data["status"], nil)
	assert.Equal(
		t,
		string(tempStatusGood),
		data["status"],
		fmt.Sprintf("status should be %s", tempStatusGood),
	)
}

func TestGetTempStatus(t *testing.T) {
	tests := []struct {
		name     string
		temp     float64
		expected tempStatus
	}{
		{
			name:     "Too Cold",
			temp:     0,
			expected: tempStatusCold,
		},
		{
			name:     "Good",
			temp:     20.2,
			expected: tempStatusGood,
		},
		{
			name:     "Too Hot",
			temp:     50.5,
			expected: tempStatusHot,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getTempStatus(tt.temp)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		temps    []float64
		expected float64
	}{
		{
			[]float64{20, 20, 20},
			20.00,
		},
		{
			[]float64{20, 30, 40},
			30.00,
		},
		{
			[]float64{17.8, 34.9, 55.888888},
			36.2,
		},
		{
			[]float64{17.8, 34.4, 55.888888},
			36.03,
		},
		{
			[]float64{},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Average of %v", tt.temps), func(t *testing.T) {
			result := average(tt.temps)
			assert.Equal(t, tt.expected, result)
		})
	}
}
