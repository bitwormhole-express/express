// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	dao0xf91732 "github.com/bitwomrhole-express/express/community-server/app/data/dao"
	impl0x31d227 "github.com/bitwomrhole-express/express/community-server/app/data/dao/impl"
	security0xe21654 "github.com/bitwomrhole-express/express/community-server/app/security"
	service0x602df7 "github.com/bitwomrhole-express/express/community-server/app/service"
	impl0x79e5ab "github.com/bitwomrhole-express/express/community-server/app/service/impl"
	controller0xc95f51 "github.com/bitwomrhole-express/express/community-server/app/web/controller"
	interceptor0x488949 "github.com/bitwomrhole-express/express/community-server/app/web/interceptor"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	mail0xcd88fb "github.com/bitwormhole/starter-mail/mail"
	buckets0xc61cfb "github.com/bitwormhole/starter-object-bucket/buckets"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComAccountDaoImpl struct {
	instance *impl0x31d227.AccountDaoImpl
	 markup0x23084a.Component `id:"express-data-account-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	UUIDGenerator service0x602df7.UUIDGenerator `inject:"#the-uuid-generator"`
}


type pComAutoMigratorManager struct {
	instance *impl0x31d227.AutoMigratorManager
	 markup0x23084a.Component `initMethod:"Init"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	Items []impl0x31d227.AutoMigrator `inject:".express-server-data-auto-migrator"`
}


type pComBucketDaoImpl struct {
	instance *impl0x31d227.BucketDaoImpl
	 markup0x23084a.Component `id:"express-data-bucket-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	UUIDGenerator service0x602df7.UUIDGenerator `inject:"#the-uuid-generator"`
}


type pComBucketAgentImpl struct {
	instance *impl0x31d227.BucketAgentImpl
	 markup0x23084a.Component `id:"express-data-bucket-agent" initMethod:"Init"`
	BucketDriver string `inject:"${bucket.default.driver}"`
	BM buckets0xc61cfb.Manager `inject:"#buckets.Manager"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComPackageDaoImpl struct {
	instance *impl0x31d227.PackageDaoImpl
	 markup0x23084a.Component `id:"express-data-package-dao" class:"express-server-data-auto-migrator"`
	DS datasource0x68a737.Source `inject:"#gorm-datasource-default"`
	UUIDGenerator service0x602df7.UUIDGenerator `inject:"#the-uuid-generator"`
}


type pComEmailAuthenticator struct {
	instance *security0xe21654.EmailAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
	AccountDAO dao0xf91732.Account `inject:"#express-data-account-dao"`
	EmailVeriService service0x602df7.EmailVerificationService `inject:"#express-EmailVerificationService"`
}


type pComPasswordAuthenticator struct {
	instance *security0xe21654.PasswordAuthenticator
	 markup0x23084a.Component `class:"keeper-authenticator-registry"`
	AccountDAO dao0xf91732.Account `inject:"#express-data-account-dao"`
	PasswordService service0x602df7.PasswordService `inject:"#express-PasswordService"`
}


type pComEmailVerificationServiceImpl struct {
	instance *impl0x79e5ab.EmailVerificationServiceImpl
	 markup0x23084a.Component `id:"express-EmailVerificationService" initMethod:"Init"`
	Context application0x67f6c5.Context `inject:"context"`
	MailTemplateName string `inject:"${email.verification.template}"`
	MailTitle string `inject:"${email.verification.title}"`
	Sender mail0xcd88fb.Sender `inject:"#mail.Sender"`
	SenderAddr string `inject:"${mail.sender.address}"`
}


type pComPackageServiceImpl struct {
	instance *impl0x79e5ab.PackageServiceImpl
	 markup0x23084a.Component `id:"express-PackageService"`
	PackageDAO dao0xf91732.PackageDAO `inject:"#express-data-package-dao"`
}


type pComPasswordServiceImpl struct {
	instance *impl0x79e5ab.PasswordServiceImpl
	 markup0x23084a.Component `id:"express-PasswordService"`
	AccountDAO dao0xf91732.Account `inject:"#express-data-account-dao"`
	EmailVeriService service0x602df7.EmailVerificationService `inject:"#express-EmailVerificationService"`
}


type pComUploadServiceImpl struct {
	instance *impl0x79e5ab.UploadServiceImpl
	 markup0x23084a.Component `id:"express-UploadService"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
	BucketDrivers buckets0xc61cfb.Manager `inject:"#buckets.Manager"`
	BucketDAO dao0xf91732.BucketDAO `inject:"#express-data-bucket-dao"`
}


type pComUUIDGeneratorImpl struct {
	instance *impl0x79e5ab.UUIDGeneratorImpl
	 markup0x23084a.Component `id:"the-uuid-generator"`
}


type pComAuthController struct {
	instance *controller0xc95f51.AuthController
	 markup0x23084a.Component `class:"rest-controller"`
}


type pComBucketController struct {
	instance *controller0xc95f51.BucketController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComEmailVerificationController struct {
	instance *controller0xc95f51.EmailVerificationController
	 markup0x23084a.Component `class:"rest-controller"`
	EmailVeriService service0x602df7.EmailVerificationService `inject:"#express-EmailVerificationService"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComExampleController struct {
	instance *controller0xc95f51.ExampleController
	 markup0x23084a.Component `class:"rest-controller"`
}


type pComLogoutController struct {
	instance *controller0xc95f51.LogoutController
	 markup0x23084a.Component `class:"rest-controller"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
}


type pComPackageController struct {
	instance *controller0xc95f51.PackageController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	PackageService service0x602df7.PackageService `inject:"#express-PackageService"`
}


type pComPasswordController struct {
	instance *controller0xc95f51.PasswordController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	PasswordService service0x602df7.PasswordService `inject:"#express-PasswordService"`
}


type pComUploadController struct {
	instance *controller0xc95f51.UploadController
	 markup0x23084a.Component `class:"rest-controller"`
	Responder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	UploadService service0x602df7.UploadService `inject:"#express-UploadService"`
}


type pComDebugInterceptor struct {
	instance *interceptor0x488949.DebugInterceptor
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
}

