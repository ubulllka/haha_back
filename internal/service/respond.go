package service

import (
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type RespondService struct {
	repoRespond Respond
	repoVac     Vacancy
	repoRes     Resume
	logg        *logger.Logger
}

func NewRespondService(repoRespond Respond, repoVac Vacancy, repoRes Resume, logg *logger.Logger) *RespondService {
	return &RespondService{repoRespond: repoRespond, repoVac: repoVac, repoRes: repoRes, logg: logg}
}

func (s *RespondService) CreateRespond(userId uint, userRole string, respond DTO.RespondModel) error {
	switch userRole {
	case models.APPLICANT:
		if _, err := s.repoVac.GetOneAnon(respond.ModalId); err != nil {
			s.logg.Error(err)
			return err
		}

		res, err := s.repoRes.GetOneAnon(respond.MyId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != res.ApplicantID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.CreateResToVac(respond)

	case models.EMPLOYER:
		if _, err := s.repoRes.GetOneAnon(respond.ModalId); err != nil {
			s.logg.Error(err)
			return err
		}

		vac, err := s.repoVac.GetOneAnon(respond.MyId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != vac.EmployerID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.CreateVacToRes(respond)
	}
	return nil
}

func (s *RespondService) UpdateRespond(userId uint, userRole string, id uint, respond DTO.RespondUpdate) error {
	switch userRole {

	case models.APPLICANT:
		vacToRes, err := s.repoRespond.GetVacToRes(id)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		resume, err := s.repoRes.GetOneAnon(vacToRes.ResumeID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != resume.ApplicantID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.UpdateVacToRes(id, respond)

	case models.EMPLOYER:
		resToVac, err := s.repoRespond.GetResToVac(id)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		vac, err := s.repoVac.GetOneAnon(resToVac.VacancyID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != vac.EmployerID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.UpdateResToVac(id, respond)
	}
	return nil
}

func (s *RespondService) GetMyAllResponds(userId uint, userRole string, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	switch userRole {

	case models.APPLICANT:
		return s.repoRespond.GetMyAllRespondsAppl(userId, page, filter)
	case models.EMPLOYER:
		return s.repoRespond.GetMyAllRespondsEmpl(userId, page, filter)

	}
	return nil, models.PaginationData{}, nil
}

func (s *RespondService) GetOtherAllResponds(userId uint, userRole string, page int64, filter string) ([]DTO.Respond, models.PaginationData, error) {
	switch userRole {

	case models.APPLICANT:
		return s.repoRespond.GetOtherAllRespondsAppl(userId, page, filter)
	case models.EMPLOYER:
		return s.repoRespond.GetOtherAllRespondsEmpl(userId, page, filter)

	}

	return nil, models.PaginationData{}, nil
}

func (s *RespondService) DeleteMyRespond(userId uint, userRole string, respondId uint) error {
	switch userRole {

	case models.APPLICANT:
		resToVac, err := s.repoRespond.GetResToVac(respondId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		res, err := s.repoRes.GetOneAnon(resToVac.ResumeID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != res.ApplicantID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.DeleteResToVac(respondId)

	case models.EMPLOYER:

		vacToRes, err := s.repoRespond.GetVacToRes(respondId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		vac, err := s.repoVac.GetOneAnon(vacToRes.VacancyID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != vac.EmployerID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.DeleteVacToRes(respondId)
	}

	return nil
}
func (s *RespondService) DeleteOtherRespond(userId uint, userRole string, respondId uint) error {
	switch userRole {

	case models.APPLICANT:
		vacToRes, err := s.repoRespond.GetVacToRes(respondId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		vac, err := s.repoRes.GetOneAnon(vacToRes.ResumeID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != vac.ApplicantID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.DeleteVacToRes(respondId)

	case models.EMPLOYER:
		resToVac, err := s.repoRespond.GetResToVac(respondId)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		vac, err := s.repoVac.GetOneAnon(resToVac.VacancyID)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		if userId != vac.EmployerID {
			s.logg.Error(errAuth)
			return errAuth
		}

		return s.repoRespond.DeleteResToVac(respondId)
	}

	return nil
}
