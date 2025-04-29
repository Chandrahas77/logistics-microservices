-- +goose Up
CREATE TABLE IF NOT EXISTS inventory (
    item_id TEXT PRIMARY KEY,
    item_name TEXT NOT NULL,
    quantity INTEGER NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS inventory;
