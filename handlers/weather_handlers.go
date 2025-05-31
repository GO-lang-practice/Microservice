package handlers

import (
	"go-lang-api/models"
	"go-lang-api/services"

	"github.com/gofiber/fiber/v2"
)

// weatherService is the service instance for weather operations
var weatherService = services.New()

// GetWeatherData retrieves all weather data
func GetWeatherData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(weatherService.GetAllWeather())
}

// GetWeatherByLocation retrieves weather data for a specific location
func GetWeatherByLocation(c *fiber.Ctx) error {
	location := c.Params("location")
	if location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Location parameter is required",
		})
	}

	data, err := weatherService.GetWeatherByLocation(location)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

// CreateWeatherData adds new weather data
func CreateWeatherData(c *fiber.Ctx) error {
	var data models.WeatherData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newData, err := weatherService.CreateWeather(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newData)
}

// UpdateWeatherData updates existing weather data
func UpdateWeatherData(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	var data models.WeatherData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	updatedData, err := weatherService.UpdateWeather(id, data)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedData)
}

// DeleteWeatherData removes weather data by ID
func DeleteWeatherData(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID parameter is required",
		})
	}

	if err := weatherService.DeleteWeather(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Weather data deleted successfully",
	})
}
