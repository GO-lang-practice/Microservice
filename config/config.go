package config

import (
	"os"
	"strconv"
	"time"
)

// Config stores all configuration of the application.
// The values are read by using environment variables.
type Config struct {
	ServerAddress string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
}

// LoadConfig reads the configuration from environment variables
func LoadConfig() *Config {
	cfg := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":3000"),
		ReadTimeout:   time.Duration(getEnvAsInt("READ_TIMEOUT", 10)) * time.Second,
		WriteTimeout:  time.Duration(getEnvAsInt("WRITE_TIMEOUT", 10)) * time.Second,
		IdleTimeout:   time.Duration(getEnvAsInt("IDLE_TIMEOUT", 10)) * time.Second,
	}
	return cfg
}

// getEnv reads an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt reads an environment variable as an integer with a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
