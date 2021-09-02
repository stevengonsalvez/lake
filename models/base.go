package models

import "time"

type Model struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}