package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type VacancyPostgres struct {
	db *gorm.DB
}

func NewVacancyPostgres(db *gorm.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}

func (r *VacancyPostgres) GetAllV() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy

	if err := r.db.Order("updated_at desc").Find(&vacancies).Error; err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyPostgres) SearchAnon(page int64, q string) ([]models.Vacancy, models.PaginationData, error) {
	var vacancies []models.Vacancy

	var count int64
	query := "%" + q + "%"
	dbBefore := r.db.Model(&models.Vacancy{}).Where("post LIKE ?", query).Count(&count)
	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&vacancies).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return vacancies, pag, nil
}

func (r *VacancyPostgres) GetOneAnon(vacancyId uint) (models.Vacancy, error) {
	var vacancy models.Vacancy
	if err := r.db.First(&vacancy, vacancyId).Error; err != nil {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

func (r *VacancyPostgres) Search(userId uint, page int64, q string) ([]DTO.VacancyDTO, models.PaginationData, error) {
	var vacancies []DTO.VacancyDTO

	var ids []string
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	var count int64
	query := "%" + q + "%"

	dbBefore := r.db.Table("vacancies").
		Select("vacancies.id as ID, post, description, employer_id, vacancies.created_at as created_at, vacancies.updated_at as updated_at, status").
		Joins("left join res_to_vacs on res_to_vacs.vacancy_id=vacancies.id AND resume_id IN (?)", ids).
		Where("post LIKE ?", query).Count(&count)
	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&vacancies).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return vacancies, pag, nil
}

func (r *VacancyPostgres) GetOne(userId, vacancyId uint) (DTO.VacancyDTO, error) {
	var vacancy DTO.VacancyDTO
	var ids []string
	if err := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Pluck("id", &ids).Error; err != nil {
		return DTO.VacancyDTO{}, err
	}

	if err := r.db.Table("vacancies").
		Select("vacancies.id as ID, post, description, employer_id, vacancies.created_at as created_at, vacancies.updated_at as updated_at, status").
		Joins("left join res_to_vacs on res_to_vacs.vacancy_id=vacancies.id AND resume_id IN (?)", ids).
		Find(&vacancy, vacancyId).Error; err != nil {
		return DTO.VacancyDTO{}, err
	}
	return vacancy, nil
}

func (r *VacancyPostgres) GetEmplAllPag(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error) {
	var vacancies []models.Vacancy

	var count int64
	dbBefore := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).Count(&count)
	if err := dbBefore.Error; err != nil {
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&vacancies).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return vacancies, pag, nil
}

func (r *VacancyPostgres) GetEmplAll(userId uint) ([]DTO.ItemList, error) {
	var vacancies []DTO.ItemList

	if err := r.db.Table("vacancies").Select("id, post").Where("employer_id = ?", userId).
		Order("updated_at desc").Find(&vacancies).Error; err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (r *VacancyPostgres) Create(vacancy models.Vacancy) error {
	return r.db.Create(&vacancy).Error
}

func (r *VacancyPostgres) Update(vacancyId uint, input DTO.VacancyUpdate) error {
	args := make(map[string]interface{})

	if input.Post != nil {
		args["post"] = *input.Post
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	vacancy, err := r.GetOneAnon(vacancyId)
	if err != nil {
		return err
	}

	if err := r.db.Model(&vacancy).Updates(args).Error; err != nil {
		return err
	}

	return nil
}

func (r *VacancyPostgres) Delete(vacancyId uint) error {
	return r.db.Unscoped().Delete(&models.Vacancy{}, vacancyId).Error
}
