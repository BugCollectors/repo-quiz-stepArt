package service

import (
	"study/internal/repos"
	"study/internal/types"
)

type UserService struct {
	repo repos.RepoUser
}

func NewAuthService(repos repos.RepoUser) *UserService {
	return &UserService{
		repo: repos,
	}
}

func (u *UserService) CreateUser(user types.User) (int, error) {
	id, err := u.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserService) GetUser(id interface{}) (*types.User, error) {
	res, err := u.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) DeleteUser() {}
