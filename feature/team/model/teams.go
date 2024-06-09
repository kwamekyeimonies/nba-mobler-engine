package model

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
}

type GetTeamRequest struct {
	TeamId string `json:"team_id"`
}
