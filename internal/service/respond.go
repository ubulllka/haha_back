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
