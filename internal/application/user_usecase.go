package application

import (
	"realworld-go-gin/internal/domain/user"
)

type UserUseCase struct {
	repo user.Repository
}

func NewUserUseCase(repo user.Repository) *UserUseCase {
	return &UserUseCase{repo}
}

func (u *UserUseCase) Register(email, password string) (*user.User, error) {
	newUser, err := user.NewUser(email, password)
	if err != nil {
	return nil, err
	}

	if err := u.repo.Save(newUser); err != nil {
	return nil, err
	}

	return newUser, nil
}