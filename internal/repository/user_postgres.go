package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type UserPostgres struct {
	db   *gorm.DB
	logg *logger.Logger
}

func NewUserPostgres(db *gorm.DB, logg *logger.Logger) *UserPostgres {
	return &UserPostgres{db: db, logg: logg}
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Order("updated_at desc").Preload("Vacancies").Preload("Resumes").
		Find(&users).Error; err != nil {
		r.logg.Error(err)
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetOneById(id uint) (models.User, error) {
	var user models.User

	if err := r.db.Preload("Vacancies").Preload("Resumes").First(&user, id).Error; err != nil {
		r.logg.Error(err)
		return models.User{}, err
	}

	return user, nil
}

func (r *UserPostgres) Update(id uint, input DTO.UserUpdate) error {
	args := make(map[string]interface{})

	if input.Name != nil {
		args["name"] = *input.Name
	}

	if input.Email != nil {
		args["email"] = *input.Email
	}

	if input.Telegram != nil {
		args["telegram"] = *input.Telegram
	}

	if input.Password != nil {
		args["password"] = *input.Password
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	if input.Password != nil {
		args["status"] = *input.Status
	}

	user, err := r.GetOneById(id)
	if err != nil {
		r.logg.Error(err)
		return err
	}

	return r.db.Model(&user).Updates(args).Error
}
