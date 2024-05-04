package models

import "github.com/jinzhu/gorm"

type ResponseApplicants struct {
	gorm.Model
	ApplicantID uint   `json:"applicant_id"`
	VacancyID   uint   `json:"vacancy_id"`
	Status      string `json:"status"`
}
