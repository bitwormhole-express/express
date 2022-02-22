// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package app

import (
	security0x9ba940 "bitwomrhole.com/djaf/express-go-server/server/security"
	controller0x21caa6 "bitwomrhole.com/djaf/express-go-server/server/web/controller"
	interceptor0xbd9b3b "bitwomrhole.com/djaf/express-go-server/server/web/interceptor"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: keeper-authenticator-registry-2233
	cominfobuilder.Next()
	cominfobuilder.ID("keeper-authenticator-registry-2233").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLoginWithPasswordHandler{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-interceptor0xbd9b3b.DebugInterceptor
	cominfobuilder.Next()
	cominfobuilder.ID("com1-interceptor0xbd9b3b.DebugInterceptor").Class("rest-interceptor-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDebugInterceptor{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-controller0x21caa6.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com2-controller0x21caa6.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLoginWithPasswordHandler : the factory of component: keeper-authenticator-registry-2233
type comFactory4pComLoginWithPasswordHandler struct {

    mPrototype * security0x9ba940.LoginWithPasswordHandler

	

}

func (inst * comFactory4pComLoginWithPasswordHandler) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLoginWithPasswordHandler) newObject() * security0x9ba940.LoginWithPasswordHandler {
	return & security0x9ba940.LoginWithPasswordHandler {}
}

func (inst * comFactory4pComLoginWithPasswordHandler) castObject(instance application.ComponentInstance) * security0x9ba940.LoginWithPasswordHandler {
	return instance.Get().(*security0x9ba940.LoginWithPasswordHandler)
}

func (inst * comFactory4pComLoginWithPasswordHandler) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLoginWithPasswordHandler) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLoginWithPasswordHandler) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLoginWithPasswordHandler) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLoginWithPasswordHandler) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLoginWithPasswordHandler) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDebugInterceptor : the factory of component: com1-interceptor0xbd9b3b.DebugInterceptor
type comFactory4pComDebugInterceptor struct {

    mPrototype * interceptor0xbd9b3b.DebugInterceptor

	

}

func (inst * comFactory4pComDebugInterceptor) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDebugInterceptor) newObject() * interceptor0xbd9b3b.DebugInterceptor {
	return & interceptor0xbd9b3b.DebugInterceptor {}
}

func (inst * comFactory4pComDebugInterceptor) castObject(instance application.ComponentInstance) * interceptor0xbd9b3b.DebugInterceptor {
	return instance.Get().(*interceptor0xbd9b3b.DebugInterceptor)
}

func (inst * comFactory4pComDebugInterceptor) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDebugInterceptor) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDebugInterceptor) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDebugInterceptor) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDebugInterceptor) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDebugInterceptor) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com2-controller0x21caa6.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0x21caa6.ExampleController

	

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0x21caa6.ExampleController {
	return & controller0x21caa6.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0x21caa6.ExampleController {
	return instance.Get().(*controller0x21caa6.ExampleController)
}

func (inst * comFactory4pComExampleController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComExampleController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComExampleController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComExampleController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComExampleController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




