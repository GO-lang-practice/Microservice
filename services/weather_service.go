package services

import (
	"errors"
	"go-lang-api/models"
	"log"
	"time"
)

// WeatherService handles business logic for weather data
type WeatherService struct {
	// This could be extended to include external API clients, caches, etc.
}

// New creates a new WeatherService
func New() *WeatherService {
	return &WeatherService{}
}

// GetAllWeather retrieves all weather data
func (s *WeatherService) GetAllWeather() []models.WeatherData {
	log.Println("Service: Getting all weather data")
	return models.GetAllWeatherData()
}

// GetWeatherByLocation retrieves weather data for a specific location
func (s *WeatherService) GetWeatherByLocation(location string) (models.WeatherData, error) {
	log.Printf("Service: Getting weather data for location: %s", location)
	if location == "" {
		return models.WeatherData{}, errors.New("location cannot be empty")
	}
	return models.GetWeatherByLocation(location)
}

// GetWeatherByID retrieves weather data by ID
func (s *WeatherService) GetWeatherByID(id string) (models.WeatherData, error) {
	log.Printf("Service: Getting weather data for ID: %s", id)
	if id == "" {
		return models.WeatherData{}, errors.New("id cannot be empty")
	}
	return models.GetWeatherByID(id)
}

// CreateWeather adds new weather data
func (s *WeatherService) CreateWeather(data models.WeatherData) (models.WeatherData, error) {
	log.Printf("Service: Creating weather data for location: %s", data.Location)
	// Validate required fields
	if data.Location == "" {
		return models.WeatherData{}, errors.New("location is required")
	}

	// Additional business logic can be added here
	// For example, rounding temperature to one decimal place
	data.Temperature = float64(int(data.Temperature*10)) / 10

	// Set default values if needed
	if data.Conditions == "" {
		data.Conditions = "Unknown"
	}

	return models.CreateWeatherData(data), nil
}

// UpdateWeather updates existing weather data
func (s *WeatherService) UpdateWeather(id string, data models.WeatherData) (models.WeatherData, error) {
	log.Printf("Service: Updating weather data for ID: %s", id)
	if id == "" {
		return models.WeatherData{}, errors.New("id cannot be empty")
	}

	// Get existing data to validate it exists
	existing, err := models.GetWeatherByID(id)
	if err != nil {
		return models.WeatherData{}, err
	}

	// Preserve created timestamp if it exists
	if !existing.UpdatedAt.IsZero() {
		data.UpdatedAt = time.Now()
	}

	// Additional business logic can be added here

	return models.UpdateWeatherData(id, data)
}

// DeleteWeather removes weather data by ID
func (s *WeatherService) DeleteWeather(id string) error {
	log.Printf("Service: Deleting weather data for ID: %s", id)
	if id == "" {
		return errors.New("id cannot be empty")
	}

	return models.DeleteWeatherData(id)
}
