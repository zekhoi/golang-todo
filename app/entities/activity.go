package entities

import (
	"time"
)

type Activity struct {
	ID        int        `json:"id" gorm:"primaryKey,autoIncrement"`
	Title     string     `json:"title" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
