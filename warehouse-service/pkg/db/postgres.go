package db

import (
    "database/sql"
    "log"
    "warehouse-service/pkg/models"

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
