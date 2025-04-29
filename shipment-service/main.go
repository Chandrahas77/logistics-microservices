package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Chandrahas77/logistics-microservices/shipment-service/constants"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/internal/service"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/pkg/db"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/shipmentpb"
	"github.com/pressly/goose/v3"

	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	lis, err := net.Listen("tcp", constants.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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

	dbConn, err := db.NewPostgresShipmentStore(connStr)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// Goose migrations
	sqlDB := dbConn.DB() // <-- call the new DB() getter here
	log.Println("🚀 Running Goose migrations...")
	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("❌ Goose migration failed: %v", err)
	}
	log.Println("✅ Goose migrations applied successfully!")
	log.Println("✅ Connected to Postgres successfully.")

	grpcServer := grpc.NewServer()
	shipmentpb.RegisterShipmentServiceServer(grpcServer, service.NewShipmentServer(dbConn))

	log.Println("🚀 Shipment Service started on port", constants.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
