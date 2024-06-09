package model

type CreatePlayerRequest struct {
	PlayerName string `json:"player_name"`
	TeamID     string `json:"team_id"`
}

type GetPlayerByRequest struct {
	PlayerName string `json:"player_name"`
	PlayerId   string `json:"player_id"`
}
