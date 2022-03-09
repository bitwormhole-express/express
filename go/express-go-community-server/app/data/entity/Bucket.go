package entity

import "bitwomrhole.com/djaf/express-go-server/server/data/dxo"

// Bucket 表示一个存储桶
type Bucket struct {
	Base

	ID dxo.BucketID `gorm:"primaryKey"`

	Provider string // 供应商
	Params   string
}
