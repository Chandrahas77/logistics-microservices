# syntax=docker/dockerfile:1

FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# COPY --from=golang:1.23 /go/bin/goose /usr/local/bin/goose


RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shipment-service main.go

FROM gcr.io/distroless/base-debian11:latest

WORKDIR /app

COPY --from=builder /app/shipment-service .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations ./migrations

ENV POSTGRES_USER=shipmentuser
ENV POSTGRES_PASSWORD=shipmentpass
ENV POSTGRES_DB=shipmentdb
ENV POSTGRES_HOST=shipment-postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_SSLMODE=disable

ENTRYPOINT ["./shipment-service"]
