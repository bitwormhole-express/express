package entity

import (
	"time"

	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"gorm.io/gorm"
)

// Base 基本的数据库实体结构
type Base struct {

	// ID        uint           `gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time

	DeleteAt  gorm.DeletedAt `gorm:"index"`
	UUID      string         `gorm:"index:idx_name,unique"`
	Owner     dxo.AccountID
	Committer dxo.AccountID
}
