package repos

import "github.com/jmoiron/sqlx"

type RepoUser interface {
	CreateUser()
	DeleteUser()
}

type Repository struct {
	RepoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoUser: NewAuthPostgres(db),
	}
}
