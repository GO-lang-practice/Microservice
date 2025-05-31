# Makefile for Weather Microservice

# Variables
BINARY_NAME=weather-service
MAIN_PACKAGE=.
DOCKER_IMAGE_NAME=weather-microservice

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the microservice
run: build
	./$(BINARY_NAME)

# Run without building
dev:
	go run $(MAIN_PACKAGE)

# Run tests
test:
	go test ./...

# Clean up binaries
clean:
	go clean
	rm -f $(BINARY_NAME)

# Format code
fmt:
	go fmt ./...

# Run linting
lint:
	golint ./...

# Docker commands
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -p 3000:3000 $(DOCKER_IMAGE_NAME)

docker-compose-up:
	docker-compose up

docker-compose-down:
	docker-compose down

# Generate documentation
docs:
	godoc -http=:6060

# All targets are phony
.PHONY: build run dev test clean fmt lint docker-build docker-run docker-compose-up docker-compose-down docs
