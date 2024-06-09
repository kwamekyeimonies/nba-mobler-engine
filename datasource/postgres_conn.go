package datasource

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/kwamekyeimonies/nba-mobler-engine/config"
	db "github.com/kwamekyeimonies/nba-mobler-engine/db_migration/db_functions"
	"github.com/kwamekyeimonies/nba-mobler-engine/models"
	"log"
	"time"
)

type NewPgQueryCfg struct {
	*models.DBQuery
}

func (*NewPgQueryCfg) PostgresDBConnector() *NewPgQueryCfg {
	dbUrl := config.WithConfigInjector.KeystoreConfig.DBConnectionURL

	if len(dbUrl) == 0 {
		log.Fatal("Empty Database Connection URL")
		return nil
	}

	keysConfig, err := pgx.ParseConfig(dbUrl)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	dbConn, err := pgx.ConnectConfig(ctx, keysConfig)
	if err != nil {
		log.Fatalf("Error connectiion to database, reason: %v\n", err)
		return nil
	}

	queries := db.New(dbConn)

	pgApiCfg := &NewPgQueryCfg{
		&models.DBQuery{
			PgQuery: queries,
		},
	}

	return pgApiCfg
}
