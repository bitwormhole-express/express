package vo

import (
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
	"github.com/bitwormhole/starter-restful/api/vo"
)

type Package struct {
	vo.BaseVO
	Packages []*dto.Package `json:"packages"`
}
