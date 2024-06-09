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