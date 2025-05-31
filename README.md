# Weather Microservice

A lightweight microservice built with Go Fiber that provides weather data within a distributed system.

## Features

- RESTful API for weather data
- In-memory database with sample data
- Environment-based configuration
- Custom middleware for request tracking
- Healthcheck endpoint
- Docker and Docker Compose support
- Service layer architecture
- Utility functions for common operations
- Unit tests

## API Endpoints

### Weather Endpoints

- `GET /api/v1/weather`: Get all weather data
- `GET /api/v1/weather/:location`: Get weather data for a specific location
- `POST /api/v1/weather`: Create new weather data
- `PUT /api/v1/weather/:id`: Update existing weather data
- `DELETE /api/v1/weather/:id`: Delete weather data

### Health Check

- `GET /health`: Service health status

## Configuration

The service can be configured using environment variables:

- `SERVER_ADDRESS`: Server address and port (default: `:3000`)
- `READ_TIMEOUT`: Read timeout in seconds (default: `10`)
- `WRITE_TIMEOUT`: Write timeout in seconds (default: `10`)
- `IDLE_TIMEOUT`: Idle timeout in seconds (default: `10`)

## Getting Started

1. Clone the repository
2. Set environment variables (optional)
3. Run the service:

```bash
go run main.go
```

### Using Docker

Build and run with Docker:

```bash
docker build -t weather-microservice .
docker run -p 3000:3000 weather-microservice
```

Or use Docker Compose:

```bash
docker-compose up
```

### Running Tests

Run all tests:

```bash
go test ./...
```

Run a specific test:

```bash
go test ./services -run TestWeatherService_GetAllWeather
```

## Example Usage

### Get all weather data

```bash
curl -X GET http://localhost:3000/api/v1/weather
```

### Get weather for a specific location

```bash
curl -X GET http://localhost:3000/api/v1/weather/London
```

### Create new weather data

```bash
curl -X POST http://localhost:3000/api/v1/weather \
  -H "Content-Type: application/json" \
  -d '{
    "location": "Paris",
    "temperature": 24.5,
    "humidity": 60.0,
    "wind_speed": 5.2,
    "conditions": "Clear"
  }'
```

### Update weather data

```bash
curl -X PUT http://localhost:3000/api/v1/weather/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "location": "Paris",
    "temperature": 26.0,
    "humidity": 55.0,
    "wind_speed": 7.0,
    "conditions": "Sunny"
  }'
```

### Delete weather data

```bash
curl -X DELETE http://localhost:3000/api/v1/weather/{id}
```
