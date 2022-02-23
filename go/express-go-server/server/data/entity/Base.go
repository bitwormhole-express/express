package entity

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}
