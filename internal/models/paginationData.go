package models

import (
	"github.com/jinzhu/gorm"
)

type PaginationData struct {
	PreviousPage int64 `json:"previous_page"`
	CurrentPage  int64 `json:"current_page"`
	NextPage     int64 `json:"next_page"`
	MaxPage      int64 `json:"max_page"`
}

func (pg *PaginationData) GetPagination(db *gorm.DB, page int64, q string, value interface{}) {
	var count int64
	db.Model(value).Where("post LIKE ?", "%"+q+"%").Count(&count)

	maxCount := count / 10
	if count%10 != 0 {
		maxCount++
	}

	if page < 1 {
		page = 1
	} else if page > maxCount {
		page = maxCount
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := page + 1
	if nextPage > maxCount {
		nextPage = maxCount
	}

	pg.PreviousPage = prevPage
	pg.CurrentPage = page
	pg.NextPage = nextPage
	pg.MaxPage = maxCount
}

func (pg *PaginationData) GetPaginationCnt(count, page, size int64) {
	maxCount := count / 10
	if count%10 != 0 {
		maxCount++
	}

	if page < 1 {
		page = 1
	} else if page > maxCount {
		page = maxCount
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := page + 1
	if nextPage > maxCount {
		nextPage = maxCount
	}

	pg.PreviousPage = prevPage
	pg.CurrentPage = page
	pg.NextPage = nextPage
	pg.MaxPage = maxCount
}
