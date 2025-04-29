package service

import (
	"context"
	"log"

	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/db"
	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/models"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	store *db.PostgresOrderStore
}

func NewOrderServer(store *db.PostgresOrderStore) *OrderServer {
	return &OrderServer{store: store}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	log.Printf("Creating order: %v", req)

	order := &models.Order{
		OrderID:      req.GetOrderId(),
		ItemID:       req.GetItemId(),
		Quantity:     req.GetQuantity(),
		CustomerName: req.GetCustomerName(),
	}

	err := s.store.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		Status: "Order created successfully",
	}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	log.Printf("Fetching order: %v", req.GetOrderId())

	order, err := s.store.GetOrder(req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return &orderpb.GetOrderResponse{
		ItemId:   order.ItemID,
		Quantity: order.Quantity,
		CustomerName: order.CustomerName,
		OrderId: order.OrderID,
	}, nil
}
