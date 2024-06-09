
-- +goose Up
ALTER TABLE player ADD COLUMN team_id uuid REFERENCES teams(id) NOT NULL;