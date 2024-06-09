
-- +goose Up
CREATE TABLE teams(
    id uuid primary key,
    name text not null,
    created_at timestamptz not null default  now(),
    updated_at timestamptz
);