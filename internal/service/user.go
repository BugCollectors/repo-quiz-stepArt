package service

import (
	"net/http"
	"study/internal/repos"
)

type UserService struct {
	repo repos.RepoUser
}

func NewAuthService(repos repos.RepoUser) *UserService {
	return &UserService{
		repo: repos,
	}
}

func (u *UserService) CreateUser(_ http.ResponseWriter, _ *http.Request) {
	u.repo.CreateUser()
}

func (u *UserService) DeleteUser() {}
