// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	security0x9ba940 "bitwomrhole.com/djaf/express-go-server/server/security"
	controller0x21caa6 "bitwomrhole.com/djaf/express-go-server/server/web/controller"
	interceptor0xbd9b3b "bitwomrhole.com/djaf/express-go-server/server/web/interceptor"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComLoginWithPasswordHandler struct {
	instance *security0x9ba940.LoginWithPasswordHandler
	 markup0x23084a.Component `id:"keeper-authenticator-registry-2233" class:"keeper-authenticator-registry"`
}


type pComDebugInterceptor struct {
	instance *interceptor0xbd9b3b.DebugInterceptor
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
}


type pComExampleController struct {
	instance *controller0x21caa6.ExampleController
	 markup0x23084a.Component `class:"rest-controller"`
}

