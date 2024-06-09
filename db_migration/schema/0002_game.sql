
-- +goose Up
CREATE TABLE game(
    id uuid primary key,
    name text not null default '',
    created_at timestamptz not null default  now(),
    updated_at timestamptz
);