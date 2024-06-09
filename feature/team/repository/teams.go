package repository

import (
	"context"
	"github.com/kwamekyeimonies/nba-mobler-engine/datasource"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/pkg"
)

type TeamRepository struct {
	datasource *datasource.NewPgQueryCfg
	ctx        context.Context
}

func NewTeamRepository(datasource *datasource.NewPgQueryCfg, ctx context.Context) pkg.ITeamsRepository {
	return &TeamRepository{
		datasource: datasource,
		ctx:        ctx,
	}
}
