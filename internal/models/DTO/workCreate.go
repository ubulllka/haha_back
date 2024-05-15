package DTO

import "time"

type WorkCreate struct {
	Post        string    `json:"post"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
