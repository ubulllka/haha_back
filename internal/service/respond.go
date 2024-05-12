package service

import (
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
	switch userRole {
	case models.APPLICANT:
		if _, err := s.repoRes.GetOne(respond.MyId); err != nil {
			return err
		}
		if _, err := s.repoVac.GetOne(respond.ModalId); err != nil {
			return err
		}
		return s.repoRespond.CreateResToVac(respond)
	case models.EMPLOYER:
		if _, err := s.repoVac.GetOne(respond.MyId); err != nil {
			return err
		}
		if _, err := s.repoRes.GetOne(respond.ModalId); err != nil {
			return err
		}
		return s.repoRespond.CreateVacToRes(respond)
	}
	return nil
}

func (s *RespondService) GetMyRespondAppl(userId uint, page int64) ([]DTO.RespondVacancy, models.PaginationData, error) {
	return s.repoRespond.GetMyRespondAppl(userId, page)
}

func (s *RespondService) GetMyRespondEmpl(userId uint, page int64) ([]DTO.RespondResume, models.PaginationData, error) {
	return s.repoRespond.GetMyRespondEmpl(userId, page)
}

func (s *RespondService) GetOtherRespondAppl(userId uint, page int64) ([]DTO.RespondVacancy, models.PaginationData, error) {
	return s.repoRespond.GetOtherRespondAppl(userId, page)
}

func (s *RespondService) GetOtherRespondEmpl(userId uint, page int64) ([]DTO.RespondResume, models.PaginationData, error) {
	return s.repoRespond.GetOtherRespondEmpl(userId, page)
}
