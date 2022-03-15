package service

import (
	"context"

	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
)

// BucketService 是管理包裹存储桶的服务
// 【inject:"#express-BucketService"】
type BucketService interface {
	ListMyBuckets(c context.Context) ([]*dto.Bucket, error)

	InsertBucket(c context.Context, b *dto.Bucket) (*dto.Bucket, error)

	UpdateBucket(c context.Context, id dxo.BucketID, b *dto.Bucket) (*dto.Bucket, error)

	DeleteBucket(c context.Context, id dxo.BucketID) error
}
