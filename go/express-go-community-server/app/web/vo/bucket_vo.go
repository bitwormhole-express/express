package vo

import "github.com/bitwomrhole-express/express/community-server/app/web/dto"

type Bucket struct {
	Base

	Buckets []*dto.Bucket
}
