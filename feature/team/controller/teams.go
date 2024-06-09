package controller

import "github.com/kwamekyeimonies/nba-mobler-engine/feature/team/repository"

type TeamController struct {
	teamRepo repository.TeamRepository
}

func NewTeamController(teamRepo repository.TeamRepository) *TeamController {
	return &TeamController{
		teamRepo: teamRepo,
	}
}
