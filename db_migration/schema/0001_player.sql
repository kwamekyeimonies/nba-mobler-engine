
-- +goose Up
CREATE TABLE player(
    id uuid primary key,
    name text not null default '',
    team text not null default '',
    created_at timestamptz not null default  now(),
    updated_at timestamptz
);