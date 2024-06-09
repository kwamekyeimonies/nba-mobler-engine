package pkg

import (
	"context"
	"github.com/google/uuid"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/feature/team/model"
)

type ITeamsRepository interface {
	CreateTeam(ctx context.Context, request *model.CreateTeamRequest) (*string, error)
	GetAllTeams(ctx context.Context) ([]db.Team, error)
	GetTeamById(ctx context.Context, teamId uuid.UUID) (*db.Team, error)
}
