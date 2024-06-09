package models

import db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"

type DBQuery struct {
	PgQuery *db.Queries
}
