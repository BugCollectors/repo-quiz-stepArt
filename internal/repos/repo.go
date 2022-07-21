package repos

import (
	"github.com/jmoiron/sqlx"
	"study/internal/types"
)

type RepoUser interface {
	CreateUser(user types.User) (int, error)
	GetUser(id interface{}) (*types.User, error)
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
