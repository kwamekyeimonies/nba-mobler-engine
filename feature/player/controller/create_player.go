package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/player/model"
	"net/http"
)

func (playerController *PlayerController) CreateNewPlayer(ctx *gin.Context) {
	var request *model.CreatePlayerRequest

	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty request",
		})
		return
	}

	response, err := playerController.playerRepo.CreatePlayer(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": response,
	})
}
