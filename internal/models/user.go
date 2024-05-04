package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name             string               `json:"name"`
	Email            string               `json:"email" gorm:"type:varchar(100);unique"`
	Telegram         string               `json:"telegram"`
	Password         string               `json:"password"`
	Role             string               `json:"role"`
	Status           string               `json:"status"`
	ApplicantVacancy []ResponseApplicants `json:"applicant_vacancy" gorm:"foreignKey:ApplicantID"`
	EmployerVacancy  []Vacancy            `json:"employer_vacancy" gorm:"foreignKey:EmployerID"`
	ApplicantResume  []Resume             `json:"applicant_resume" gorm:"foreignKey:ApplicantID"`
	EmployerResume   []ResponseEmployers  `json:"employer_resume" gorm:"foreignKey:EmployerID"`
}
