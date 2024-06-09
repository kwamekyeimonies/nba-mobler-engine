-- name: CreateGameStats :one
INSERT INTO game_stats (
    id,
    game_id,
    points,
    rebounds,
    assists,
    steals,
    blocks,
    fouls,
    turn_overs,
    minutes_played,
    created_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW()
         ) RETURNING *;

-- name: GetAllGameStats :many
SELECT * FROM game_stats;

-- name: GetGameStatsById :one
SELECT * FROM game_stats WHERE id = $1;

-- name: GetGameStatsByGameId :many
SELECT * FROM game_stats WHERE game_id = $1;

-- name: UpdateGameStats :one
UPDATE game_stats
SET
    points = $2,
    rebounds = $3,
    assists = $4,
    steals = $5,
    blocks = $6,
    fouls = $7,
    turn_overs = $8,
    minutes_played = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
