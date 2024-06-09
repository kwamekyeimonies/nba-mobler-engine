-- name: CreateGame :one
INSERT INTO game (
    id,
    name,
    created_at
) VALUES (
             $1, $2, NOW()
         ) RETURNING *;

-- name: GetAllGames :many
SELECT * FROM game;

-- name: GetGameById :one
SELECT * FROM game WHERE id = $1;
