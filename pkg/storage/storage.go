package storage

import (
	"database/sql"

	"github.com/xsadia/kgallery/config"
)

type Storage struct {
	SQL *sql.DB
}

func NewStorage(ctx *config.Context) *Storage {
	return &Storage{
		SQL: newDb(ctx),
	}
}
