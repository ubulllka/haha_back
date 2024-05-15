package DTO

type ResumeUpdate struct {
	Post        string       `json:"post"`
	Description string       `json:"description"`
	OldWorksOld []WorkUpdate `json:"old_works_old"`
	OldWorksNew []WorkUpdate `json:"old_works_new"`
}
