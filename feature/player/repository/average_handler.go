package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (playerRepo PlayerRepository) GetAllPlayersAverage(ctx context.Context) ([]db.GetAllPlayersAverageRow, error) {
	allPlayersAverage, err := playerRepo.datasource.PostgresDBConnector().PgQuery.GetAllPlayersAverage(ctx)
	if err != nil {
		return nil, errors.New("unable to compute player average")
	}

	return allPlayersAverage, nil
}

func (playerRepo PlayerRepository) PlayerAverageCalculator(ctx context.Context, requestId uuid.UUID) (*db.CalculatePlayerAverageRow, error) {
	averageResult, err := playerRepo.datasource.PostgresDBConnector().PgQuery.CalculatePlayerAverage(ctx, requestId)
	if err != nil {
		return nil, errors.New("unable to compute player average")
	}

	return &averageResult, nil
}
