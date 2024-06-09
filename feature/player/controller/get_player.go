package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (playerController *PlayerController) GetAllPlayers(ctx *gin.Context) {

	response, err := playerController.playerRepo.GetAllPlayers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"players": response,
	})
}

func (playerController *PlayerController) GetPlayerById(ctx *gin.Context) {

	playerId := ctx.Param("playerId")

	if len(playerId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty player id",
		})
		return
	}

	playerUUID, err := uuid.Parse(playerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty player id",
		})
		return
	}

	response, err := playerController.playerRepo.GetPlayerById(ctx, playerUUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"players": response,
	})
}

func (playerController *PlayerController) GetPlayerAverage(ctx *gin.Context) {

	playerId := ctx.Param("playerId")

	if len(playerId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty player id",
		})
		return
	}

	playerUUID, err := uuid.Parse(playerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty player id",
		})
		return
	}

	response, err := playerController.playerRepo.PlayerAverageCalculator(ctx, playerUUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"players": response,
	})
}

func (playerController *PlayerController) GetAllPlayersAverage(ctx *gin.Context) {

	response, err := playerController.playerRepo.GetAllPlayersAverage(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"players": response,
	})
}
