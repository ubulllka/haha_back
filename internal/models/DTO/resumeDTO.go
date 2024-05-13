package DTO

import "time"

type ResumeDTO struct {
	Id          uint      `json:"ID"`
	Post        string    `json:"post"`
	Description string    `json:"description"`
	ApplicantID uint      `json:"applicant_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

//vacToRes
