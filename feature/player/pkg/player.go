package pkg

import (
	"context"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/player/model"
)

type IPlayerRepository interface {
	CreatePlayer(ctx context.Context, request *model.CreatePlayerRequest) (*string, error)
	GetPlayerById(ctx context.Context, requestId uuid.UUID) (*db.Player, error)
	GetAllPlayers(ctx context.Context) ([]db.Player, error)
	PlayerAverageCalculator(ctx context.Context, requestId uuid.UUID) (*db.CalculatePlayerAverageRow, error)
	GetAllPlayersAverage(ctx context.Context) ([]db.GetAllPlayersAverageRow, error)
}
