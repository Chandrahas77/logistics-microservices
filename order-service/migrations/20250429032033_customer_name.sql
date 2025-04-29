-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
ADD COLUMN customer_name TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
DROP COLUMN customer_name;
-- +goose StatementEnd
