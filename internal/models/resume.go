package models

import "github.com/jinzhu/gorm"

type Resume struct {
	gorm.Model
	Post              string              `json:"post"`
	Description       string              `json:"description"`
	ApplicantID       uint                `json:"applicant_id"`
	ResponseEmployers []ResponseEmployers `json:"response_employers" gorm:"foreignKey:ResumeID"`
	OldWorks          []Work              `json:"old_works" gorm:"foreignKey:ResumeID"`
}
