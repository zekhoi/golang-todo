package entities

import (
	"time"
)

type ToDo struct {
	ID              int        `json:"id" gorm:"primaryKey,autoIncrement"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title" gorm:"not null"`
	Priority        string     `json:"priority" gorm:"not null"`
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	CreatedAt       *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
