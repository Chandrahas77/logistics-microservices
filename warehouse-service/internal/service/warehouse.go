package service

import (
    "context"
    "log"
    "warehouse-service/pkg/db"
    "warehouse-service/pkg/models"
    "warehouse-service/warehousepb"
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

    // Build inventory model
    inv := &models.Inventory{
        ItemID:   req.GetItemId(),
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
