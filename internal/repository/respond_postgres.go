package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
)

type RespondPostgres struct {
	db *gorm.DB
}

func NewRespondPostgres(db *gorm.DB) *RespondPostgres {
	return &RespondPostgres{db: db}
}

func (r *RespondPostgres) GetResToVac(id uint) (models.ResToVac, error) {
	var resToVac models.ResToVac
	if err := r.db.Exec("select * from res_to_vacs where id = ?", id).First(&resToVac).Error; err != nil {
		return models.ResToVac{}, err
	}
	return resToVac, nil
}

func (r *RespondPostgres) GetVacToRes(id uint) (models.VacToRes, error) {
	var vacToRes models.VacToRes
	if err := r.db.Exec("select * from vac_to_res where id = ?", id).First(&vacToRes).Error; err != nil {
		return models.VacToRes{}, err
	}
	return vacToRes, nil
}

func (r *RespondPostgres) CreateResToVac(respond DTO.RespondModel) error {
	resToVac := &models.ResToVac{
		ResumeID:  respond.MyId,
		VacancyID: respond.ModalId,
		Letter:    respond.Letter,
		Status:    models.WAIT,
	}
	return r.db.Create(&resToVac).Error
}

func (r *RespondPostgres) CreateVacToRes(respond DTO.RespondModel) error {
	vacToRes := &models.VacToRes{
		VacancyID: respond.MyId,
		ResumeID:  respond.ModalId,
		Letter:    respond.Letter,
		Status:    models.WAIT,
	}
	return r.db.Create(&vacToRes).Error
}

func (r *RespondPostgres) UpdateResToVac(id uint, respond DTO.RespondUpdate) error {
	return r.db.Model(&models.ResToVac{}).Where("id = ?", id).Update("status", respond.Status).Error
}
func (r *RespondPostgres) UpdateVacToRes(id uint, respond DTO.RespondUpdate) error {
	return r.db.Model(&models.VacToRes{}).Where("id = ?", id).Update("status", respond.Status).Error
}

func (r *RespondPostgres) DeleteResToVac(id uint) error {
	return r.db.Unscoped().Delete(&models.ResToVac{}, id).Error
}

func (r *RespondPostgres) DeleteVacToRes(id uint) error {
	return r.db.Unscoped().Delete(&models.VacToRes{}, id).Error
}

func (r *RespondPostgres) GetMyRespondAppl(id uint) (DTO.Respond, error) {
	var result DTO.Respond
	if err := r.db.Table("res_to_vacs").
		Select("res_to_vacs.id as id, vacancies.id as vacancy_id, status, letter, "+
			"vacancies.post as post, vacancies.description as description, resume_id, "+
			"resumes.post as other_post, res_to_vacs.created_at as created_at, "+
			"res_to_vacs.updated_at as updated_at").
		Joins("inner join vacancies on res_to_vacs.vacancy_id=vacancies.id").
		Joins("inner join resumes on res_to_vacs.resume_id=resumes.id").
		First(&result, id).Error; err != nil {
		return DTO.Respond{}, err
	}
	return result, nil
}

func (r *RespondPostgres) GetMyRespondEmpl(id uint) (DTO.Respond, error) {
	var result DTO.Respond
	if err := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, resumes.id as resume_id, status, letter, "+
			"resumes.post as post, resumes.description as description, vacancy_id, "+
			"vacancies.post as other_post, vac_to_res.created_at as created_at, "+
			"vac_to_res.updated_at as updated_at").
		Joins("inner join resumes on vac_to_res.resume_id=resumes.id").
		Joins("inner join vacancies on vac_to_res.vacancy_id=vacancies.id").
		First(&result, id).Error; err != nil {
		return DTO.Respond{}, err
	}
	return result, nil
}

func (r *RespondPostgres) GetOtherRespondAppl(id uint) (DTO.Respond, error) {
	var result DTO.Respond
	if err := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, vacancies.id as vacancy_id, status, letter, "+
			"vacancies.post as post,vacancies.description as description, resume_id, "+
			"resumes.post as other_post, vac_to_res.created_at as created_at, "+
			"vac_to_res.updated_at as updated_at").
		Joins("inner join vacancies on vac_to_res.vacancy_id=vacancies.id").
		Joins("inner join resumes on vac_to_res.resume_id=resumes.id").
		First(&result, id).Error; err != nil {
		return DTO.Respond{}, err
	}
	return result, nil
}

func (r *RespondPostgres) GetOtherRespondEmpl(id uint) (DTO.Respond, error) {
	var result DTO.Respond
	if err := r.db.Table("res_to_vacs").
		Select("res_to_vacs.id as id, resumes.id as resume_id, status, letter, "+
			"resumes.post as post, resumes.description as description, vacancy_id, "+
			"vacancies.post as other_post, res_to_vacs.created_at as created_at, "+
			"res_to_vacs.updated_at as updated_at").
		Joins("inner join resumes on res_to_vacs.resume_id=resumes.id").
		Joins("inner join vacancies on res_to_vacs.vacancy_id=vacancies.id").
		First(&result, id).Error; err != nil {
		return DTO.Respond{}, err
	}
	return result, nil
}

func (r *RespondPostgres) GetMyAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	var result []DTO.Respond
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	if filter == "" {
		filter = "WAIT,ACCEPT,DECLINE"
	}
	filterField := strings.Split(filter, ",")

	dbBefore := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, resumes.id as resume_id, status, letter, "+
			"resumes.post as post, resumes.description as description, vacancy_id, "+
			"vacancies.post as other_post, vac_to_res.created_at as created_at, "+
			"vac_to_res.updated_at as updated_at").
		Joins("inner join resumes on vac_to_res.resume_id=resumes.id").
		Joins("inner join vacancies on vac_to_res.vacancy_id=vacancies.id").
		Where("resume_id IN (?)", ids).Where("status IN (?)", filterField).Count(&cnt)

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

func (r *RespondPostgres) GetMyAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	var result []DTO.Respond
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	if filter == "" {
		filter = "WAIT,ACCEPT,DECLINE"
	}
	filterField := strings.Split(filter, ",")

	dbBefore := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, resumes.id as resume_id, status, letter, "+
			"resumes.post as post, resumes.description as description, vacancy_id, "+
			"vacancies.post as other_post, vac_to_res.created_at as created_at, "+
			"vac_to_res.updated_at as updated_at").
		Joins("inner join resumes on vac_to_res.resume_id=resumes.id").
		Joins("inner join vacancies on vac_to_res.vacancy_id=vacancies.id").
		Where("vacancy_id IN (?)", ids).Where("status IN (?)", filterField).Count(&cnt)

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

func (r *RespondPostgres) GetOtherAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	var result []DTO.Respond
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	if filter == "" {
		filter = "WAIT,ACCEPT,DECLINE"
	}
	filterField := strings.Split(filter, ",")

	dbBefore := r.db.Table("vac_to_res").
		Select("vac_to_res.id as id, vacancies.id as vacancy_id, status, letter, "+
			"vacancies.post as post,vacancies.description as description, resume_id, "+
			"resumes.post as other_post, vac_to_res.created_at as created_at, "+
			"vac_to_res.updated_at as updated_at").
		Joins("inner join vacancies on vac_to_res.vacancy_id=vacancies.id").
		Joins("inner join resumes on vac_to_res.resume_id=resumes.id").
		Where("resume_id IN (?)", ids).Where("status IN (?)", filterField).Count(&cnt)

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

func (r *RespondPostgres) GetOtherAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	var result []DTO.Respond
	var ids []string
	var cnt int64
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	if filter == "" {
		filter = "WAIT,ACCEPT,DECLINE"
	}
	filterField := strings.Split(filter, ",")

	dbBefore := r.db.Table("res_to_vacs").
		Select("res_to_vacs.id as id, resumes.id as resume_id, status, letter, "+
			"resumes.post as post, resumes.description as description, vacancy_id, "+
			"vacancies.post as other_post, res_to_vacs.created_at as created_at, "+
			"res_to_vacs.updated_at as updated_at").
		Joins("inner join resumes on res_to_vacs.resume_id=resumes.id").
		Joins("inner join vacancies on res_to_vacs.vacancy_id=vacancies.id").
		Where("vacancy_id IN (?)", ids).Where("status IN (?)", filterField).Count(&cnt)

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
