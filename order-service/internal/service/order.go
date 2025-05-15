package service

import (
	"context"
	"log"

	"github.com/Chandrahas77/logistics-microservices/order-service/orderpb"
	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/db"
	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	orderId := ""
	if req.OrderId == "" {
		orderId = uuid.New().String()
	}

	order := &models.Order{
		OrderID:      orderId,
		ItemID:       req.ItemId,
		Quantity:     req.Quantity,
		CustomerName: req.CustomerName,
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
		ItemId:       order.ItemID,
		Quantity:     order.Quantity,
		CustomerName: order.CustomerName,
		OrderId:      order.OrderID,
	}, nil
}

func (s *OrderServer) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	// Fetch all orders from your store
	orders, err := s.store.GetAllOrders()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders: %v", err)
	}

	// Map to protobuf Order messages
	var orderList []*orderpb.Order
	for _, o := range orders {
		orderList = append(orderList, &orderpb.Order{
			Id:           o.OrderID,
			CustomerName: o.CustomerName,
			Item:         o.ItemID,
			Quantity:     o.Quantity,
		})

	}

	return &orderpb.ListOrdersResponse{
		Orders: orderList,
	}, nil
}
