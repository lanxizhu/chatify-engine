package service

import (
	"chatify-engine/internal/model"
	"chatify-engine/internal/repository"
	"errors"
	"fmt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) Register(username string, password string) error {
	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to find user: %s", err))
	}
	if user != nil {
		return errors.New("user already exists")
	}

	user = &model.User{
		Username: username,
		Password: password,
	}

	if err = user.HashPassword(); err != nil {
		return errors.New(fmt.Sprintf("failed to hash password: %s", err))
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create user: %s", err))
	}

	return nil
}
