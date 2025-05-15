package service

import (
	"context"
	"log"

	"github.com/Chandrahas77/logistics-microservices/warehouse-service/pkg/db"
	"github.com/Chandrahas77/logistics-microservices/warehouse-service/pkg/models"
	"github.com/Chandrahas77/logistics-microservices/warehouse-service/warehousepb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WarehouseServer struct {
	warehousepb.UnimplementedWarehouseServiceServer
	inventory *db.PostgresInventory
}

func NewWarehouseServer(inventory *db.PostgresInventory) *WarehouseServer {
	return &WarehouseServer{
		inventory: inventory,
	}
}

func (s *WarehouseServer) AddInventory(ctx context.Context, req *warehousepb.AddInventoryRequest) (*warehousepb.AddInventoryResponse, error) {
	log.Printf("AddInventory called with: itemId=%s, itemName=%s, quantity=%d", req.GetItemId(), req.GetItemName(), req.GetQuantity())

	itemId := ""

	if req.ItemId == "" {
		itemId = uuid.New().String()
	}

	inv := &models.Inventory{
		ItemID:   itemId,
		ItemName: req.GetItemName(),
		Quantity: req.GetQuantity(),
	}

	err := s.inventory.Add(inv)
	if err != nil {
		log.Printf("Failed to add inventory: %v", err)
		return nil, err
	}

	return &warehousepb.AddInventoryResponse{
		Status: "Item added successfully",
	}, nil
}

func (s *WarehouseServer) GetInventory(ctx context.Context, req *warehousepb.GetInventoryRequest) (*warehousepb.GetInventoryResponse, error) {
	log.Printf("GetInventory called with: itemId=%s", req.GetItemId())

	inv, err := s.inventory.Get(req.GetItemId())
	if err != nil {
		log.Printf("Failed to get inventory: %v", err)
		return nil, err
	}

	if inv == nil {
		return &warehousepb.GetInventoryResponse{
			ItemName: "Item not found",
			Quantity: 0,
		}, nil
	}

	return &warehousepb.GetInventoryResponse{
		ItemName: inv.ItemName,
		Quantity: inv.Quantity,
	}, nil
}

func (s *WarehouseServer) ListInventory(ctx context.Context, req *warehousepb.ListInventoryRequest) (*warehousepb.ListInventoryResponse, error) {
	records, err := s.inventory.GetAllInventory()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch inventory: %v", err)
	}

	var res warehousepb.ListInventoryResponse
	for _, r := range records {
		res.Inventory = append(res.Inventory, &warehousepb.InventoryItem{
			Id:       r.ItemID,
			Name:     r.ItemName,
			Quantity: r.Quantity,
		})
	}
	return &res, nil
}
