package models

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// DB is a simple in-memory database for weather data
var (
	weatherDB   = make(map[string]WeatherData)
	weatherMux  sync.RWMutex
	initialized bool
)

// WeatherData represents weather information for a location
type WeatherData struct {
	ID          string    `json:"id"`
	Location    string    `json:"location"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	Conditions  string    `json:"conditions"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// InitDB initializes the database with some sample data
func InitDB() error {
	if initialized {
		return nil
	}

	// Sample data
	sampleData := []WeatherData{
		{
			ID:          uuid.New().String(),
			Location:    "New York",
			Temperature: 22.5,
			Humidity:    65.0,
			WindSpeed:   10.2,
			Conditions:  "Partly Cloudy",
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Location:    "London",
			Temperature: 18.0,
			Humidity:    70.0,
			WindSpeed:   8.5,
			Conditions:  "Rainy",
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Location:    "Tokyo",
			Temperature: 28.5,
			Humidity:    55.0,
			WindSpeed:   6.0,
			Conditions:  "Sunny",
			UpdatedAt:   time.Now(),
		},
	}

	// Add sample data to the DB
	weatherMux.Lock()
	defer weatherMux.Unlock()

	for _, data := range sampleData {
		weatherDB[data.ID] = data
	}

	initialized = true
	return nil
}

// GetAllWeatherData retrieves all weather data
func GetAllWeatherData() []WeatherData {
	weatherMux.RLock()
	defer weatherMux.RUnlock()

	data := make([]WeatherData, 0, len(weatherDB))
	for _, w := range weatherDB {
		data = append(data, w)
	}
	return data
}

// GetWeatherByLocation retrieves weather data for a specific location
func GetWeatherByLocation(location string) (WeatherData, error) {
	weatherMux.RLock()
	defer weatherMux.RUnlock()

	for _, data := range weatherDB {
		if data.Location == location {
			return data, nil
		}
	}
	return WeatherData{}, fmt.Errorf("no weather data found for location: %s", location)
}

// GetWeatherByID retrieves weather data by ID
func GetWeatherByID(id string) (WeatherData, error) {
	weatherMux.RLock()
	defer weatherMux.RUnlock()

	data, exists := weatherDB[id]
	if !exists {
		return WeatherData{}, fmt.Errorf("weather data not found for ID: %s", id)
	}
	return data, nil
}

// CreateWeatherData adds new weather data
func CreateWeatherData(data WeatherData) WeatherData {
	data.ID = uuid.New().String()
	data.UpdatedAt = time.Now()

	weatherMux.Lock()
	defer weatherMux.Unlock()

	weatherDB[data.ID] = data
	return data
}

// UpdateWeatherData updates existing weather data
func UpdateWeatherData(id string, data WeatherData) (WeatherData, error) {
	weatherMux.Lock()
	defer weatherMux.Unlock()

	_, exists := weatherDB[id]
	if !exists {
		return WeatherData{}, fmt.Errorf("weather data not found for ID: %s", id)
	}

	// Preserve the ID
	data.ID = id
	data.UpdatedAt = time.Now()
	weatherDB[id] = data

	return data, nil
}

// DeleteWeatherData removes weather data by ID
func DeleteWeatherData(id string) error {
	weatherMux.Lock()
	defer weatherMux.Unlock()

	if _, exists := weatherDB[id]; !exists {
		return fmt.Errorf("weather data not found for ID: %s", id)
	}

	delete(weatherDB, id)
	return nil
}
