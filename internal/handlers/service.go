package handlers

import (
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"haha/internal/service"
)

type Authorization interface {
	CreateUser(name, email, telegram, password, role string) (uint, error)
	GenerateToken(email, password string) (string, string, error)
	ParseToken(token string) (uint, error)
	GetUser(id uint) (models.User, error)
}

type User interface {
	GetAllUsers() ([]models.User, error)
	GetUser(id uint) (models.User, error)
	UpdateUser(id uint, user DTO.UserUpdate) error
}

type Vacancy interface {
	GetAllVacancies() ([]models.Vacancy, error)
	SearchVacanciesAnon(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetVacancyAnon(vacancyId uint) (models.Vacancy, error)
	SearchVacancies(userId uint, page int64, q string) ([]DTO.VacancyDTO, models.PaginationData, error)
	GetVacancy(userId, vacancyId uint) (DTO.VacancyDTO, error)

	GetEmplAllVacanciesPag(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAllVacancies(userId uint) ([]DTO.ItemList, error)

	CreateVacancy(userId uint, vacancy DTO.VacancyCreate) error
	UpdateVacancy(userId, vacancyId uint, userRole string, vacancy DTO.VacancyUpdate) error
	DeleteVacancy(userId, vacancyId uint, userRole string) error
}

type Resume interface {
	GetAllResumes() ([]models.Resume, error)
	SearchResumesAnon(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetResumeAnon(resumeId uint) (models.Resume, error)
	SearchResumes(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error)
	GetResume(userId, resumeId uint) (DTO.ResumeDTO, error)

	GetApplAllResumesPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error)
	GetApplAllResumes(userId uint) ([]DTO.ItemList, error)

	CreateResume(userId uint, resume DTO.ResumeCreate) error
	UpdateResume(userId, resumeId uint, userRole string, resume DTO.ResumeUpdate) error
	DeleteResume(userId, resumeId uint, userRole string) error
}

type Work interface {
	GetListWork(resumeId uint) ([]models.Work, error)
	DeleteWork(userId, workId uint, userRole string) error
}

type Respond interface {
	CreateRespond(userId uint, userRole string, respond DTO.RespondModel) error
	UpdateRespond(userId uint, userRole string, id uint, respond DTO.RespondUpdate) error

	GetMyAllResponds(userId uint, userRole string, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetOtherAllResponds(userId uint, userRole string, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)

	DeleteMyRespond(userId uint, userRole string, respondId uint) error
	DeleteOtherRespond(userId uint, userRole string, respondId uint) error
}

type Service struct {
	Authorization
	User
	Vacancy
	Resume
	Work
	Respond
}

func NewService(repos *service.Repository, logger *logger.Logger) *Service {
	return &Service{
		Authorization: service.NewAuthService(repos.Authorization, logger),
		User:          service.NewUserService(repos.User, logger),
		Vacancy:       service.NewVacancyService(repos.Vacancy, logger),
		Resume:        service.NewResumeService(repos.Resume, logger),
		Work:          service.NewWorkService(repos.Resume, repos.Work, logger),
		Respond:       service.NewRespondService(repos.Respond, repos.Vacancy, repos.Resume, logger),
	}
}
