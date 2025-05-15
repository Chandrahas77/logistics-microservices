package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Chandrahas77/logistics-microservices/warehouse-service/pkg/models"

	_ "github.com/lib/pq"
)

type PostgresInventory struct {
	db *sql.DB
}

func NewPostgresInventory(connStr string) (*PostgresInventory, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to Postgres successfully.")
	return &PostgresInventory{db: db}, nil
}

func (p *PostgresInventory) Add(inv *models.Inventory) error {
	_, err := p.db.Exec(`
        INSERT INTO inventory (item_id, item_name, quantity)
        VALUES ($1, $2, $3)
        ON CONFLICT (item_id)
        DO UPDATE SET item_name = EXCLUDED.item_name, quantity = EXCLUDED.quantity
    `, inv.ItemID, inv.ItemName, inv.Quantity)
	return err
}

func (p *PostgresInventory) Get(itemID string) (*models.Inventory, error) {
	row := p.db.QueryRow(`SELECT item_id, item_name, quantity FROM inventory WHERE item_id = $1`, itemID)

	var inv models.Inventory
	if err := row.Scan(&inv.ItemID, &inv.ItemName, &inv.Quantity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &inv, nil
}

func (s *PostgresInventory) GetAllInventory() ([]*models.Inventory, error) {
	rows, err := s.db.Query("SELECT item_id, item_name, quantity FROM inventory")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inventory: %v", err)
	}
	defer rows.Close()

	var inventoryList []*models.Inventory
	for rows.Next() {
		var inv models.Inventory
		if err := rows.Scan(&inv.ItemID, &inv.ItemName, &inv.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		inventoryList = append(inventoryList, &inv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return inventoryList, nil
}
