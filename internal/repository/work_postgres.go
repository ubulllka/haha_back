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
