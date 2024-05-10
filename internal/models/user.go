package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email" gorm:"type:varchar(100);unique"`
	Telegram    string    `json:"telegram"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Resumes     []Resume  `json:"resumes"`   // Applicant
	Vacancies   []Vacancy `json:"vacancies"` //Employer
}
