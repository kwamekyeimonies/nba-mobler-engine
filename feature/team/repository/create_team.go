package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/model"
	"strings"
)

func (teamRepo TeamRepository) CreateTeam(ctx context.Context, request *model.CreateTeamRequest) (*string, error) {
	if len(request.TeamName) == 0 {
		return nil, errors.New("empty/invalid team name")
	}

	teamName := strings.ToLower(request.TeamName)
	_, teamCheckErr := teamRepo.datasource.PostgresDBConnector().PgQuery.GetTeamByName(ctx, teamName)
	if teamCheckErr == nil {
		return nil, errors.New("team already exist")
	}

	_, err := teamRepo.datasource.PostgresDBConnector().PgQuery.CreateTeam(ctx, db.CreateTeamParams{
		ID:   uuid.New(),
		Name: teamName,
	})

	if err != nil {
		return nil, errors.New("unable to create team")
	}

	responseMessage := "team created successfully"

	return &responseMessage, nil
}
