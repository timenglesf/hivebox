package sensebox

import "time"

var dummySensebox1 = SenseboxJSON{
	ID:        "dummy_id_1",
	CreatedAt: time.Date(2024, 6, 20, 18, 56, 38, 316, time.UTC),
	UpdatedAt: time.Date(2025, 1, 12, 3, 16, 10, 948, time.UTC),
	Name:      "Dummy Location 1",
	CurrentLocation: CurrentLocation{
		Timestamp:   time.Date(2024, 6, 20, 18, 56, 38, 310, time.UTC),
		Coordinates: []float64{8.345586, 49.632145, 103},
		Type:        "Point",
	},
	Exposure: "outdoor",
	Sensors: []Sensors{
		{
			Title:      "PM10",
			Unit:       "µg/m³",
			SensorType: "SDS 011",
			ID:         "dummy_sensor_id_1",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 1, 12, 3, 16, 10, 910, time.UTC),
				Value:     "12.90",
			},
		},
		{
			Title:      "PM2.5",
			Unit:       "µg/m³",
			SensorType: "SDS 011",
			ID:         "dummy_sensor_id_2",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 1, 12, 3, 16, 10, 910, time.UTC),
				Value:     "7.87",
			},
		},
		{
			Title:      "Temperatur",
			Unit:       "°C",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_3",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 1, 12, 3, 16, 10, 910, time.UTC),
				Value:     "2.79",
			},
		},
	},
	Model:           "luftdaten_sds011_bme280",
	LastMeasurement: time.Date(2025, 1, 12, 3, 16, 10, 910, time.UTC),
}

var dummySensebox2 = SenseboxJSON{
	ID:        "dummy_id_2",
	CreatedAt: time.Date(2024, 7, 15, 12, 30, 45, 123, time.UTC),
	UpdatedAt: time.Date(2025, 2, 20, 5, 45, 30, 456, time.UTC),
	Name:      "Dummy Location 2",
	CurrentLocation: CurrentLocation{
		Timestamp:   time.Date(2024, 7, 15, 12, 30, 45, 120, time.UTC),
		Coordinates: []float64{9.123456, 48.987654, 200},
		Type:        "Point",
	},
	Exposure: "indoor",
	Sensors: []Sensors{
		{
			Title:      "Temperature",
			Unit:       "°C",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_4",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 2, 20, 5, 45, 30, 450, time.UTC),
				Value:     "22.5",
			},
		},
		{
			Title:      "Humidity",
			Unit:       "%",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_5",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 2, 20, 5, 45, 30, 450, time.UTC),
				Value:     "55.0",
			},
		},
		{
			Title:      "Temperatur",
			Unit:       "°C",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_6",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 2, 20, 5, 45, 30, 450, time.UTC),
				Value:     "22.5",
			},
		},
	},
	Model:           "indoor_bme280",
	LastMeasurement: time.Date(2025, 2, 20, 5, 45, 30, 450, time.UTC),
}

var dummySensebox3 = SenseboxJSON{
	ID:        "dummy_id_3",
	CreatedAt: time.Date(2024, 8, 10, 14, 20, 50, 789, time.UTC),
	UpdatedAt: time.Date(2025, 3, 25, 7, 55, 40, 321, time.UTC),
	Name:      "Dummy Location 3",
	CurrentLocation: CurrentLocation{
		Timestamp:   time.Date(2024, 8, 10, 14, 20, 50, 780, time.UTC),
		Coordinates: []float64{10.654321, 47.123456, 150},
		Type:        "Point",
	},
	Exposure: "outdoor",
	Sensors: []Sensors{
		{
			Title:      "Pressure",
			Unit:       "Pa",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_7",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 3, 25, 7, 55, 40, 320, time.UTC),
				Value:     "101325.0",
			},
		},
		{
			Title:      "CO2",
			Unit:       "ppm",
			SensorType: "MH-Z19",
			ID:         "dummy_sensor_id_8",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 3, 25, 7, 55, 40, 320, time.UTC),
				Value:     "400.0",
			},
		},
		{
			Title:      "Temperatur",
			Unit:       "°C",
			SensorType: "BME280",
			ID:         "dummy_sensor_id_9",
			LastMeasurement: LastMeasurement{
				CreatedAt: time.Date(2025, 3, 25, 7, 55, 40, 320, time.UTC),
				Value:     "15.0",
			},
		},
	},
	Model:           "outdoor_bme280_mhz19",
	LastMeasurement: time.Date(2025, 3, 25, 7, 55, 40, 320, time.UTC),
}
