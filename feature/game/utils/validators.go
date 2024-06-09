package utils

import (
	"errors"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/model"
)

func GameStatsValidators(request *model.CreateGameRequest) error {
	if request.Points < 1 {
		return errors.New("invalid points, points must be a positive integer")
	}

	if request.Rebounds < 1 {
		return errors.New("invalid rebounds, rebounds must be a positive integer")
	}

	if request.Assists < 1 {
		return errors.New("invalid assists, assists must be a positive integer")
	}

	if request.Steals < 1 {
		return errors.New("invalid steal, steal must be a positive integer")
	}

	if request.Blocks < 1 {
		return errors.New("invalid blocks, blocks must be a positive integer")
	}

	if request.TurnOvers < 1 {
		return errors.New("invalid turnovers, turnovers must be a positive integer")
	}

	return nil

}

func GameStatFoulValidators(request *model.CreateGameRequest) error {
	if request.Fouls > 6 && request.Fouls < 0 {
		return errors.New("maximum number of fouls should not exceed 6 and should be a positive integer")
	}

	return nil
}

func GameStatMinutesValidator(request *model.CreateGameRequest) error {
	if request.MinutesPlayed > 48 && request.MinutesPlayed < 0 {
		return errors.New("minutes playes should be between 0 and 48.0")
	}

	return nil
}
