# Stage 1: Build stage using the official Golang image
FROM golang:1.23.3-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the application binary
RUN go build -o gin-app

# Stage 2: Run stage using a minimal Alpine image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/gin-app .

# Expose port 8080 so the app can be accessed
EXPOSE 8080

# Run the binary when the container starts
CMD ["./gin-app"]
