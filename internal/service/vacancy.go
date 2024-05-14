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

func (s *VacancyService) GetAllVacancies() ([]models.Vacancy, error) {
	return s.repo.GetAllV()
}

func (s *VacancyService) SearchVacanciesAnon(page int64, q string) ([]models.Vacancy, models.PaginationData, error) {
	return s.repo.SearchAnon(page, q)
}

func (s *VacancyService) GetVacancyAnon(resumeId uint) (models.Vacancy, error) {
	return s.repo.GetOneAnon(resumeId)
}

func (s *VacancyService) SearchVacancies(userId uint, page int64, q string) ([]DTO.VacancyDTO, models.PaginationData, error) {
	return s.repo.Search(userId, page, q)
}

func (s *VacancyService) GetVacancy(userId, resumeId uint) (DTO.VacancyDTO, error) {
	return s.repo.GetOne(userId, resumeId)
}

func (s *VacancyService) GetEmplAllVacanciesPag(id uint, page int64) ([]models.Vacancy, models.PaginationData, error) {
	return s.repo.GetEmplAllPag(id, page)
}

func (s *VacancyService) GetEmplAllVacancies(id uint) ([]DTO.ItemList, error) {
	return s.repo.GetEmplAll(id)
}

func (s *VacancyService) CreateVacancy(userId uint, vacancy DTO.VacancyCreate) error {
	newVacancy := models.Vacancy{
		Post:        vacancy.Post,
		Description: vacancy.Description,
		EmployerID:  userId,
	}

	return s.repo.Create(newVacancy)
}

func (s *VacancyService) UpdateVacancy(userId, vacancyId uint, userRole string, vacancy DTO.VacancyUpdate) error {
	vac, err := s.GetVacancyAnon(vacancyId)
	if err != nil {
		return err
	}
	if userId != vac.EmployerID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Update(vacancyId, vacancy)
}

func (s *VacancyService) DeleteVacancy(userId, vacancyId uint, userRole string) error {
	vac, err := s.GetVacancyAnon(vacancyId)
	if err != nil {
		return err
	}
	if userId != vac.EmployerID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Delete(vacancyId)
}
