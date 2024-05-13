package DTO

import "time"

type Respond struct {
	ID          uint      `json:"id"`
	Post        string    `json:"post"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Letter      string    `json:"letter"`
	VacancyId   uint      `json:"vacancy_id"`
	ResumeId    uint      `json:"resume_id"`
	OtherPost   string    `json:"other_post"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
