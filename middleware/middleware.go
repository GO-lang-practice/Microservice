package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestID adds a unique request ID to each request
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate a UUID for each request
		requestID := uuid.New().String()
		c.Set("X-Request-ID", requestID)
		return c.Next()
	}
}

// RequestTime logs the time it takes to process a request
func RequestTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start timer
		start := time.Now()

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Set header
		c.Set("X-Response-Time", duration.String())

		return err
	}
}
