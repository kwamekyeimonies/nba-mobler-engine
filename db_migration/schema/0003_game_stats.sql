
-- +goose Up
CREATE TABLE game_stats(
    id uuid primary key,
    game_id uuid references game(id) not null,
    points int not null,
    rebounds int not null,
    assists int not null,
    steals int not null default 0,
    blocks int not null default 0,
    fouls int not null default 0,
    turn_overs int not null default 0,
    minutes_played int not null default 0,
    created_at timestamptz not null default  now(),
    updated_at timestamptz
);