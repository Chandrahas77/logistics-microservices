#!/bin/sh

# Run Goose migrations
echo "Running Goose migrations..."
goose -dir ./migrations postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}" up

# Start warehouse-service
echo "Starting warehouse service..."
./warehouse-service
