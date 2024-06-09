package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

func (gameRepo GameRepository) GetAllGameStatistics(ctx context.Context) ([]db.GameStat, error) {
	dbResponsePayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.GetAllGameStats(ctx)
	if err != nil {
		return nil, err
	}

	return dbResponsePayload, nil
}

func (gameRepo GameRepository) GetGameByGameId(ctx context.Context, requestId uuid.UUID) (*db.Game, error) {
	dbResponsePayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.GetGameById(ctx, requestId)
	if err != nil {
		return nil, err
	}

	return &dbResponsePayload, nil
}

func (gameRepo GameRepository) GetAllGamesStatsByGameId(ctx context.Context, requestId uuid.UUID) ([]db.GameStat, error) {
	dbResponsePayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.GetGameStatsByGameId(ctx, requestId)
	if err != nil {
		return nil, err
	}

	return dbResponsePayload, nil
}

func (gameRepo GameRepository) GetAllGame(ctx context.Context) ([]db.GetGamesStatsByGameGroupingRow, error) {

	gamesDBResponsePayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.GetGamesStatsByGameGrouping(ctx)
	if err != nil {
		return nil, err
	}

	return gamesDBResponsePayload, nil

}

func (gameRepo GameRepository) GetGameStatById(ctx context.Context, requestId uuid.UUID) (*db.GameStat, error) {
	responsePayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.GetGameStatsById(ctx, requestId)
	if err != nil {
		return nil, errors.New("game stats with such id doesnot exist")
	}

	return &responsePayload, nil
}
