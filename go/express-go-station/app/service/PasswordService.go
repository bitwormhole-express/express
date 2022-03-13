package service

import "github.com/bitwomrhole-express/express/station/app/web/dto"

// PasswordService 提供密码相关的服务
// 【inject:"#express-PasswordService"】
type PasswordService interface {
	Verify(username string, password []byte) error
	CreateAccount(o *dto.Password) error
	ResetPassword(o *dto.Password) error
}
