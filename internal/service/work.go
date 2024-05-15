package service

import (
	"haha/internal/logger"
	"haha/internal/models"
	"strings"
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
