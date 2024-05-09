package models

import "github.com/jinzhu/gorm"

type Vacancy struct {
	gorm.Model
	Post        string     `json:"post"`
	Description string     `json:"description"`
	EmployerID  uint       `json:"employer_id"`
	ResToVac    []ResToVac `json:"res_to_vac" gorm:"foreignKey:VacancyID"`
	VacToRes    []VacToRes `json:"vac_to_res" gorm:"foreignKey:VacancyID"`
}
