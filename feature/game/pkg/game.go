package pkg

import (
	"context"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/model"
)

type IGameRepository interface {
	CreateGameStatistics(ctx context.Context, request *model.CreateGameRequest) (*string, error)
	UpdateGameStatistics(ctx context.Context, request *model.UpdateGameRequest) (*string, error)
	GetAllGameStatistics(ctx context.Context) ([]db.GameStat, error)
	GetGameByGameId(ctx context.Context, requestId uuid.UUID) (*db.Game, error)
	GetAllGamesStatsByGameId(ctx context.Context, requestId uuid.UUID) ([]db.GameStat, error)
	GetAllGame(ctx context.Context) ([]db.GetGamesStatsByGameGroupingRow, error)
}
