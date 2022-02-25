// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	dao0xf4c226 "bitwomrhole.com/djaf/express-go-server/server/data/dao"
	security0x9ba940 "bitwomrhole.com/djaf/express-go-server/server/security"
	service0xd29b29 "bitwomrhole.com/djaf/express-go-server/server/service"
	impl0x29d072 "bitwomrhole.com/djaf/express-go-server/server/service/impl"
	controller0x21caa6 "bitwomrhole.com/djaf/express-go-server/server/web/controller"
	interceptor0xbd9b3b "bitwomrhole.com/djaf/express-go-server/server/web/interceptor"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	mail0xcd88fb "github.com/bitwormhole/starter-mail/mail"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComAccountDaoImpl struct {
	instance *dao0xf4c226.AccountDaoImpl
	 markup0x23084a.Component `id:"express-data-account-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
}


type pComAutoMigratorManager struct {
	instance *dao0xf4c226.AutoMigratorManager
	 markup0x23084a.Component `initMethod:"Init"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	Items []dao0xf4c226.AutoMigrator `inject:".express-server-data-auto-migrator"`
}


type pComEmailAuthenticator struct {
	instance *security0x9ba940.EmailAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
}


type pComPasswordAuthenticator struct {
	instance *security0x9ba940.PasswordAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
	AccountDAO dao0xf4c226.Account `inject:"#express-data-account-dao"`
	PasswordService service0xd29b29.PasswordService `inject:"#express-PasswordService"`
}


type pComEmailVerificationServiceImpl struct {
	instance *impl0x29d072.EmailVerificationServiceImpl
	 markup0x23084a.Component `id:"express-EmailVerificationService" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	MailTemplateName string `inject:"${email.verification.template}"`
	MailTitle string `inject:"${email.verification.title}"`
	Sender mail0xcd88fb.Sender `inject:"#mail.Sender"`
	SenderAddr string `inject:"${mail.sender.address}"`
}


type pComPasswordServiceImpl struct {
	instance *impl0x29d072.PasswordServiceImpl
	 markup0x23084a.Component `id:"express-PasswordService"`
	AccountDAO dao0xf4c226.Account `inject:"#express-data-account-dao"`
	EmailVeriService service0xd29b29.EmailVerificationService `inject:"#express-EmailVerificationService"`
}


type pComEmailVerificationController struct {
	instance *controller0x21caa6.EmailVerificationController
	 markup0x23084a.Component `class:"rest-controller"`
	EmailVeriService service0xd29b29.EmailVerificationService `inject:"#express-EmailVerificationService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExampleController struct {
	instance *controller0x21caa6.ExampleController
	 markup0x23084a.Component `class:"rest-controller"`
}


type pComLogoutController struct {
	instance *controller0x21caa6.LogoutController
	 markup0x23084a.Component `class:"rest-controller"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
}


type pComPasswordController struct {
	instance *controller0x21caa6.PasswordController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	PasswordService service0xd29b29.PasswordService `inject:"#express-PasswordService"`
}


type pComDebugInterceptor struct {
	instance *interceptor0xbd9b3b.DebugInterceptor
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
}

