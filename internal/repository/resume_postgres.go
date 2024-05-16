package repository

import (
	"github.com/jinzhu/gorm"
	"haha/internal/logger"
	"haha/internal/models"
	"haha/internal/models/DTO"
)

type ResumePostgres struct {
	db   *gorm.DB
	logg *logger.Logger
}

func NewResumePostgres(db *gorm.DB, logg *logger.Logger) *ResumePostgres {
	return &ResumePostgres{db: db, logg: logg}
}

func (r *ResumePostgres) GetAllR() ([]models.Resume, error) {
	var resumes []models.Resume

	if err := r.db.Order("updated_at desc").Preload("OldWorks").
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, err
	}

	return resumes, nil
}

func (r *ResumePostgres) SearchAnon(page int64, q string) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume

	var count int64
	query := "%" + q + "%"

	dbBefore := r.db.Model(&models.Resume{}).Where("post LIKE ?", query).Count(&count)
	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}
	return resumes, pag, nil
}

func (r *ResumePostgres) GetOneAnon(resumeId uint) (models.Resume, error) {
	var resume models.Resume
	if err := r.db.First(&resume, resumeId).Error; err != nil {
		r.logg.Error(err)
		return models.Resume{}, err
	}
	return resume, nil
}

func (r *ResumePostgres) Search(userId uint, page int64, q string) ([]DTO.ResumeDTO, models.PaginationData, error) {
	var resumes []DTO.ResumeDTO

	var ids []string
	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).
		Pluck("id", &ids).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	var count int64
	query := "%" + q + "%"

	dbBefore := r.db.Table("resumes").
		Select("resumes.id as ID, post, description, applicant_id, resumes.created_at as created_at, resumes.updated_at as updated_at, status").
		Joins("left join vac_to_res on vac_to_res.resume_id=resumes.id AND vacancy_id IN (?)", ids).
		Where("post LIKE ?", query).Count(&count)

	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}
	pageSize := int64(10)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Find(&resumes).Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	return resumes, pag, nil
}

func (r *ResumePostgres) GetOne(userId, resumeId uint) (DTO.ResumeDTO, error) {
	var resume DTO.ResumeDTO
	var ids []string

	if err := r.db.Model(&models.Vacancy{}).Where("employer_id = ?", userId).
		Pluck("id", &ids).Error; err != nil {
		r.logg.Error(err)
		return DTO.ResumeDTO{}, err
	}

	if err := r.db.Table("resumes").
		Select("resumes.id as ID, post, description, applicant_id, resumes.created_at as created_at, resumes.updated_at as updated_at, status").
		Joins("left join vac_to_res on vac_to_res.resume_id=resumes.id AND vacancy_id IN (?)", ids).
		Find(&resume, resumeId).Error; err != nil {
		r.logg.Error(err)
		return DTO.ResumeDTO{}, err
	}
	return resume, nil
}

func (r *ResumePostgres) GetApplAllPag(userId uint, page int64) ([]models.Resume, models.PaginationData, error) {
	var resumes []models.Resume
	var count int64

	dbBefore := r.db.Model(&models.Resume{}).Where("applicant_id = ?", userId).Count(&count)
	if err := dbBefore.Error; err != nil {
		r.logg.Error(err)
		return nil, models.PaginationData{}, err
	}

	pageSize := int64(5)
	pag := models.PaginationData{}
	pag.Get(count, page, pageSize)

	if err := dbBefore.Order("updated_at desc").Scopes(Paginate(page, pageSize)).
		Preload("OldWorks").Find(&resumes).Error; err != nil {
		return nil, models.PaginationData{}, err
	}
	return resumes, pag, nil
}

func (r *ResumePostgres) GetApplAll(userId uint) ([]DTO.ItemList, error) {
	var resumes []DTO.ItemList

	if err := r.db.Table("resumes").Select("id, post").Where("applicant_id = ?", userId).
		Order("updated_at desc").Find(&resumes).Error; err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *ResumePostgres) Create(userId uint, resumeDTO DTO.ResumeCreate) error {
	arrWork := make([]models.Work, 0)
	for _, v := range resumeDTO.OldWorks {
		arrWork = append(arrWork, models.Work{
			Post:        v.Post,
			Description: v.Description,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
		})
	}

	resume := models.Resume{
		Post:        resumeDTO.Post,
		Description: resumeDTO.Description,
		ApplicantID: userId,
		OldWorks:    arrWork,
	}
	return r.db.Create(&resume).Error
}

func remove(arr []DTO.WorkUpdate, id uint) []DTO.WorkUpdate {
	newArr := make([]DTO.WorkUpdate, 0)
	for _, v := range arr {
		if v.Id != id {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func (r *ResumePostgres) Update(resumeId uint, input DTO.ResumeUpdate) error {

	oldArr := input.OldWorksOld
	newArr := input.OldWorksNew
	tx := r.db.Begin()

	oldWork := make([]models.Work, 0)
	for _, v := range newArr {
		if v.WorkId != 0 {
			var work models.Work
			if err := tx.First(&work, v.WorkId).Error; err != nil {
				r.logg.Error(err)
				tx.Rollback()
				return err
			}
			work.Post = v.Post
			work.Description = v.Description
			work.StartTime = v.StartTime
			work.EndTime = v.EndTime
			if err := tx.Save(&work).Error; err != nil {
				r.logg.Error(err)
				tx.Rollback()
				return err
			}
			oldArr = remove(oldArr, v.Id)
			oldWork = append(oldWork, work)
		} else {
			work := models.Work{
				Post:        v.Post,
				Description: v.Description,
				StartTime:   v.StartTime,
				EndTime:     v.EndTime,
			}
			if err := tx.Save(&work).Error; err != nil {
				r.logg.Error(err)
				tx.Rollback()
				return err
			}
			oldArr = remove(oldArr, v.Id)
			oldWork = append(oldWork, work)
		}
	}

	for _, v := range oldArr {
		if err := tx.Unscoped().Delete(models.Work{}, v.WorkId).Error; err != nil {
			r.logg.Error(err)
			tx.Rollback()
			return err
		}
	}

	var resume models.Resume
	if err := tx.First(&resume, resumeId).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}

	resume.Post = input.Post
	resume.Description = input.Description
	resume.OldWorks = oldWork

	if err := tx.Save(&resume).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *ResumePostgres) Delete(resumeId uint) error {
	tx := r.db.Begin()

	if err := tx.Unscoped().Where("resume_id = ?", resumeId).Delete(models.Work{}).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("resume_id = ?", resumeId).Delete(models.ResToVac{}).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("resume_id = ?", resumeId).Delete(models.VacToRes{}).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Delete(&models.Resume{}, resumeId).Error; err != nil {
		r.logg.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
