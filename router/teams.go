package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/controller"
)

type TeamRouter struct {
	teamController controller.TeamController
}

func NewTeamRouter(teamController controller.TeamController) TeamRouter {
	return TeamRouter{
		teamController: teamController,
	}
}

func (teamRoute *TeamRouter) TeamRoutes(route *gin.RouterGroup) {
	route.POST("/team", teamRoute.teamController.CreateNewTeam)
	route.GET("/team", teamRoute.teamController.GetAllTeams)
	route.GET("/team/:teamId", teamRoute.teamController.GetTeamById)
	route.GET("/team//average/:teamId", teamRoute.teamController.GetTeamAverage)
	route.GET("/team/average", teamRoute.teamController.GetAllTeamsAverage)
}
