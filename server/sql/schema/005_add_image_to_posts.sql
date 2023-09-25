-- +goose Up
ALTER TABLE posts ADD COLUMN image TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE posts DROP COLUMN image;