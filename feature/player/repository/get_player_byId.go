package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (playerRepo PlayerRepository) GetPlayerById(ctx context.Context, requestId uuid.UUID) (*db.Player, error) {
	playersDBPayload, err := playerRepo.datasource.PostgresDBConnector().PgQuery.GetPlayerById(ctx, requestId)
	if err != nil {
		return nil, errors.New("unable to get players")
	}

	return &playersDBPayload, nil
}
