// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	dao0xf4c226 "bitwomrhole.com/djaf/express-go-server/server/data/dao"
	security0x9ba940 "bitwomrhole.com/djaf/express-go-server/server/security"
	controller0x21caa6 "bitwomrhole.com/djaf/express-go-server/server/web/controller"
	interceptor0xbd9b3b "bitwomrhole.com/djaf/express-go-server/server/web/interceptor"
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
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
}


type pComDebugInterceptor struct {
	instance *interceptor0xbd9b3b.DebugInterceptor
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
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

