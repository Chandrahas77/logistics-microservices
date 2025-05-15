package db

import (
	"database/sql"
	"fmt"

	"github.com/Chandrahas77/logistics-microservices/order-service/pkg/models"

	_ "github.com/lib/pq"
)

type PostgresOrderStore struct {
	db *sql.DB
}

func NewPostgresOrderStore(connStr string) (*PostgresOrderStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping Postgres: %v", err)
	}
	return &PostgresOrderStore{db: db}, nil
}

func (p *PostgresOrderStore) CreateOrder(order *models.Order) error {
	query := `
		INSERT INTO orders (order_id, item_id, quantity, customer_name)
		VALUES ($1, $2, $3, $4)
	`
	_, err := p.db.Exec(query, order.OrderID, order.ItemID, order.Quantity, order.CustomerName)
	return err
}

func (p *PostgresOrderStore) GetOrder(orderID string) (*models.Order, error) {
	query := `
		SELECT order_id, item_id, quantity, customer_name
		FROM orders
		WHERE order_id = $1
	`
	row := p.db.QueryRow(query, orderID)

	var order models.Order
	err := row.Scan(&order.OrderID, &order.ItemID, &order.Quantity, &order.CustomerName)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *PostgresOrderStore) GetAllOrders() ([]*models.Order, error) {
	rows, err := s.db.Query("SELECT order_id, customer_name, item_id, quantity FROM orders")
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %v", err)
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.OrderID, &o.CustomerName, &o.ItemID, &o.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan order: %v", err)
		}
		orders = append(orders, &o)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return orders, nil
}
