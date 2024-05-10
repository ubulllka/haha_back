package handlers

import (
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
	GetAllVacancies(page int64) ([]models.Vacancy, models.PaginationData, error)
	SearchVacancies(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAllVacancies(userId uint) ([]models.Vacancy, error)
	GetVacancy(vacancyId uint) (models.Vacancy, error)
	CreateVacancy(userId uint, vacancy DTO.VacancyCreate) (uint, error)
	UpdateVacancy(userId, vacancyId uint, userRole string, vacancy DTO.VacancyUpdate) error
	DeleteVacancy(userId, vacancyId uint, userRole string) error
}

type Resume interface {
	GetAllResumes(page int64) ([]models.Resume, models.PaginationData, error)
	SearchResumes(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetApplAllResumes(userId uint) ([]models.Resume, error)
	GetResume(resumeId uint) (models.Resume, error)
	CreateResume(userId uint, resume DTO.ResumeCreate) (uint, error)
	UpdateResume(userId, resumeId uint, userRole string, resume DTO.ResumeUpdate) error
	DeleteResume(userId, resumeId uint, userRole string) error
}

type Work interface {
	CreateWork(userId, resumeId uint, userRole string, work DTO.WorkCreate) (uint, error)
	UpdateWork(userId, workId uint, userRole string, work DTO.WorkUpdate) error
	DeleteWork(userId, workId uint, userRole string) error
}

type Service struct {
	Authorization
	User
	Vacancy
	Resume
	Work
}

func NewService(repos *service.Repository) *Service {
	return &Service{
		Authorization: service.NewAuthService(repos.Authorization),
		User:          service.NewUserService(repos.User),
		Vacancy:       service.NewVacancyService(repos.Vacancy),
		Resume:        service.NewResumeService(repos.Resume),
		Work:          service.NewWorkService(repos.Resume, repos.Work),
	}
}
