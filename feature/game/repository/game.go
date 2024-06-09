package repository

import (
	"context"
	"github.com/kwamekyeimonies/nba-mobler-engine/datasource"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/pkg"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/player/repository"
)

type GameRepository struct {
	datasource *datasource.NewPgQueryCfg
	playerRepo repository.PlayerRepository
	ctx        context.Context
}

func NewGameRepository(datasource *datasource.NewPgQueryCfg, ctx context.Context, playerRepo repository.PlayerRepository) pkg.IGameRepository {
	return &GameRepository{
		datasource: datasource,
		ctx:        ctx,
		playerRepo: playerRepo,
	}
}
