package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Chandrahas77/logistics-microservices/order-service/constants"
	"github.com/Chandrahas77/logistics-microservices/order-service/internal/service"
	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/db"
	_ "github.com/jackc/pgx/v5/stdlib"


	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", constants.GRPCPort)
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		constants.PostgresUser,
		constants.PostgresPassword,
		constants.PostgresHost,
		constants.PostgresPort,
		constants.PostgresDB,
		constants.PostgresSSLMode,
	)

	log.Println("🚀 Running Goose migrations...")

	dbInstance, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to open DB for migrations: %v", err)
	}
	defer dbInstance.Close()

	goose.SetDialect("postgres")
	if err := goose.Up(dbInstance, "./migrations"); err != nil {
		log.Fatalf("❌ Goose migration failed: %v", err)
	}

	log.Println("✅ Goose migrations applied successfully!")

	store, err := db.NewPostgresOrderStore(connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to Postgres: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, service.NewOrderServer(store))

	log.Println("🚀 Order Service started on port", constants.GRPCPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
