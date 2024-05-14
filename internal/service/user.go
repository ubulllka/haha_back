package service

import (
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type UserService struct {
	repo User
	logg *logger.Logger
}

func NewUserService(repo User, logg *logger.Logger) *UserService {
	return &UserService{repo: repo, logg: logg}
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
