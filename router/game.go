package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/controller"
)

type GameRouter struct {
	gameController controller.GameController
}

func NewGameRouter(gameController controller.GameController) GameRouter {
	return GameRouter{
		gameController: gameController,
	}
}

func (playerRoute *GameRouter) GameRoutes(route *gin.RouterGroup) {
	route.POST("/game", playerRoute.gameController.CreateNewGame)

}
