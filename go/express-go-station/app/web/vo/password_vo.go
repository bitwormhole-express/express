package vo

import (
	"github.com/bitwomrhole-express/express/station/app/web/dto"
	"github.com/bitwormhole/starter-restful/api/vo"
)

// Password 是操作密码的VO
type Password struct {
	vo.BaseVO
	Password dto.Password `json:"password"`
}
