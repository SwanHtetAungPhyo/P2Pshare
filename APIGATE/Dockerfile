# Use a base image with Go 1.24
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code into the container
COPY . .

# Build the Go application
RUN go build -o api-gateway ./main.go

# Expose the port that the app will listen on
EXPOSE 8082


CMD ["./api-gateway"]
