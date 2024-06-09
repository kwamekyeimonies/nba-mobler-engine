package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (gameController *GameController) GetAllGames(ctx *gin.Context) {

	response, err := gameController.gameRepo.GetAllGame(ctx)
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

func (gameController *GameController) GetAllGamesStats(ctx *gin.Context) {

	response, err := gameController.gameRepo.GetAllGameStatistics(ctx)
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

func (gameController *GameController) GetGameStatsByGameId(ctx *gin.Context) {

	gameId := ctx.Param("gameId")
	gameUUID, err := uuid.Parse(gameId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid game id",
		})
		return
	}

	response, err := gameController.gameRepo.GetAllGamesStatsByGameId(ctx, gameUUID)
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
