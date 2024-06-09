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


-- name: GetGamesStatsByGameGrouping :many
SELECT
    game_id,
    array_agg(id) AS stat_ids,
    array_agg(points) AS points,
    array_agg(rebounds) AS rebounds,
    array_agg(assists) AS assists,
    array_agg(steals) AS steals,
    array_agg(blocks) AS blocks,
    array_agg(fouls) AS fouls,
    array_agg(turn_overs) AS turn_overs,
    array_agg(minutes_played) AS minutes_played
FROM
    game_stats
GROUP BY
    game_id;
