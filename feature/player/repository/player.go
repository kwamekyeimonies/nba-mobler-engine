package repository

import (
	"context"
	"github.com/kwamekyeimonies/nba-mobler-engine/datasource"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/player/pkg"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/repository"
)

type PlayerRepository struct {
	datasource *datasource.NewPgQueryCfg
	teamRepo   repository.TeamRepository
	ctx        context.Context
}

func NewPlayerRepository(ctx context.Context, datasource *datasource.NewPgQueryCfg, teamRepo repository.TeamRepository) pkg.IPlayerRepository {

	return &PlayerRepository{
		ctx:        ctx,
		datasource: datasource,
		teamRepo:   teamRepo,
	}
}
