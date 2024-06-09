-- name: CreateGame :one
INSERT INTO game (
    id,
    name,
    player_id,
    created_at
) VALUES (
             $1, $2,$3, NOW()
         ) RETURNING *;

-- name: GetAllGames :many
SELECT * FROM game;

-- name: GetGameById :one
SELECT * FROM game WHERE id = $1;

-- name: GetGameByPlayerId :many
SELECT * FROM game where player_id=$1;
