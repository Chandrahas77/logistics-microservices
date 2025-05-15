FROM golang:1.23-alpine

WORKDIR /app

# Install git (for Go modules)
RUN apk add --no-cache git

# Copy the entire monorepo to maintain relative imports in go.mod
COPY . .

# Go mod download
WORKDIR /app/api-gateway
RUN go mod download

# Build
RUN go build -o api-gateway main.go

# Run
CMD ["./api-gateway"]
