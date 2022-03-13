// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	dao0xeb8470 "github.com/bitwomrhole-express/express/station/app/data/dao"
	impl0x7304fe "github.com/bitwomrhole-express/express/station/app/data/dao/impl"
	security0x584176 "github.com/bitwomrhole-express/express/station/app/security"
	service0x336bd6 "github.com/bitwomrhole-express/express/station/app/service"
	impl0xeadaf8 "github.com/bitwomrhole-express/express/station/app/service/impl"
	controller0xebdfab "github.com/bitwomrhole-express/express/station/app/web/controller"
	interceptor0x20f6c3 "github.com/bitwomrhole-express/express/station/app/web/interceptor"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	mail0xcd88fb "github.com/bitwormhole/starter-mail/mail"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComPasswordServiceImpl struct {
	instance *impl0xeadaf8.PasswordServiceImpl
	 markup0x23084a.Component `id:"express-PasswordService"`
	AccountDAO dao0xeb8470.Account `inject:"#express-data-account-dao"`
	EmailVeriService service0x336bd6.EmailVerificationService `inject:"#express-EmailVerificationService"`
}


type pComEmailVerificationServiceImpl struct {
	instance *impl0xeadaf8.EmailVerificationServiceImpl
	 markup0x23084a.Component `id:"express-EmailVerificationService" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	MailTemplateName string `inject:"${email.verification.template}"`
	MailTitle string `inject:"${email.verification.title}"`
	Sender mail0xcd88fb.Sender `inject:"#mail.Sender"`
	SenderAddr string `inject:"${mail.sender.address}"`
}


type pComUUIDGeneratorImpl struct {
	instance *impl0xeadaf8.UUIDGeneratorImpl
	 markup0x23084a.Component `id:"the-uuid-generator"`
}


type pComPackageServiceImpl struct {
	instance *impl0xeadaf8.PackageServiceImpl
	 markup0x23084a.Component `id:"express-PackageService"`
	PackageDAO dao0xeb8470.PackageDAO `inject:"#express-data-package-dao"`
}


type pComAccountDaoImpl struct {
	instance *impl0x7304fe.AccountDaoImpl
	 markup0x23084a.Component `id:"express-data-account-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	UUIDGenerator service0x336bd6.UUIDGenerator `inject:"#the-uuid-generator"`
}


type pComPackageDaoImpl struct {
	instance *impl0x7304fe.PackageDaoImpl
	 markup0x23084a.Component `id:"express-data-package-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	UUIDGenerator service0x336bd6.UUIDGenerator `inject:"#the-uuid-generator"`
}


type pComAutoMigratorManager struct {
	instance *impl0x7304fe.AutoMigratorManager
	 markup0x23084a.Component `initMethod:"Init"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	Items []impl0x7304fe.AutoMigrator `inject:".express-server-data-auto-migrator"`
}


type pComEmailAuthenticator struct {
	instance *security0x584176.EmailAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
	AccountDAO dao0xeb8470.Account `inject:"#express-data-account-dao"`
	EmailVeriService service0x336bd6.EmailVerificationService `inject:"#express-EmailVerificationService"`
}


type pComPasswordAuthenticator struct {
	instance *security0x584176.PasswordAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
	AccountDAO dao0xeb8470.Account `inject:"#express-data-account-dao"`
	PasswordService service0x336bd6.PasswordService `inject:"#express-PasswordService"`
}


type pComDebugInterceptor struct {
	instance *interceptor0x20f6c3.DebugInterceptor
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
}


type pComExampleController struct {
	instance *controller0xebdfab.ExampleController
	 markup0x23084a.Component `class:"rest-controller"`
}


type pComAuthController struct {
	instance *controller0xebdfab.AuthController
	 markup0x23084a.Component `class:"rest-controller"`
}


type pComPasswordController struct {
	instance *controller0xebdfab.PasswordController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	PasswordService service0x336bd6.PasswordService `inject:"#express-PasswordService"`
}


type pComLogoutController struct {
	instance *controller0xebdfab.LogoutController
	 markup0x23084a.Component `class:"rest-controller"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
}


type pComEmailVerificationController struct {
	instance *controller0xebdfab.EmailVerificationController
	 markup0x23084a.Component `class:"rest-controller"`
	EmailVeriService service0x336bd6.EmailVerificationService `inject:"#express-EmailVerificationService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComPackageController struct {
	instance *controller0xebdfab.PackageController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	PackageService service0x336bd6.PackageService `inject:"#express-PackageService"`
}

