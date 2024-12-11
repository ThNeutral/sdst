-- +goose Up

ALTER TABLE projects ADD COLUMN owner_id UUID UNIQUE NOT NULL;

-- +goose Down

ALTER TABLE projects DROP COLUMN owner_id;