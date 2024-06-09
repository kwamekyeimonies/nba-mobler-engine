
-- +goose Up
ALTER TABLE game ADD COLUMN player_id uuid REFERENCES player(id) NOT NULL;