package dto

import "github.com/bitwomrhole-express/express/community-server/app/data/dxo"

type Bucket struct {
	Base

	ID dxo.BucketID

	Driver string

	Properties map[string]string
}
