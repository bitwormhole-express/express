package vo

import (
	"bitwomrhole.com/djaf/express-go-server/server/web/dto"
	"github.com/bitwormhole/starter-restful/api/vo"
)

// EmailVerification 定义邮件验证VO
type EmailVerification struct {
	vo.BaseVO

	Verification dto.EmailVerification `json:"verification"`
}
