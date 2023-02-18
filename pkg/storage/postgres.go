package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/xsadia/kgallery/config"
)

func newDb(ctx *config.Context) *sql.DB {
	connString := fmt.Sprintf(ctx.Env["DB_HOST"],
		ctx.Env["DB_USERNAME"],
		ctx.Env["DB_PASSWORD"],
		ctx.Env["DB_NAME"])

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatalf("[Error]: Error while connectiong to database, %v", err)
	}

	return db
}
