FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o auth-service ./cmd/main.go

EXPOSE 8081
CMD ["./auth-service"]
