package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
	"haha/internal/repository"
)

var (
	errAuth = errors.New("not enough rights")
)

type Authorization interface {
	GetOne(email, password string) (models.User, error)
	GetOneById(id uint) (models.User, error)

	Create(user models.User) (uint, error)
}

type User interface {
	GetAll() ([]models.User, error)
	GetOneById(id uint) (models.User, error)

	Update(id uint, user DTO.UserUpdate) error
}

type Vacancy interface {
	GetAllV() ([]models.Vacancy, error)
	SearchAnon(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetOneAnon(vacancyId uint) (models.Vacancy, error)
	Search(userId uint, page int64, q string) ([]DTO.VacancyDTO, models.PaginationData, error)
	GetOne(userId, vacancyId uint) (DTO.VacancyDTO, error)

	GetEmplAllPag(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAll(userId uint) ([]DTO.ItemList, error)

	Create(vacancy models.Vacancy) error
	Update(vacancyId uint, vacancy DTO.VacancyUpdate) error
	Delete(vacancyId uint) error
}

type Resume interface {
	GetAllR() ([]models.Resume, error)
	SearchAnon(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetOneAnon(resumeId uint) (models.Resume, error)
	Search(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error)
	GetOne(userId, resumeId uint) (DTO.ResumeDTO, error)

	GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error)
	GetApplAll(userId uint) ([]DTO.ItemList, error)

	Create(userId uint, resume DTO.ResumeCreate) error
	Update(resumeId uint, resume DTO.ResumeUpdate) error
	Delete(resumeId uint) error
}

type Work interface {
	GetList(resumeId uint) ([]models.Work, error)
	GetOne(userId uint) (models.Work, error)

	//Create(work models.Work) error
	//Update(workId uint, input DTO.WorkUpdate) error
	Delete(workId uint) error
}

type Respond interface {
	GetResToVac(id uint) (models.ResToVac, error)
	GetVacToRes(id uint) (models.VacToRes, error)

	GetMyAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetMyAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetOtherAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetOtherAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)

	CreateResToVac(respond DTO.RespondModel) error
	CreateVacToRes(respond DTO.RespondModel) error
	UpdateResToVac(id uint, respond DTO.RespondUpdate) error
	UpdateVacToRes(id uint, respond DTO.RespondUpdate) error
	DeleteResToVac(id uint) error
	DeleteVacToRes(id uint) error
}

type Repository struct {
	Authorization
	User
	Vacancy
	Resume
	Work
	Respond
}

func NewRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{
		Authorization: repository.NewAuthPostgres(db, logger),
		User:          repository.NewUserPostgres(db, logger),
		Vacancy:       repository.NewVacancyPostgres(db, logger),
		Resume:        repository.NewResumePostgres(db, logger),
		Work:          repository.NewWorkPostgres(db, logger),
		Respond:       repository.NewRespondPostgres(db, logger),
	}
}
