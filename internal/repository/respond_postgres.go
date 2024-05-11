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

func (r *RespondPostgres) GetMyRespondAppl(userId uint, page int64) ([]DTO.RespondVacancy, models.PaginationData, error) {
	var result []DTO.RespondVacancy
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	dbBefore := r.db.Table("res_to_vacs").
		Select("res_to_vacs.id as id, vacancies.id as vacancy_id, status, letter, post, description, resume_id, res_to_vacs.created_at as created_at, res_to_vacs.updated_at as updated_at").
		Joins("Inner join vacancies on res_to_vacs.vacancy_id=vacancies.id").Where("resume_id IN (?)", ids).
		Count(&cnt)

	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(cnt, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&result).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return result, pag, nil
}

func (r *RespondPostgres) GetMyRespondEmpl(userId uint, page int64) ([]DTO.RespondResume, models.PaginationData, error) {
	var result []DTO.RespondResume
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	dbBefore := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, resumes.id as resume_id, status, letter, post, description, vacancy_id, vac_to_res.created_at as created_at, vac_to_res.updated_at as updated_at").
		Joins("Inner join resumes on vac_to_res.resume_id=resumes.id").Where("vacancy_id IN (?)", ids).
		Count(&cnt)

	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(cnt, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&result).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return result, pag, nil
}

func (r *RespondPostgres) GetOtherRespondAppl(userId uint, page int64) ([]DTO.RespondVacancy, models.PaginationData, error) {
	var result []DTO.RespondVacancy
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	dbBefore := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, vacancies.id as vacancy_id, status, letter, post, description, resume_id, vac_to_res.created_at as created_at, vac_to_res.updated_at as updated_at").
		Joins("Inner join vacancies on vac_to_res.vacancy_id=vacancies.id").Where("resume_id IN (?)", ids).
		Count(&cnt)

	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(cnt, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&result).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return result, pag, nil
}

func (r *RespondPostgres) GetOtherRespondEmpl(userId uint, page int64) ([]DTO.RespondResume, models.PaginationData, error) {
	var result []DTO.RespondResume
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	dbBefore := r.db.Table("res_to_vacs").
		Select("res_to_vacs.id as id, resumes.id as resume_id, status, letter, post, description, vacancy_id, res_to_vacs.created_at as created_at, res_to_vacs.updated_at as updated_at").
		Joins("Inner join resumes on res_to_vacs.resume_id=resumes.id").Where("vacancy_id IN (?)", ids).
		Count(&cnt)

	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(cnt, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&result).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return result, pag, nil
}
