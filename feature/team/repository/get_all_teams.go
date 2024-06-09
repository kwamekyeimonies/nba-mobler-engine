package repository

import (
	"context"
	"errors"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (teamRepo TeamRepository) GetAllTeams(ctx context.Context) ([]db.Team, error) {
	dbResponsePayload, err := teamRepo.datasource.PostgresDBConnector().PgQuery.GetAllTeams(ctx)
	if err != nil {
		return nil, errors.New("no teams available")
	}

	return dbResponsePayload, nil
}
