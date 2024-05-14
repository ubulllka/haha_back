package service

import (
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
	"time"
)

type ResumeService struct {
	repo Resume
	logg *logger.Logger
}

func NewResumeService(repo Resume, logg *logger.Logger) *ResumeService {
	return &ResumeService{repo: repo, logg: logg}
}

func (s *ResumeService) GetAllResumes() ([]models.Resume, error) {
	return s.repo.GetAllR()
}

func (s *ResumeService) SearchResumesAnon(page int64, q string) ([]models.Resume, models.PaginationData, error) {
	return s.repo.SearchAnon(page, q)
}

func (s *ResumeService) GetResumeAnon(resumeId uint) (models.Resume, error) {
	return s.repo.GetOneAnon(resumeId)
}

func (s *ResumeService) SearchResumes(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error) {
	return s.repo.Search(userId, page, q)
}

func (s *ResumeService) GetResume(userId, resumeId uint) (DTO.ResumeDTO, error) {
	return s.repo.GetOne(userId, resumeId)
}

func (s *ResumeService) GetApplAllResumesPag(id uint, page int64) ([]models.Resume, models.PaginationData, error) {
	return s.repo.GetApplAllPag(id, page)
}

func (s *ResumeService) GetApplAllResumes(id uint) ([]DTO.ItemList, error) {
	return s.repo.GetApplAll(id)
}

func (s *ResumeService) CreateResume(userId uint, resume DTO.ResumeCreate) error {

	newResume := models.Resume{
		Post:        resume.Post,
		Description: resume.Description,
		ApplicantID: userId,
	}

	for _, v := range resume.OldWork {
		tStart, err := time.Parse(models.PARSEDATE, v.StartTime)
		if err != nil {
			s.logg.Error(err)
			return err
		}

		var tEnd time.Time
		if !strings.EqualFold(v.EndTime, "") {
			tEnd, err = time.Parse(models.PARSEDATE, v.EndTime)
			if err != nil {
				s.logg.Error(err)
				return err
			}
		} else {
			tEnd = time.Time{}
		}

		newWork := models.Work{
			Post:        v.Post,
			Description: v.Description,
			StartTime:   tStart,
			EndTime:     tEnd,
		}

		newResume.OldWorks = append(newResume.OldWorks, newWork)
	}
	return s.repo.Create(newResume)
}

func (s *ResumeService) UpdateResume(userId, resumeId uint, userRole string, resume DTO.ResumeUpdate) error {
	res, err := s.GetResumeAnon(resumeId)
	if err != nil {
		s.logg.Error(err)
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		s.logg.Error(errAuth)
		return errAuth
	}
	return s.repo.Update(resumeId, resume)
}

func (s ResumeService) DeleteResume(userId, resumeId uint, userRole string) error {
	res, err := s.GetResumeAnon(resumeId)
	if err != nil {
		s.logg.Error(err)
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		s.logg.Error(errAuth)
		return errAuth
	}
	return s.repo.Delete(resumeId)
}
