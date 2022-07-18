package repos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	tableName = "users"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (a *UserRepo) CreateUser() {
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1)", tableName)

	row := a.db.QueryRow(query, "TOXA")

	fmt.Println(row)
}

func (a *UserRepo) DeleteUser() {}
