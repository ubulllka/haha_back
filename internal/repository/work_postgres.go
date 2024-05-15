package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"time"
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

func (r *WorkPostgres) Create(work models.Work) error {
	return r.db.Create(&work).Error
}

func (r *WorkPostgres) Update(workId uint, input DTO.WorkUpdate) error {
	args := make(map[string]interface{})

	if input.Post != nil {
		args["post"] = *input.Post
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	if input.StartTime != nil {
		tStart, err := time.Parse(models.PARSEDATE, *input.StartTime)
		if err != nil {
			r.logg.Error(err)
			return err
		}
		args["start_time"] = tStart
	}

	if input.EndTime != nil {
		tEnd, err := time.Parse(models.PARSEDATE, *input.EndTime)
		if err != nil {
			r.logg.Error(err)
			return err
		}
		args["end_time"] = tEnd
	}

	work, err := r.GetOne(workId)
	if err != nil {
		r.logg.Error(err)
		return err
	}

	return r.db.Model(&work).Updates(args).Error
}

func (r *WorkPostgres) Delete(workId uint) error {
	return r.db.Unscoped().Delete(&models.Work{}, workId).Error
}
