package dto

import (
	"github.com/bitwormhole/starter-restful/api/dto"
	"github.com/bitwormhole/starter/util"
)

// EmailVerification 定义邮件验证DTO
type EmailVerification struct {
	dto.BaseDTO

	Code      string    `json:"code"`   // 验证码
	Email     string    `json:"email"`  // email 地址
	Nonce     string    `json:"nonce"`  // 随机数
	Secret    string    `json:"secret"` // 密钥
	TimeBegin util.Time `json:"time1"`  // 起效时间戳
	TimeEnd   util.Time `json:"time2"`  // 过期时间戳

	Sha256sum string `json:"sha256sum"`
}
