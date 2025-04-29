-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shipments (
    shipment_id TEXT PRIMARY KEY,
    order_id TEXT NOT NULL,
    status TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shipments;
-- +goose StatementEnd
