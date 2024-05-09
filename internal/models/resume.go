package models

import "github.com/jinzhu/gorm"

type Resume struct {
	gorm.Model
	Post        string     `json:"post"`
	Description string     `json:"description"`
	ApplicantID uint       `json:"applicant_id"`
	OldWorks    []Work     `json:"old_works" gorm:"foreignKey:ResumeID"`
	ResToVac    []ResToVac `json:"res_to_vac" gorm:"foreignKey:ResumeID"`
	VacToRes    []VacToRes `json:"vac_to_res" gorm:"foreignKey:ResumeID"`
}
