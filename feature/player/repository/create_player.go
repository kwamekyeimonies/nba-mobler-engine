package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/player/model"
	"strings"
)

func (playerRepo PlayerRepository) CreatePlayer(ctx context.Context, request *model.CreatePlayerRequest) (*string, error) {
	if len(request.PlayerName) == 0 {
		return nil, errors.New("player name cannot be empty")
	}

	if len(request.TeamID) == 0 {
		return nil, errors.New("invalid/empty team id")
	}

	teamUUID, err := uuid.Parse(request.TeamID)
	if err != nil {
		return nil, errors.New("invalid team id")
	}

	teamDBPayload, err := playerRepo.teamRepo.GetTeamById(ctx, teamUUID)
	if err != nil {
		return nil, err
	}

	_, err = playerRepo.datasource.PostgresDBConnector().PgQuery.CreatePlayer(ctx, db.CreatePlayerParams{
		ID:     uuid.New(),
		Name:   strings.ToLower(request.PlayerName),
		Team:   strings.ToLower(teamDBPayload.Name),
		TeamID: teamDBPayload.ID,
	})

	if err != nil {
		return nil, errors.New("unable to create player")
	}

	responseMessage := "Player created successfully"

	return &responseMessage, nil

}
