package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/xsadia/kgallery/config"
)

func newDb(ctx *config.Context) *sql.DB {
	url := `host=%s user=%s password=%s dbname=%s sslmode=disable`
	connString := fmt.Sprintf(url, ctx.Env["PGSQL_HOST"],
		ctx.Env["PGSQL_USER"],
		ctx.Env["PGSQL_PASSWORD"],
		ctx.Env["PGSQL_DBNAME"])

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatalf("[Error]: Error while connectiong to database, %v", err)
	}

	return db
}
