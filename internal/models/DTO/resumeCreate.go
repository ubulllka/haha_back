package DTO

type ResumeCreate struct {
	Post        string       `json:"post"`
	Description string       `json:"description"`
	OldWorks    []WorkCreate `json:"old_works"`
}
