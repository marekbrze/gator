-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_fetched_at DATE;

-- +goose Down
ALTER TABLE feeds
DELETE COLUMN last_fetched_at ;
