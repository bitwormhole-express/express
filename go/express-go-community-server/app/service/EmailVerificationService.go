package service

import "bitwomrhole.com/djaf/express-go-server/server/web/dto"

// EmailVerificationService 是提供邮件验证的服务
// 【inject:"#express-EmailVerificationService"】
type EmailVerificationService interface {

	// 开始验证：生成验证码，并发送邮件；
	StartVerification(email string) (*dto.EmailVerification, error)

	// 执行验证：检查验证码是否正确
	Verify(v *dto.EmailVerification) error
}
