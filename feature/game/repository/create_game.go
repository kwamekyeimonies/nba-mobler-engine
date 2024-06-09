package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/model"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/utils"
)

func (gameRepo GameRepository) CreateGameStatistics(ctx context.Context, request *model.CreateGameRequest) (*string, error) {
	//	Stats validators
	err := utils.GameStatsValidators(request)
	if err != nil {
		return nil, err
	}

	//	Fouls validator
	err = utils.GameStatFoulValidators(request)
	if err != nil {
		return nil, err
	}

	//	Minutes validator
	err = utils.GameStatMinutesValidator(request)
	if err != nil {
		return nil, err
	}

	playerUUID, err := uuid.Parse(request.PlayerId)
	if err != nil {
		return nil, errors.New("invalid player id")
	}

	//	Lets create game and game stats for player
	_, err = gameRepo.playerRepo.GetPlayerById(ctx, playerUUID)
	if err != nil {
		return nil, errors.New("invalid player id, player does not exist")
	}
	gameDBPayload, err := gameRepo.datasource.PostgresDBConnector().PgQuery.CreateGame(ctx, db.CreateGameParams{
		ID:       uuid.New(),
		Name:     request.Name,
		PlayerID: playerUUID,
	})
	if err != nil {
		return nil, errors.New("unable to create game")
	}

	//	create game-stats for player
	_, err = gameRepo.datasource.PostgresDBConnector().PgQuery.CreateGameStats(ctx, db.CreateGameStatsParams{
		ID:            uuid.New(),
		GameID:        gameDBPayload.ID,
		Points:        request.Points,
		Rebounds:      request.Rebounds,
		Assists:       request.Assists,
		Steals:        request.Steals,
		Blocks:        request.Blocks,
		Fouls:         request.Fouls,
		TurnOvers:     request.TurnOvers,
		MinutesPlayed: request.MinutesPlayed,
	})
	if err != nil {
		return nil, errors.New("unable to create game statistics")
	}

	responseMessage := "game statistics created successfully for player"

	return &responseMessage, nil

}
