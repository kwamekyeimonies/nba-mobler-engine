package repository

import (
	"context"
	"github.com/kwamekyeimonies/nba-mobler-engine/datasource"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/game/pkg"
)

type GameRepository struct {
	datasource *datasource.NewPgQueryCfg
	ctx        context.Context
}

func NewGameRepository(datasource *datasource.NewPgQueryCfg, ctx context.Context) pkg.IGameRepository {
	return &GameRepository{
		datasource: datasource,
		ctx:        ctx,
	}
}
