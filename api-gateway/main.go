package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/shipmentpb"
	"github.com/Chandrahas77/logistics-microservices/warehouse-service/warehousepb"
)

func main() {
	r := gin.Default()

	// Connect to Order Service
	orderConn, err := grpc.Dial("order-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// Connect to Shipment Service
	shipmentConn, err := grpc.Dial("shipment-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to shipment service: %v", err)
	}
	shipmentClient := shipmentpb.NewShipmentServiceClient(shipmentConn)

	// Connect to Warehouse Service
	warehouseConn, err := grpc.Dial("warehouse-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to warehouse service: %v", err)
	}
	warehouseClient := warehousepb.NewWarehouseServiceClient(warehouseConn)

	// Endpoint: List Orders
	r.GET("/orders", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := orderClient.ListOrders(ctx, &orderpb.ListOrdersRequest{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res.Orders)
	})

	// Endpoint: List Shipments
	r.GET("/shipments", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := shipmentClient.ListShipments(ctx, &shipmentpb.ListShipmentsRequest{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res.Shipments)
	})

	// Endpoint: List Inventory
	r.GET("/inventory", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := warehouseClient.ListInventory(ctx, &warehousepb.ListInventoryRequest{})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res.Inventory)
	})

	// Start the gateway
	r.Run(":8080")
}
