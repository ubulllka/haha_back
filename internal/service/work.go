package service

import (
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
	"time"
)

type WorkService struct {
	repoRes  Resume
	repoWork Work
	logg     *logger.Logger
}

func NewWorkService(repoRes Resume, repoWork Work, logg *logger.Logger) *WorkService {
	return &WorkService{repoRes: repoRes, repoWork: repoWork, logg: logg}
}

func (s *WorkService) GetListWork(resumeId uint) ([]models.Work, error) {
	return s.repoWork.GetList(resumeId)
}

func (s *WorkService) GetWork(userId uint) (models.Work, error) {
	return s.repoWork.GetOne(userId)
}

func (s *WorkService) CreateWork(userId, resumeId uint, userRole string, work DTO.WorkCreate) error {
	res, err := s.repoRes.GetOneAnon(resumeId)
	if err != nil {
		s.logg.Error(err)
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		s.logg.Error(errAuth)
		return errAuth
	}
	tStart, err := time.Parse(models.PARSEDATE, work.StartTime)
	if err != nil {
		s.logg.Error(err)
		return err
	}
	var tEnd time.Time
	if !strings.EqualFold(work.EndTime, "") {
		tEnd, err = time.Parse(models.PARSEDATE, work.EndTime)
		if err != nil {
			s.logg.Error(err)
			return err
		}
	} else {
		tEnd = time.Time{}
	}

	newWork := models.Work{
		Post:        work.Post,
		Description: work.Description,
		StartTime:   tStart,
		EndTime:     tEnd,
		ResumeID:    resumeId,
	}
	return s.repoWork.Create(newWork)
}

func (s *WorkService) UpdateWork(userId, workId uint, userRole string, work DTO.WorkUpdate) error {
	oldWork, err := s.GetWork(workId)
	if err != nil {
		s.logg.Error(err)
		return err
	}

	res, err := s.repoRes.GetOneAnon(oldWork.ResumeID)
	if err != nil {
		s.logg.Error(err)
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		s.logg.Error(errAuth)
		return errAuth
	}

	return s.repoWork.Update(workId, work)
}

func (s *WorkService) DeleteWork(userId, workId uint, userRole string) error {
	oldWork, err := s.GetWork(workId)
	if err != nil {
		return err
	}

	res, err := s.repoRes.GetOneAnon(oldWork.ResumeID)
	if err != nil {
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		s.logg.Error(errAuth)
		return errAuth
	}

	return s.repoWork.Delete(workId)
}
