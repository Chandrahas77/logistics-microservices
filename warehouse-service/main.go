package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"warehouse-service/constants"
	"warehouse-service/internal/service"
	"warehouse-service/pkg/db"
	"warehouse-service/warehousepb"

	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to Postgres
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		constants.PostgresUser,
		constants.PostgresPassword,
		constants.PostgresHost,
		constants.PostgresPort,
		constants.PostgresDB,
		constants.PostgresSSLMode,
	)

	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ failed to connect to db: %v", err)
	}

	// 🚀 Run Goose migrations programmatically
	log.Println("🚀 Running Goose migrations...")
	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("❌ Goose migration failed: %v", err)
	}
	log.Println("✅ Goose migrations applied successfully!")

	// Set up inventory database connection
	inventoryDB, err := db.NewPostgresInventory(connStr)
	if err != nil {
		log.Fatalf("❌ failed to set up inventory db: %v", err)
	}

	// Start gRPC server
	lis, err := net.Listen("tcp", constants.GRPCPort)
	if err != nil {
		log.Fatalf("❌ failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	warehousepb.RegisterWarehouseServiceServer(grpcServer, service.NewWarehouseServer(inventoryDB))

	log.Println("🚀 Warehouse Service started on port", constants.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}
