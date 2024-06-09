package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/model"
	"net/http"
)

func (gameController *GameController) UpdateGame(ctx *gin.Context) {
	var request *model.UpdateGameRequest

	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	response, err := gameController.gameRepo.UpdateGameStatistics(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": response,
	})
	return
}
