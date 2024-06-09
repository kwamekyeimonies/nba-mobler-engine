package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (teamRepo TeamRepository) GetAllTeamsAverage(ctx context.Context) ([]db.CalculateAllTeamsAverageRow, error) {
	responsePayload, err := teamRepo.datasource.PostgresDBConnector().PgQuery.CalculateAllTeamsAverage(ctx)
	if err != nil {
		return nil, errors.New("unable to compute teams average")
	}

	return responsePayload, nil
}

func (teamRepo TeamRepository) GetTeamAverage(ctx context.Context, requestId uuid.UUID) (*db.CalculateTeamAverageRow, error) {
	responsePayload, err := teamRepo.datasource.PostgresDBConnector().PgQuery.CalculateTeamAverage(ctx, requestId)
	if err != nil {
		return nil, errors.New("unable to compute teams average")
	}

	return &responsePayload, nil
}
