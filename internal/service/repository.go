package service

import (
	"github.com/jinzhu/gorm"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"haha/internal/repository"
)

type Authorization interface {
	Create(user models.User) (uint, error)
	GetOne(email, password string) (models.User, error)
	GetOneById(id uint) (models.User, error)
}

type User interface {
	GetAll() ([]models.User, error)
	Update(id uint, user DTO.UserUpdate) error
	GetOneById(id uint) (models.User, error)
}

type Vacancy interface {
	GetAll(page int64) ([]models.Vacancy, models.PaginationData, error)
	Search(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAll(userId uint) ([]models.Vacancy, error)
	GetOne(vacancyId uint) (models.Vacancy, error)
	Create(vacancy models.Vacancy) (uint, error)
	Update(vacancyId uint, vacancy DTO.VacancyUpdate) error
	Delete(vacancyId uint) error
}

type Resume interface {
	GetAll(page int64) ([]models.Resume, models.PaginationData, error)
	Search(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetApplAll(userId uint) ([]models.Resume, error)
	GetOne(resumeId uint) (models.Resume, error)
	Create(resume models.Resume) (uint, error)
	Update(resumeId uint, resume DTO.ResumeUpdate) error
	Delete(resumeId uint) error
}

type Work interface {
	GetOne(userId uint) (models.Work, error)
	Create(work models.Work) (uint, error)
	Update(workId uint, input DTO.WorkUpdate) error
	Delete(workId uint) error
}

type Repository struct {
	Authorization
	User
	Vacancy
	Resume
	Work
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: repository.NewAuthPostgres(db),
		User:          repository.NewUserPostgres(db),
		Vacancy:       repository.NewVacancyPostgres(db),
		Resume:        repository.NewResumePostgres(db),
		Work:          repository.NewWorkPostgres(db),
	}
}
