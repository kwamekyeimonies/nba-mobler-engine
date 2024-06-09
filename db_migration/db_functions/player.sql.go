// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: player.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO player
(
    id,
    name,
    team,
    team_id,
    created_at
)VALUES (
         $1,$2,$3,$4,NOW()
)RETURNING id, name, team, created_at, updated_at, team_id
`

type CreatePlayerParams struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Team   string    `json:"team"`
	TeamID uuid.UUID `json:"team_id"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, createPlayer,
		arg.ID,
		arg.Name,
		arg.Team,
		arg.TeamID,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Team,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TeamID,
	)
	return i, err
}

const getAllPlayers = `-- name: GetAllPlayers :many
SELECT id, name, team, created_at, updated_at, team_id FROM player
`

func (q *Queries) GetAllPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.Query(ctx, getAllPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Player{}
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Team,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TeamID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPlayersByTeamId = `-- name: GetAllPlayersByTeamId :many
SELECT id, name, team, created_at, updated_at, team_id FROM player WHERE team_id=$1
`

func (q *Queries) GetAllPlayersByTeamId(ctx context.Context, teamID uuid.UUID) ([]Player, error) {
	rows, err := q.db.Query(ctx, getAllPlayersByTeamId, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Player{}
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Team,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TeamID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayerById = `-- name: GetPlayerById :one
SELECT id, name, team, created_at, updated_at, team_id FROM player WHERE id=$1
`

func (q *Queries) GetPlayerById(ctx context.Context, id uuid.UUID) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayerById, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Team,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TeamID,
	)
	return i, err
}
