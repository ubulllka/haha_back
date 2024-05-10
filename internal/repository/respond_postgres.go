package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type RespondPostgres struct {
	db *gorm.DB
}

func NewRespondPostgres(db *gorm.DB) *RespondPostgres {
	return &RespondPostgres{db: db}
}

func (r *RespondPostgres) CreateResToVac(respond DTO.RespondModel) error {
	resToVac := &models.ResToVac{
		ResumeID:  respond.ResumeId,
		VacancyID: respond.VacancyId,
		Letter:    respond.Letter,
		Status:    models.WAIT,
	}
	return r.db.Create(&resToVac).Error
}

func (r *RespondPostgres) CreateVacToRes(respond DTO.RespondModel) error {
	vacToRes := &models.VacToRes{
		VacancyID: respond.VacancyId,
		ResumeID:  respond.ResumeId,
		Letter:    respond.Letter,
		Status:    models.WAIT,
	}
	return r.db.Create(&vacToRes).Error
}
