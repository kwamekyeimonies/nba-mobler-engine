
-- name: CreatePlayer :one
INSERT INTO player
(
    id,
    name,
    team,
    team_id,
    created_at
)VALUES (
         $1,$2,$3,$4,NOW()
)RETURNING *;

-- name: GetAllPlayers :many
SELECT * FROM player;

-- name: GetPlayerById :one
SELECT * FROM player WHERE id=$1;