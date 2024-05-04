package DTO

type WorkUpdate struct {
	Post        *string `json:"post"`
	Description *string `json:"description"`
	StartTime   *string `json:"start_time"`
	EndTime     *string `json:"end_time"`
}
