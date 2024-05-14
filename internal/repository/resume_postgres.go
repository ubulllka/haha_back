package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type ResumePostgres struct {
	db   *gorm.DB
	logg *logger.Logger
}

func NewResumePostgres(db *gorm.DB, logg *logger.Logger) *ResumePostgres {
	return &ResumePostgres{db: db, logg: logg}
}

func (r *ResumePostgres) GetAllR() ([]models.Resume, error) {
	var resumes []models.Resume

	if err := r.db.Order("updated_at desc").Preload("OldWorks").
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, err
	}

	return resumes, nil
}

func (r *ResumePostgres) SearchAnon(page int64, q string) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume

	var count int64
	query := "%" + q + "%"

	dbBefore := r.db.Model(&models.Resume{}).Where("post LIKE ?", query).Count(&count)
	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}
	return resumes, pag, nil
}

func (r *ResumePostgres) GetOneAnon(resumeId uint) (models.Resume, error) {
	var resume models.Resume
	if err := r.db.First(&resume, resumeId).Error; err != nil {
		r.logg.Error(err)
		return models.Resume{}, err
	}
	return resume, nil
}

func (r *ResumePostgres) Search(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error) {
	var resumes []DTO.ResumeDTO

	var ids []string
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).
		Pluck("id", &ids).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	var count int64
	query := "%" + q + "%"

	dbBefore := r.db.Table("resumes").
		Select("resumes.id as ID, post, description, applicant_id, resumes.created_at as created_at, resumes.updated_at as updated_at, status").
		Joins("left join vac_to_res on vac_to_res.resume_id=resumes.id AND vacancy_id IN (?)", ids).
		Where("post LIKE ?", query).Count(&count)

	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}
	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	return resumes, pag, nil
}

func (r *ResumePostgres) GetOne(userId, resumeId uint) (DTO.ResumeDTO, error) {
	var resume DTO.ResumeDTO
	var ids []string

	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).
		Pluck("id", &ids).Error; err != nil {
		r.logg.Error(err)
		return DTO.ResumeDTO{}, err
	}

	if err := r.db.Table("resumes").
		Select("resumes.id as ID, post, description, applicant_id, resumes.created_at as created_at, resumes.updated_at as updated_at, status").
		Joins("left join vac_to_res on vac_to_res.resume_id=resumes.id AND vacancy_id IN (?)", ids).
		Find(&resume, resumeId).Error; err != nil {
		r.logg.Error(err)
		return DTO.ResumeDTO{}, err
	}
	return resume, nil
}

func (r *ResumePostgres) GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume
	var count int64

	dbBefore := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Count(&count)
	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
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

func (r *ResumePostgres) Create(resume models.Resume) error {
	return r.db.Create(&resume).Error
}

func (r *ResumePostgres) Update(resumeId uint, input DTO.ResumeUpdate) error {
	args := make(map[string]interface{})

	if input.Post != nil {
		args["post"] = *input.Post
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	resume, err := r.GetOneAnon(resumeId)
	if err != nil {
		r.logg.Error(err)
		return err
	}

	return r.db.Model(&resume).Updates(args).Error
}

func (r *ResumePostgres) Delete(resumeId uint) error {
	return r.db.Unscoped().Delete(&models.Resume{}, resumeId).Error
}
