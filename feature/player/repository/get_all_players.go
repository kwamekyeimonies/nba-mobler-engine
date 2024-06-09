package repository

import (
	"context"
	"errors"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (playerRepo PlayerRepository) GetAllPlayers(ctx context.Context) ([]db.Player, error) {
	playersDBPayload, err := playerRepo.datasource.PostgresDBConnector().PgQuery.GetAllPlayers(ctx)
	if err != nil {
		return nil, errors.New("unable to get players")
	}

	return playersDBPayload, nil
}
