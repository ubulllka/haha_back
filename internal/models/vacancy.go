package models

import "github.com/jinzhu/gorm"

type Vacancy struct {
	gorm.Model
	Post        string `json:"post"`
	Description string `json:"description"`
	EmployerID  uint   `json:"employer_id"`
}
