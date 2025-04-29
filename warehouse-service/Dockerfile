# Build Stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install Goose CLI
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Build Go app binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o warehouse-service main.go

# --- Production Stage ---
FROM gcr.io/distroless/static-debian11

WORKDIR /app

# Copy built binary
COPY --from=builder /app/warehouse-service .

# Copy the migrations folder (❗ This was missing)
COPY --from=builder /app/migrations ./migrations

EXPOSE 50051

# Final ENTRYPOINT
ENTRYPOINT ["/app/warehouse-service"]
