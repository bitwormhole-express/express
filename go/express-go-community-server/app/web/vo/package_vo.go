package vo

import (
	"bitwomrhole.com/djaf/express-go-server/server/web/dto"
	"github.com/bitwormhole/starter-restful/api/vo"
)

type Package struct {
	vo.BaseVO
	Packages []*dto.Package `json:"packages"`
}
