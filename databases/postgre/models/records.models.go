package models

import "time"

type Records struct {
	Id        int64     `json:"id" gorm:"column:id"`
	Total     int       `json:"totalMarks" gorm:"column:total"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}
