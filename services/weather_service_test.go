package services

import (
	"go-lang-api/models"
	"testing"
)

func TestWeatherService_GetAllWeather(t *testing.T) {
	// Initialize DB with test data
	err := models.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new service
	service := New()

	// Get all weather data
	weatherData := service.GetAllWeather()

	// Check if we got any data
	if len(weatherData) == 0 {
		t.Error("Expected some weather data, got none")
	}
}

func TestWeatherService_GetWeatherByLocation(t *testing.T) {
	// Initialize DB with test data
	err := models.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new service
	service := New()

	// Test cases
	tests := []struct {
		name      string
		location  string
		wantError bool
	}{
		{
			name:      "Valid location",
			location:  "New York",
			wantError: false,
		},
		{
			name:      "Empty location",
			location:  "",
			wantError: true,
		},
		{
			name:      "Invalid location",
			location:  "InvalidLocation",
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := service.GetWeatherByLocation(tc.location)

			// Check error
			if (err != nil) != tc.wantError {
				t.Errorf("GetWeatherByLocation() error = %v, wantError %v", err, tc.wantError)
				return
			}

			// If we expect a successful result, check the returned data
			if !tc.wantError && data.Location != tc.location {
				t.Errorf("GetWeatherByLocation() got location = %v, want %v", data.Location, tc.location)
			}
		})
	}
}

func TestWeatherService_CreateWeather(t *testing.T) {
	// Initialize DB with test data
	err := models.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new service
	service := New()

	// Test cases
	tests := []struct {
		name      string
		data      models.WeatherData
		wantError bool
	}{
		{
			name: "Valid weather data",
			data: models.WeatherData{
				Location:    "Berlin",
				Temperature: 20.5,
				Humidity:    60.0,
				WindSpeed:   5.0,
				Conditions:  "Sunny",
			},
			wantError: false,
		},
		{
			name: "Missing location",
			data: models.WeatherData{
				Temperature: 20.5,
				Humidity:    60.0,
				WindSpeed:   5.0,
				Conditions:  "Sunny",
			},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := service.CreateWeather(tc.data)

			// Check error
			if (err != nil) != tc.wantError {
				t.Errorf("CreateWeather() error = %v, wantError %v", err, tc.wantError)
				return
			}

			// If we expect a successful result, check the returned data
			if !tc.wantError {
				if data.Location != tc.data.Location {
					t.Errorf("CreateWeather() got location = %v, want %v", data.Location, tc.data.Location)
				}
				if data.ID == "" {
					t.Error("CreateWeather() didn't generate an ID")
				}
			}
		})
	}
}
