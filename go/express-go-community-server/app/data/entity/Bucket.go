package entity

import (
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwormhole/starter/util"
)

// Bucket 表示一个存储桶
type Bucket struct {
	Base

	ID dxo.BucketID `gorm:"primaryKey"`

	Driver string // 供应商

	Properties util.Base64
}

// TableName 设置表名
func (Bucket) TableName() string {
	return TableNameBucket
}
