package service

import (
	"errors"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type RespondService struct {
	repoRespond Respond
	repoVac     Vacancy
	repoRes     Resume
}

func NewRespondService(repoRespond Respond, repoVac Vacancy, repoRes Resume) *RespondService {
	return &RespondService{repoRespond: repoRespond, repoVac: repoVac, repoRes: repoRes}
}

func (s *RespondService) CreateRespond(userRole string, respond DTO.RespondModel) error {
	if _, err := s.repoVac.GetOne(respond.VacancyId); err != nil {
		return err
	}
	if _, err := s.repoRes.GetOne(respond.ResumeId); err != nil {
		return err
	}
	switch userRole {
	case models.APPLICANT:
		return s.repoRespond.CreateResToVac(respond)
	case models.EMPLOYER:

		return s.repoRespond.CreateVacToRes(respond)
	}
	return errors.New("wrong user's role")
}

func (s *RespondService) GetMyRespondAppl(userId uint, page int64) ([]DTO.RespondVacancy, models.PaginationData, error) {
	return s.repoRespond.GetMyRespondAppl(userId, page)
}

func (s *RespondService) GetMyRespondEmpl(userId uint, page int64) ([]models.VacToRes, models.PaginationData, error) {
	return s.repoRespond.GetMyRespondEmpl(userId, page)
}

func (s *RespondService) GetOtherRespondAppl(userId uint, page int64) ([]models.VacToRes, models.PaginationData, error) {
	return s.repoRespond.GetOtherRespondAppl(userId, page)
}

func (s *RespondService) GetOtherRespondEmpl(userId uint, page int64) ([]models.ResToVac, models.PaginationData, error) {
	return s.repoRespond.GetOtherRespondEmpl(userId, page)
}
