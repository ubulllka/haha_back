package DTO

type ResumeCreate struct {
	Post        string       `json:"post"`
	Description string       `json:"description"`
	OldWork     []WorkCreate `json:"old_work"`
}
