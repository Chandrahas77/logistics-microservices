package service

import (
	"context"
	"log"

	"github.com/Chandrahas77/logistics-microservices/shipment-service/pkg/db"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/pkg/models"
	"github.com/Chandrahas77/logistics-microservices/shipment-service/shipmentpb"
)

type ShipmentServer struct {
	shipmentpb.UnimplementedShipmentServiceServer
	store *db.PostgresShipmentStore
}

func NewShipmentServer(store *db.PostgresShipmentStore) *ShipmentServer {
	return &ShipmentServer{store: store}
}

func (s *ShipmentServer) CreateShipment(ctx context.Context, req *shipmentpb.CreateShipmentRequest) (*shipmentpb.CreateShipmentResponse, error) {
	log.Printf("Creating shipment: %v", req)

	shipment := &models.Shipment{
		ShipmentID: req.GetShipmentId(),
		OrderID:    req.GetOrderId(),
		Status:     req.GetStatus(),
	}

	err := s.store.CreateShipment(shipment)
	if err != nil {
		return nil, err
	}

	return &shipmentpb.CreateShipmentResponse{
		Status: "Shipment created successfully",
	}, nil
}

func (s *ShipmentServer) GetShipment(ctx context.Context, req *shipmentpb.GetShipmentRequest) (*shipmentpb.GetShipmentResponse, error) {
	log.Printf("Fetching shipment: %v", req.GetShipmentId())

	shipment, err := s.store.GetShipment(req.GetShipmentId())
	if err != nil {
		return nil, err
	}

	return &shipmentpb.GetShipmentResponse{
		OrderId:    shipment.OrderID,
		Status:     shipment.Status,
		ShipmentId: shipment.ShipmentID,
	}, nil
}
