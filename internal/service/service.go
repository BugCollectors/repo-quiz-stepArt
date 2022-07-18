package service

import (
	"net/http"
	"study/internal/repos"
)

type UserServ interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
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
