package db

import (
	"database/sql"
	"fmt"

	"github.com/Chandrahas77/logistics-microservices/shipment-service/pkg/models"
	_ "github.com/lib/pq"
)

type PostgresShipmentStore struct {
	db *sql.DB
}

func (p *PostgresShipmentStore) DB() *sql.DB {
	return p.db
}


func NewPostgresShipmentStore(connStr string) (*PostgresShipmentStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Postgres: %v", err)
	}
	return &PostgresShipmentStore{db: db}, nil
}

// 🛠 ADD THIS FUNCTION
func (p *PostgresShipmentStore) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresShipmentStore) CreateShipment(shipment *models.Shipment) error {
	query := `
		INSERT INTO shipments (shipment_id, order_id, status)
		VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(query, shipment.ShipmentID, shipment.OrderID, shipment.Status)
	return err
}

func (p *PostgresShipmentStore) GetShipment(shipmentID string) (*models.Shipment, error) {
	query := `
		SELECT shipment_id, order_id, status
		FROM shipments
		WHERE shipment_id = $1
	`
	row := p.db.QueryRow(query, shipmentID)

	var shipment models.Shipment
	err := row.Scan(&shipment.ShipmentID, &shipment.OrderID, &shipment.Status)
	if err != nil {
		return nil, err
	}
	return &shipment, nil
}
