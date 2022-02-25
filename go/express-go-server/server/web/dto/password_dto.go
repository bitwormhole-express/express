package dto

import "github.com/bitwormhole/starter/util"

// Password 是操作密码的DTO
type Password struct {
	Verification EmailVerification `json:"verification"`
	NewPassword  util.Base64       `json:"newpassword"`
}
