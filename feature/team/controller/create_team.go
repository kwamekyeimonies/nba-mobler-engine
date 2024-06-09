package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/model"
	"net/http"
)

func (teamController *TeamController) CreateNewTeam(ctx *gin.Context) {
	var request *model.CreateTeamRequest

	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty body request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty body request",
		})
		return
	}

	response, err := teamController.teamRepo.CreateTeam(ctx, request)
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
