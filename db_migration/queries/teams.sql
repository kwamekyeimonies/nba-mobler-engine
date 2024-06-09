-- name: CreateTeam :one
INSERT INTO teams (
    id,
    name,
    created_at
) VALUES (
             $1, $2, NOW()
         ) RETURNING *;

-- name: GetAllTeams :many
SELECT * FROM teams;

-- name: GetTeamById :one
SELECT * FROM teams WHERE id = $1;

-- name: UpdateTeam :one
UPDATE teams
SET
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;


-- name: GetTeamByName :one
SELECT * FROM teams WHERE name=$1;

-- name: CalculateAllTeamsAverage :many
SELECT
    t.name AS team_name,
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
        JOIN
    player p ON g.player_id = p.id
        JOIN
    teams t ON p.team_id = t.id
GROUP BY
    t.name;


-- name: CalculateTeamAverage :one
SELECT
    t.name AS team_name,
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
        JOIN
    player p ON g.player_id = p.id
        JOIN
    teams t ON p.team_id = t.id
WHERE
    t.id = $1
GROUP BY
    t.name;
