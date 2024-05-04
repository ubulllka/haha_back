package service

import (
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type UserService struct {
	repo User
}

func NewUserService(repo User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(id uint, user DTO.UserUpdate) error {
	return s.repo.Update(id, user)
}

func (s *UserService) GetUser(id uint) (models.User, error) {
	return s.repo.GetOneById(id)
}
