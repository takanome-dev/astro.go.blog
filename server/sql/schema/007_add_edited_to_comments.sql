-- +goose Up
alter table comments add column edited_at timestamp with time zone;

-- +goose Down
alter table comments drop column edited_at;