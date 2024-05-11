package service

import (
	"errors"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
	"time"
)

type ResumeService struct {
	repo Resume
}

func NewResumeService(repoResume Resume) *ResumeService {
	return &ResumeService{repo: repoResume}
}

func (s *ResumeService) GetAllResumes(page int64) ([]models.Resume, models.PaginationData, error) {
	return s.repo.GetAll(page)
}

func (s *ResumeService) SearchResumes(page int64, q string) ([]models.Resume, models.PaginationData, error) {
	return s.repo.Search(page, q)
}

func (s *ResumeService) GetApplAllResumes(id uint, page int64) ([]models.Resume, models.PaginationData, error) {
	return s.repo.GetApplAll(id, page)
}

func (s *ResumeService) GetResume(resumeId uint) (models.Resume, error) {
	return s.repo.GetOne(resumeId)
}

func (s *ResumeService) CreateResume(userId uint, resume DTO.ResumeCreate) (uint, error) {

	newResume := models.Resume{
		Post:        resume.Post,
		Description: resume.Description,
		ApplicantID: userId,
	}

	for _, v := range resume.OldWork {
		tStart, err := time.Parse(models.PARSEDATE, v.StartTime)
		if err != nil {
			return 0, err
		}
		var tEnd time.Time
		if !strings.EqualFold(v.EndTime, "") {
			tEnd, err = time.Parse(models.PARSEDATE, v.EndTime)
			if err != nil {
				return 0, err
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
	resumeId, err := s.repo.Create(newResume)
	if err != nil {
		return 0, err
	}
	return resumeId, nil
}

func (s *ResumeService) UpdateResume(userId, resumeId uint, userRole string, resume DTO.ResumeUpdate) error {
	res, err := s.GetResume(resumeId)
	if err != nil {
		return err
	}
	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Update(resumeId, resume)
}

func (s ResumeService) DeleteResume(userId, resumeId uint, userRole string) error {
	res, err := s.GetResume(resumeId)
	if err != nil {
		return err
	}
	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
	}
	return s.repo.Delete(resumeId)
}
