package db

import (
	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewWritePostgres(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (w *PostgresDB) Write() {
}
