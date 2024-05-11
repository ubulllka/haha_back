package service

import (
	"errors"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
)

type VacancyService struct {
	repo Vacancy
}

func NewVacancyService(repo Vacancy) *VacancyService {
	return &VacancyService{repo: repo}
}

func (s *VacancyService) GetAllVacancies(page int64) ([]models.Vacancy, models.PaginationData, error) {
	return s.repo.GetAll(page)
}

func (s *VacancyService) SearchVacancies(page int64, q string) ([]models.Vacancy, models.PaginationData, error) {
	return s.repo.Search(page, q)
}

func (s *VacancyService) GetEmplAllVacancies(id uint, page int64) ([]models.Vacancy, models.PaginationData, error) {
	return s.repo.GetEmplAll(id, page)
}

func (s *VacancyService) GetVacancy(resumeId uint) (models.Vacancy, error) {
	return s.repo.GetOne(resumeId)
}

func (s *VacancyService) CreateVacancy(userId uint, vacancy DTO.VacancyCreate) (uint, error) {
	newVacancy := models.Vacancy{
		Post:        vacancy.Post,
		Description: vacancy.Description,
		EmployerID:  userId,
	}

	vacancyId, err := s.repo.Create(newVacancy)
	if err != nil {
		return 0, err
	}
	return vacancyId, nil
}

func (s *VacancyService) UpdateVacancy(userId, vacancyId uint, userRole string, vacancy DTO.VacancyUpdate) error {
	vac, err := s.GetVacancy(vacancyId)
	if err != nil {
		return err
	}
	if userId != vac.EmployerID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Update(vacancyId, vacancy)
}

func (s *VacancyService) DeleteVacancy(userId, vacancyId uint, userRole string) error {
	vac, err := s.GetVacancy(vacancyId)
	if err != nil {
		return err
	}
	if userId != vac.EmployerID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Delete(vacancyId)
}
