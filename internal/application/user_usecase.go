package application

import (
	"realworld-go-gin/internal/domain/user"
	"realworld-go-gin/internal/infrastructure/jwt"
)

type UserUseCase struct {
	service     *user.Service
	jwtProvider *jwt.Service
}

func NewUserUseCase(service *user.Service, jwtProvider *jwt.Service) *UserUseCase {
	return &UserUseCase{service, jwtProvider}
}

func (u *UserUseCase) Register(username, email, password string) (map[string]interface{}, error) {
	user, err := u.service.Register(username, email, password)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtProvider.GenerateToken(user.Email)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": map[string]interface{}{
			"email":    user.Email,
			"username": user.Username,
			"token":    token,
		},
	}, nil
}

func (u *UserUseCase) Login(email, password string) (map[string]interface{}, error) {
	user, err := u.service.SignIn(email, password)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtProvider.GenerateToken(user.Email)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": map[string]interface{}{
			"email":    user.Email,
			"username": user.Username,
			"token":    token,
		},
	}, nil
}
