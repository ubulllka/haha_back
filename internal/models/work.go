package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Work struct {
	gorm.Model
	Post        string    `json:"post"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	ResumeID    uint      `json:"resume_id"`
}
