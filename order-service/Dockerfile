# syntax=docker/dockerfile:1

# 1. Build Stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install goose for DB migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order-service main.go

# 2. Final Stage (using distroless base image)
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/order-service .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations ./migrations

ENTRYPOINT ["/app/order-service"]
