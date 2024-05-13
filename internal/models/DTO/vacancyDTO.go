package DTO

import "time"

type VacancyDTO struct {
	Id          uint      `json:"ID"`
	Post        string    `json:"post"`
	Description string    `json:"description"`
	EmployerID  uint      `json:"employer_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

//resToVac
