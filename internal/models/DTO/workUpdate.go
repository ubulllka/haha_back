package DTO

import "time"

type WorkUpdate struct {
	Id          uint      `json:"id"`
	WorkId      uint      `json:"work_id"`
	Post        string    `json:"post"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
