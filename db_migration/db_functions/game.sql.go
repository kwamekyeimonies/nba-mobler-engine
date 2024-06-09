// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: game.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createGame = `-- name: CreateGame :one
INSERT INTO game (
    id,
    name,
    created_at
) VALUES (
             $1, $2, NOW()
         ) RETURNING id, name, created_at, updated_at
`

type CreateGameParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRow(ctx, createGame, arg.ID, arg.Name)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllGames = `-- name: GetAllGames :many
SELECT id, name, created_at, updated_at FROM game
`

func (q *Queries) GetAllGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.Query(ctx, getAllGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Game{}
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getGameById = `-- name: GetGameById :one
SELECT id, name, created_at, updated_at FROM game WHERE id = $1
`

func (q *Queries) GetGameById(ctx context.Context, id uuid.UUID) (Game, error) {
	row := q.db.QueryRow(ctx, getGameById, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
