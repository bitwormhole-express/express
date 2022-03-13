package vo

import "github.com/bitwomrhole-express/express/community-server/app/web/dto"

type Upload struct {
	Base
	Items []*dto.Upload
}
