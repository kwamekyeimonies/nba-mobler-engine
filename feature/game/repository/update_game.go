package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/model"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/utils"
)

func (gameRepo GameRepository) UpdateGameStatistics(ctx context.Context, request *model.UpdateGameRequest) (*string, error) {
	//	Stats validators
	updateRequestPayload := &model.CreateGameRequest{
		Name:          request.Name,
		PlayerId:      request.PlayerId,
		Points:        request.Points,
		Rebounds:      request.Rebounds,
		Assists:       request.Assists,
		Steals:        request.Steals,
		Blocks:        request.Blocks,
		Fouls:         request.Fouls,
		TurnOvers:     request.TurnOvers,
		MinutesPlayed: request.MinutesPlayed,
	}
	err := utils.GameStatsValidators(updateRequestPayload)
	if err != nil {
		return nil, err
	}

	//	Fouls validator
	err = utils.GameStatFoulValidators(updateRequestPayload)
	if err != nil {
		return nil, err
	}

	//	Minutes validator
	err = utils.GameStatMinutesValidator(updateRequestPayload)
	if err != nil {
		return nil, err
	}

	gameStatUUID, err := uuid.Parse(request.GameStatsId)
	if err != nil {
		return nil, errors.New("invalid player id")
	}

	//	Lets create game and game stats for player
	_, err = gameRepo.datasource.PostgresDBConnector().PgQuery.UpdateGameStats(ctx, db.UpdateGameStatsParams{
		ID:            gameStatUUID,
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
		return nil, errors.New("unable to update game statistics")
	}

	responseMessage := "game statistics updated successfully for player"

	return &responseMessage, nil

}
