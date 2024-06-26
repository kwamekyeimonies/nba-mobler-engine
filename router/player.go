package router

import (
	"github.com/gin-gonic/gin"
	playerControl "github.com/kwamekyeimonies/nba-mobler-engine/feature/player/controller"
)

type PlayerRouter struct {
	playerController playerControl.PlayerController
}

func NewPlayerRouter(playerController playerControl.PlayerController) PlayerRouter {
	return PlayerRouter{
		playerController: playerController,
	}
}

func (playerRoute *PlayerRouter) PlayerRoutes(route *gin.RouterGroup) {
	route.POST("/player", playerRoute.playerController.CreateNewPlayer)
	route.GET("/player", playerRoute.playerController.GetAllPlayers)
	route.GET("/player/:playerId", playerRoute.playerController.GetPlayerById)
	route.GET("/player/average/:playerId", playerRoute.playerController.GetPlayerAverage)
	route.GET("/player/average", playerRoute.playerController.GetAllPlayersAverage)
}
