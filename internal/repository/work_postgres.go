package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
)

type WorkPostgres struct {
	db   *gorm.DB
	logg *logger.Logger
}

func NewWorkPostgres(db *gorm.DB, logg *logger.Logger) *WorkPostgres {
	return &WorkPostgres{db: db, logg: logg}
}

func (r *WorkPostgres) GetList(resumeId uint) ([]models.Work, error) {
	var works []models.Work

	if err := r.db.Where("resume_id = ?", resumeId).Find(&works).Error; err != nil {
		r.logg.Error(err)
		return nil, err
	}

	return works, nil
}

func (r *WorkPostgres) GetOne(userId uint) (models.Work, error) {
	var work models.Work

	if err := r.db.First(&work, userId).Error; err != nil {
		r.logg.Error(err)
		return models.Work{}, err
	}

	return work, nil
}

func (r *WorkPostgres) Delete(workId uint) error {
	return r.db.Unscoped().Delete(&models.Work{}, workId).Error
}
