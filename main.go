package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"go-lang-api/config"
	"go-lang-api/handlers"
	"go-lang-api/middleware"
	"go-lang-api/models"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database
	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestTime())

	// Routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Weather endpoints
	weather := v1.Group("/weather")
	weather.Get("/", handlers.GetWeatherData)
	weather.Get("/:location", handlers.GetWeatherByLocation)
	weather.Post("/", handlers.CreateWeatherData)
	weather.Put("/:id", handlers.UpdateWeatherData)
	weather.Delete("/:id", handlers.DeleteWeatherData)
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"service": "weather-microservice",
		})
	})

	// Start the server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := app.Listen(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
