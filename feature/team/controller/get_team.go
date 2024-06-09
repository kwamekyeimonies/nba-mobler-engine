package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (teamController *TeamController) GetAllTeams(ctx *gin.Context) {
	dbResponsePayload, err := teamController.teamRepo.GetAllTeams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"teams": dbResponsePayload,
	})
}

func (teamController *TeamController) GetAllTeamsAverage(ctx *gin.Context) {
	dbResponsePayload, err := teamController.teamRepo.GetAllTeamsAverage(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"teams": dbResponsePayload,
	})
}

func (teamController *TeamController) GetTeamById(ctx *gin.Context) {

	teamId := ctx.Param("teamId")

	teamUUID, err := uuid.Parse(teamId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team id",
		})
		return
	}

	dbResponsePayload, err := teamController.teamRepo.GetTeamById(ctx, teamUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"teams": dbResponsePayload,
	})
}

func (teamController *TeamController) GetTeamAverage(ctx *gin.Context) {

	teamId := ctx.Param("teamId")

	teamUUID, err := uuid.Parse(teamId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team id",
		})
		return
	}

	dbResponsePayload, err := teamController.teamRepo.GetTeamAverage(ctx, teamUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"teams": dbResponsePayload,
	})
}
