package model

import (
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
)

type CreateGameRequest struct {
	Name          string `json:"name"`
	PlayerId      string `json:"player_id"`
	GameID        string `json:"game_id"`
	Points        int32  `json:"points"`
	Rebounds      int32  `json:"rebounds"`
	Assists       int32  `json:"assists"`
	Steals        int32  `json:"steals"`
	Blocks        int32  `json:"blocks"`
	Fouls         int32  `json:"fouls"`
	TurnOvers     int32  `json:"turn_overs"`
	MinutesPlayed int32  `json:"minutes_played"`
}

type UpdateGameRequest struct {
	GameID        string `json:"game_id"`
	GameStatsId   string `json:"game_stats_id"`
	Name          string `json:"name"`
	Points        int32  `json:"points"`
	Rebounds      int32  `json:"rebounds"`
	Assists       int32  `json:"assists"`
	Steals        int32  `json:"steals"`
	Blocks        int32  `json:"blocks"`
	Fouls         int32  `json:"fouls"`
	TurnOvers     int32  `json:"turn_overs"`
	MinutesPlayed int32  `json:"minutes_played"`
	PlayerId      string `json:"player_id"`
}

type GetGameStatsRequest struct {
	GameID      string `json:"game_id"`
	GameStatsId string `json:"game_stats_id"`
}

type GetAllGamesResponse struct {
	GameName  string    `json:"game_name"`
	GameID    uuid.UUID `json:"game_id"`
	GameStats Game      `json:"game_stats"`
}

type Game struct {
	GameStatistics []db.GameStat `json:"game_statistics"`
}
