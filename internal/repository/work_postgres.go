package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"time"
)

type WorkPostgres struct {
	db *gorm.DB
}

func NewWorkPostgres(db *gorm.DB) *WorkPostgres {
	return &WorkPostgres{db: db}
}

func (r *WorkPostgres) GetList(resumeId uint) ([]models.Work, error) {
	var works []models.Work
	if err := r.db.Where("resume_id = ?", resumeId).Find(&works).Error; err != nil {
		return nil, err
	}
	return works, nil
}

func (r *WorkPostgres) GetOne(userId uint) (models.Work, error) {
	var work models.Work
	if err := r.db.First(&work, userId).Error; err != nil {
		return models.Work{}, err
	}
	return work, nil
}

func (r *WorkPostgres) Create(work models.Work) (uint, error) {
	if err := r.db.Create(&work).Error; err != nil {
		return 0, err
	}
	return work.ID, nil
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
			return err
		}
		args["start_time"] = tStart
	}

	if input.EndTime != nil {
		tEnd, err := time.Parse(models.PARSEDATE, *input.EndTime)
		if err != nil {
			return err
		}
		args["end_time"] = tEnd
	}

	work, err := r.GetOne(workId)
	if err != nil {
		return err
	}

	if err := r.db.Model(&work).Updates(args).Error; err != nil {
		return err
	}

	return nil
}

func (r *WorkPostgres) Delete(workId uint) error {
	return r.db.Delete(&models.Work{}, workId).Error
}
