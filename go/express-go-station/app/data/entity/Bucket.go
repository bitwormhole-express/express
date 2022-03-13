package entity

import "github.com/bitwomrhole-express/express/station/app/data/dxo"

// Bucket 表示一个存储桶
type Bucket struct {
	Base

	ID dxo.BucketID `gorm:"primaryKey"`

	Provider string // 供应商
	Params   string
}
