-- +goose Up
ALTER TABLE posts DROP COLUMN deleted_at;

-- +goose Down
ALTER TABLE posts ADD COLUMN deleted_at TIMESTAMP;