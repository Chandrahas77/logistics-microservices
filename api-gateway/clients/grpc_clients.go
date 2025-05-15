package clients

import (
	"log"

	"google.golang.org/grpc"
	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/shipmentpb"
	"github.com/Chandrahas77/logistics-microservices/warehouse-service/warehousepb"
)

type Clients struct {
	OrderClient     orderpb.OrderServiceClient
	ShipmentClient  shipmentpb.ShipmentServiceClient
	WarehouseClient warehousepb.WarehouseServiceClient
}

func InitClients() *Clients {
	orderConn, err := grpc.Dial("order-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}

	shipmentConn, err := grpc.Dial("shipment-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ShipmentService: %v", err)
	}

	warehouseConn, err := grpc.Dial("warehouse-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to WarehouseService: %v", err)
	}

	return &Clients{
		OrderClient:     orderpb.NewOrderServiceClient(orderConn),
		ShipmentClient:  shipmentpb.NewShipmentServiceClient(shipmentConn),
		WarehouseClient: warehousepb.NewWarehouseServiceClient(warehouseConn),
	}
}
