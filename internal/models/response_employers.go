package models

import "github.com/jinzhu/gorm"

type ResponseEmployers struct {
	gorm.Model
	EmployerID uint   `json:"employer_id"`
	ResumeID   uint   `json:"resume_id"`
	Status     string `json:"status"`
}

//Status
//accept - принято
//decline - отклонено
//wait - ожидание
