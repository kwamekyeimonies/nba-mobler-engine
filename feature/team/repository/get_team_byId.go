package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (teamRepo TeamRepository) GetTeamById(ctx context.Context, teamId uuid.UUID) (*db.Team, error) {
	dbResponsePayload, err := teamRepo.datasource.PostgresDBConnector().PgQuery.GetTeamById(ctx, teamId)
	if err != nil {
		return nil, errors.New("no teams available")
	}

	return &dbResponsePayload, nil
}
