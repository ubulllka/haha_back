package models

import "github.com/jinzhu/gorm"

//резюме, которые откликнулись на вакансии

type ResToVac struct {
	gorm.Model
	VacancyID uint   `json:"vacancy_id"`
	ResumeID  uint   `json:"resume_id"`
	Status    string `json:"status"`
}
