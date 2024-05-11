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

func (r *VacancyPostgres) GetAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy

	if err := r.db.Order("updated_at desc").Find(&vacancies).Error; err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyPostgres) Search(page int64, q string) ([]models.Vacancy, models.PaginationData, error) {
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

func (r *VacancyPostgres) GetEmplAll(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error) {
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

func (r *VacancyPostgres) GetOne(vacancyId uint) (models.Vacancy, error) {
	var vacancy models.Vacancy
	if err := r.db.Find(&vacancy, vacancyId).Error; err != nil {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

func (r *VacancyPostgres) Create(vacancy models.Vacancy) (uint, error) {
	if err := r.db.Create(&vacancy).Error; err != nil {
		return 0, err
	}
	return vacancy.ID, nil
}

func (r *VacancyPostgres) Update(vacancyId uint, input DTO.VacancyUpdate) error {
	args := make(map[string]interface{})

	if input.Post != nil {
		args["post"] = *input.Post
	}

	if input.Description != nil {
		args["description"] = *input.Description
	}

	vacancy, err := r.GetOne(vacancyId)
	if err != nil {
		return err
	}

	if err := r.db.Model(&vacancy).Updates(args).Error; err != nil {
		return err
	}

	return nil
}

func (r *VacancyPostgres) Delete(vacancyId uint) error {
	return r.db.Delete(&models.Vacancy{}, vacancyId).Error
}
