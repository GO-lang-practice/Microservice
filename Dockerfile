FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather-service .

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install CA certificates for HTTPS requests if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/weather-service .

# Expose port
EXPOSE 3000

# Set environment variables
ENV SERVER_ADDRESS=:3000
ENV READ_TIMEOUT=10
ENV WRITE_TIMEOUT=10
ENV IDLE_TIMEOUT=10

# Command to run the executable
CMD ["./weather-service"]
