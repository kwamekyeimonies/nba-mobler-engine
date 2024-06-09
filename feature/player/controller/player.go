package controller

import "github.com/kwamekyeimonies/nba-mobler-engine/feature/player/repository"

type PlayerController struct {
	playerRepo repository.PlayerRepository
}

func NewPlayerController(playerRepo repository.PlayerRepository) *PlayerController {
	return &PlayerController{
		playerRepo: playerRepo,
	}
}
