package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Preload("Vacancies").Preload("Resumes").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) GetOneById(id uint) (models.User, error) {
	var user models.User
	//if err := r.db.Preload("Vacancies").Preload("Resumes").First(&user, id).Error; err != nil {
	//	return models.User{}, err
	//}
	if err := r.db.First(&user, id).Error; err != nil {
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
		return err
	}

	if err := r.db.Model(&user).Updates(args).Error; err != nil {
		return err
	}

	return nil
}
