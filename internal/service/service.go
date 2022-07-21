package service

import (
	"study/internal/repos"
	"study/internal/types"
)

type UserServ interface {
	CreateUser(user types.User) (int, error)
	GetUser(id interface{}) (*types.User, error)
	DeleteUser()
}

type Service struct {
	UserServ
}

func NewService(repos *repos.Repository) *Service {
	return &Service{
		UserServ: NewAuthService(repos.RepoUser),
	}
}
