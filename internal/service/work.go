package service

import (
	"errors"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"strings"
	"time"
)

type WorkService struct {
	repoRes  Resume
	repoWork Work
}

func NewWorkService(repoRes Resume, repoWork Work) *WorkService {
	return &WorkService{repoRes: repoRes, repoWork: repoWork}
}

func (s *WorkService) GetListWork(resumeId uint) ([]models.Work, error) {
	return s.repoWork.GetList(resumeId)
}

func (s *WorkService) GetWork(userId uint) (models.Work, error) {
	return s.repoWork.GetOne(userId)
}

func (s *WorkService) CreateWork(userId, resumeId uint, userRole string, work DTO.WorkCreate) (uint, error) {
	res, err := s.repoRes.GetOneAnon(resumeId)
	if err != nil {
		return 0, err
	}
	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		return 0, errors.New("not enough rights")
	}
	tStart, err := time.Parse(models.PARSEDATE, work.StartTime)
	if err != nil {
		return 0, err
	}
	var tEnd time.Time
	if !strings.EqualFold(work.EndTime, "") {
		tEnd, err = time.Parse(models.PARSEDATE, work.EndTime)
		if err != nil {
			return 0, err
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
		return err
	}

	res, err := s.repoRes.GetOneAnon(oldWork.ResumeID)
	if err != nil {
		return err
	}

	if userId != res.ApplicantID && !strings.EqualFold(userRole, models.ADMIN) {
		return errors.New("not enough rights")
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
		return errors.New("not enough rights")
	}

	return s.repoWork.Delete(workId)
}
