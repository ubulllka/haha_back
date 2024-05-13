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
	SearchAnon(page int64, q string) ([]models.Vacancy, models.PaginationData, error)
	GetOneAnon(vacancyId uint) (models.Vacancy, error)
	Search(userId uint, page int64, q string) ([]DTO.VacancyDTO, models.PaginationData, error)
	GetOne(userId, vacancyId uint) (DTO.VacancyDTO, error)

	GetEmplAllPag(userId uint, page int64) ([]models.Vacancy, models.PaginationData, error)
	GetEmplAll(userId uint) ([]DTO.ItemList, error)

	Create(vacancy models.Vacancy) (uint, error)
	Update(vacancyId uint, vacancy DTO.VacancyUpdate) error
	Delete(vacancyId uint) error
}

type Resume interface {
	GetAll() ([]models.Resume, error)
	SearchAnon(page int64, q string) ([]models.Resume, models.PaginationData, error)
	GetOneAnon(resumeId uint) (models.Resume, error)
	Search(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error)
	GetOne(userId, resumeId uint) (DTO.ResumeDTO, error)

	GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error)
	GetApplAll(userId uint) ([]DTO.ItemList, error)

	Create(resume models.Resume) (uint, error)
	Update(resumeId uint, resume DTO.ResumeUpdate) error
	Delete(resumeId uint) error
}

type Work interface {
	GetList(resumeId uint) ([]models.Work, error)
	GetOne(userId uint) (models.Work, error)
	Create(work models.Work) (uint, error)
	Update(workId uint, input DTO.WorkUpdate) error
	Delete(workId uint) error
}

type Respond interface {
	GetResToVac(id uint) (models.ResToVac, error)
	GetVacToRes(id uint) (models.VacToRes, error)
	CreateResToVac(respond DTO.RespondModel) error
	CreateVacToRes(respond DTO.RespondModel) error
	UpdateResToVac(id uint, respond DTO.RespondUpdate) error
	UpdateVacToRes(id uint, respond DTO.RespondUpdate) error
	DeleteResToVac(id uint) error
	DeleteVacToRes(id uint) error
	GetMyRespondAppl(id uint) (DTO.Respond, error)
	GetMyRespondEmpl(id uint) (DTO.Respond, error)
	GetOtherRespondAppl(id uint) (DTO.Respond, error)
	GetOtherRespondEmpl(id uint) (DTO.Respond, error)
	GetMyAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetMyAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetOtherAllRespondsAppl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
	GetOtherAllRespondsEmpl(userId uint, page int64, filter string) ([]DTO.Respond, models.PaginationData, error)
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
