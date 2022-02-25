package entity

import (
	"time"

	"gorm.io/gorm"
)

// Base 基本的数据库实体结构
type Base struct {

	// ID        uint           `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time

	DeleteAt gorm.DeletedAt `gorm:"index"`
	UUID     string         `gorm:"index:idx_name,unique"`
}
