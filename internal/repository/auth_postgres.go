package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(user models.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *AuthPostgres) GetOne(email, password string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *AuthPostgres) GetOneById(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
