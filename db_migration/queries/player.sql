
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

-- name: GetAllPlayersByTeamId :many
SELECT * FROM player WHERE team_id=$1;


-- name: CalculatePlayerAverage :one
SELECT
    g.player_id,
    AVG(gs.points) AS avg_points,
    AVG(gs.rebounds) AS avg_rebounds,
    AVG(gs.assists) AS avg_assists,
    AVG(gs.steals) AS avg_steals,
    AVG(gs.blocks) AS avg_blocks,
    AVG(gs.fouls) AS avg_fouls,
    AVG(gs.turn_overs) AS avg_turn_overs,
    AVG(gs.minutes_played) AS avg_minutes_played
FROM
    game_stats gs
        JOIN
    game g ON gs.game_id = g.id
WHERE
    g.player_id = $1
GROUP BY
    g.player_id;


-- name: GetAllPlayersAverage :many
SELECT
    g.player_id,
    AVG(gs.points) AS avg_points,
    AVG(gs.rebounds) AS avg_rebounds,
    AVG(gs.assists) AS avg_assists,
    AVG(gs.steals) AS avg_steals,
    AVG(gs.blocks) AS avg_blocks,
    AVG(gs.fouls) AS avg_fouls,
    AVG(gs.turn_overs) AS avg_turn_overs,
    AVG(gs.minutes_played) AS avg_minutes_played
FROM
    game_stats gs
        JOIN
    game g ON gs.game_id = g.id

GROUP BY
    g.player_id;
