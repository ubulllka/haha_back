package service

import (
	"haha/internal/logger"
	"haha/internal/models"
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
