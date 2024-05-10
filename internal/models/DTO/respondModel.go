package DTO

type RespondModel struct {
	ResumeId  uint   `json:"resume_id"`
	VacancyId uint   `json:"vacancy_id"`
	Letter    string `json:"letter"`
}
