package controller

import "github.com/kwamekyeimonies/nba-mobler-engine/feature/game/repository"

type GameController struct {
	gameRepo repository.GameRepository
}

func NewGameController(gameRepo repository.GameRepository) GameController {
	return GameController{
		gameRepo: gameRepo,
	}
}
