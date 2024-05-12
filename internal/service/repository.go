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
	GetAll() ([]models.Vacancy, error)
	Search(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAllPag(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAll(userId uint) ([]DTO.ItemList, error)
	GetOne(vacancyId uint) (models.Vacancy, error)
	Create(vacancy models.Vacancy) (uint, error)
	Update(vacancyId uint, vacancy DTO.VacancyUpdate) error
	Delete(vacancyId uint) error
}

type Resume interface {
	GetAll() ([]models.Resume, error)
	Search(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error)
	GetApplAll(userId uint) ([]DTO.ItemList, error)
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

type Respond interface {
	CreateResToVac(respond DTO.RespondModel) error
	CreateVacToRes(respond DTO.RespondModel) error
	GetMyRespondAppl(userId uint, page int64, filter string) ([]DTO.RespondVacancy, models.PaginationData, error)
	GetMyRespondEmpl(userId uint, page int64, filter string) ([]DTO.RespondResume, models.PaginationData, error)
	GetOtherRespondAppl(userId uint, page int64, filter string) ([]DTO.RespondVacancy, models.PaginationData, error)
	GetOtherRespondEmpl(userId uint, page int64, filter string) ([]DTO.RespondResume, models.PaginationData, error)
}

type Repository struct {
	Authorization
	User
	Vacancy
	Resume
	Work
	Respond
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: repository.NewAuthPostgres(db),
		User:          repository.NewUserPostgres(db),
		Vacancy:       repository.NewVacancyPostgres(db),
		Resume:        repository.NewResumePostgres(db),
		Work:          repository.NewWorkPostgres(db),
		Respond:       repository.NewRespondPostgres(db),
	}
}
