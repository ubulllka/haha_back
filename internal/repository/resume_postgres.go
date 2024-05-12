package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type ResumePostgres struct {
	db *gorm.DB
}

func NewResumePostgres(db *gorm.DB) *ResumePostgres {
	return &ResumePostgres{db: db}
}

func (r *ResumePostgres) GetAll() ([]models.Resume, error) {
	var resumes []models.Resume
	if err := r.db.Order("updated_at desc").Preload("OldWorks").
		Find(&resumes).Error; err != nil {
		return nil, err
	}
	return resumes, nil
}

func (r *ResumePostgres) Search(page int64, q string) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume

	var count int64
	query := "%" + q + "%"
	dbBefore := r.db.Model(&models.Resume{}).Where("post LIKE ?", query).Count(&count)
	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Preload("OldWorks").Find(&resumes).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return resumes, pag, nil
}

func (r *ResumePostgres) GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume

	var count int64
	dbBefore := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Count(&count)
	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Preload("OldWorks").Find(&resumes).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return resumes, pag, nil
}

func (r *ResumePostgres) GetApplAll(userId uint) ([]DTO.ItemList, error) {
	var resumes []DTO.ItemList

	if err := r.db.Table("resumes").Select("id, post").Where("applicant_id = ?", userId).
		Order("updated_at desc").Find(&resumes).Error; err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *ResumePostgres) GetOne(resumeId uint) (models.Resume, error) {
	var resume models.Resume
	if err := r.db.Preload("OldWorks").First(&resume, resumeId).Error; err != nil {
		return models.Resume{}, err
	}
	return resume, nil
}

func (r *ResumePostgres) Create(resume models.Resume) (uint, error) {
	if err := r.db.Create(&resume).Error; err != nil {
		return 0, err
	}
	return resume.ID, nil
}

func (r *ResumePostgres) Update(resumeId uint, input DTO.ResumeUpdate) error {
	args := make(map[string]interface{})

	if input.Post != nil {
		args["post"] = *input.Post
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	resume, err := r.GetOne(resumeId)
	if err != nil {
		return err
	}

	if err := r.db.Model(&resume).Updates(args).Error; err != nil {
		return err
	}

	return nil
}

func (r *ResumePostgres) Delete(resumeId uint) error {
	return r.db.Delete(&models.Resume{}, resumeId).Error
}
