package models

import "github.com/jinzhu/gorm"

//вакансии, которые откликнулись на резюме

type VacToRes struct {
	gorm.Model
	VacancyID uint   `json:"vacancy_id"`
	ResumeID  uint   `json:"resume_id"`
	Status    string `json:"status"`
}
