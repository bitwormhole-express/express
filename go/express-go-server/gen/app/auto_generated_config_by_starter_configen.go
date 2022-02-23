// (todo:gen2.template) 
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

	// component: express-data-account-dao
	cominfobuilder.Next()
	cominfobuilder.ID("express-data-account-dao").Class("express-server-data-auto-migrator").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAccountDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-dao0xf4c226.AutoMigratorManager
	cominfobuilder.Next()
	cominfobuilder.ID("com1-dao0xf4c226.AutoMigratorManager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAutoMigratorManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-security0x9ba940.EmailAuthenticator
	cominfobuilder.Next()
	cominfobuilder.ID("com2-security0x9ba940.EmailAuthenticator").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEmailAuthenticator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-security0x9ba940.PasswordAuthenticator
	cominfobuilder.Next()
	cominfobuilder.ID("com3-security0x9ba940.PasswordAuthenticator").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPasswordAuthenticator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-interceptor0xbd9b3b.DebugInterceptor
	cominfobuilder.Next()
	cominfobuilder.ID("com4-interceptor0xbd9b3b.DebugInterceptor").Class("rest-interceptor-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDebugInterceptor{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-controller0x21caa6.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com5-controller0x21caa6.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-controller0x21caa6.LogoutController
	cominfobuilder.Next()
	cominfobuilder.ID("com6-controller0x21caa6.LogoutController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLogoutController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAccountDaoImpl : the factory of component: express-data-account-dao
type comFactory4pComAccountDaoImpl struct {

    mPrototype * dao0xf4c226.AccountDaoImpl

	
	mDSSelector config.InjectionSelector

}

func (inst * comFactory4pComAccountDaoImpl) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAccountDaoImpl) newObject() * dao0xf4c226.AccountDaoImpl {
	return & dao0xf4c226.AccountDaoImpl {}
}

func (inst * comFactory4pComAccountDaoImpl) castObject(instance application.ComponentInstance) * dao0xf4c226.AccountDaoImpl {
	return instance.Get().(*dao0xf4c226.AccountDaoImpl)
}

func (inst * comFactory4pComAccountDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAccountDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAccountDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAccountDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAccountDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAccountDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DS = inst.getterForFieldDSSelector(context)
	return context.LastError()
}

//getterForFieldDSSelector
func (inst * comFactory4pComAccountDaoImpl) getterForFieldDSSelector (context application.InstanceContext) datasource0x68a737.Source {

	o1 := inst.mDSSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Source)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-account-dao")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAutoMigratorManager : the factory of component: com1-dao0xf4c226.AutoMigratorManager
type comFactory4pComAutoMigratorManager struct {

    mPrototype * dao0xf4c226.AutoMigratorManager

	
	mDSSelector config.InjectionSelector
	mItemsSelector config.InjectionSelector

}

func (inst * comFactory4pComAutoMigratorManager) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)
	inst.mItemsSelector = config.NewInjectionSelector(".express-server-data-auto-migrator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAutoMigratorManager) newObject() * dao0xf4c226.AutoMigratorManager {
	return & dao0xf4c226.AutoMigratorManager {}
}

func (inst * comFactory4pComAutoMigratorManager) castObject(instance application.ComponentInstance) * dao0xf4c226.AutoMigratorManager {
	return instance.Get().(*dao0xf4c226.AutoMigratorManager)
}

func (inst * comFactory4pComAutoMigratorManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAutoMigratorManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAutoMigratorManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAutoMigratorManager) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComAutoMigratorManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAutoMigratorManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DS = inst.getterForFieldDSSelector(context)
	obj.Items = inst.getterForFieldItemsSelector(context)
	return context.LastError()
}

//getterForFieldDSSelector
func (inst * comFactory4pComAutoMigratorManager) getterForFieldDSSelector (context application.InstanceContext) datasource0x68a737.Source {

	o1 := inst.mDSSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Source)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-dao0xf4c226.AutoMigratorManager")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldItemsSelector
func (inst * comFactory4pComAutoMigratorManager) getterForFieldItemsSelector (context application.InstanceContext) []dao0xf4c226.AutoMigrator {
	list1 := inst.mItemsSelector.GetList(context)
	list2 := make([]dao0xf4c226.AutoMigrator, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(dao0xf4c226.AutoMigrator)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEmailAuthenticator : the factory of component: com2-security0x9ba940.EmailAuthenticator
type comFactory4pComEmailAuthenticator struct {

    mPrototype * security0x9ba940.EmailAuthenticator

	

}

func (inst * comFactory4pComEmailAuthenticator) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEmailAuthenticator) newObject() * security0x9ba940.EmailAuthenticator {
	return & security0x9ba940.EmailAuthenticator {}
}

func (inst * comFactory4pComEmailAuthenticator) castObject(instance application.ComponentInstance) * security0x9ba940.EmailAuthenticator {
	return instance.Get().(*security0x9ba940.EmailAuthenticator)
}

func (inst * comFactory4pComEmailAuthenticator) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComEmailAuthenticator) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComEmailAuthenticator) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComEmailAuthenticator) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEmailAuthenticator) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEmailAuthenticator) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPasswordAuthenticator : the factory of component: com3-security0x9ba940.PasswordAuthenticator
type comFactory4pComPasswordAuthenticator struct {

    mPrototype * security0x9ba940.PasswordAuthenticator

	
	mAccountDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComPasswordAuthenticator) init() application.ComponentFactory {

	
	inst.mAccountDAOSelector = config.NewInjectionSelector("#express-data-account-dao",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPasswordAuthenticator) newObject() * security0x9ba940.PasswordAuthenticator {
	return & security0x9ba940.PasswordAuthenticator {}
}

func (inst * comFactory4pComPasswordAuthenticator) castObject(instance application.ComponentInstance) * security0x9ba940.PasswordAuthenticator {
	return instance.Get().(*security0x9ba940.PasswordAuthenticator)
}

func (inst * comFactory4pComPasswordAuthenticator) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPasswordAuthenticator) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPasswordAuthenticator) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPasswordAuthenticator) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordAuthenticator) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordAuthenticator) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AccountDAO = inst.getterForFieldAccountDAOSelector(context)
	return context.LastError()
}

//getterForFieldAccountDAOSelector
func (inst * comFactory4pComPasswordAuthenticator) getterForFieldAccountDAOSelector (context application.InstanceContext) dao0xf4c226.Account {

	o1 := inst.mAccountDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf4c226.Account)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com3-security0x9ba940.PasswordAuthenticator")
		eb.Set("field", "AccountDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf4c226.Account")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDebugInterceptor : the factory of component: com4-interceptor0xbd9b3b.DebugInterceptor
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

// comFactory4pComExampleController : the factory of component: com5-controller0x21caa6.ExampleController
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



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLogoutController : the factory of component: com6-controller0x21caa6.LogoutController
type comFactory4pComLogoutController struct {

    mPrototype * controller0x21caa6.LogoutController

	
	mSubjectsSelector config.InjectionSelector

}

func (inst * comFactory4pComLogoutController) init() application.ComponentFactory {

	
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLogoutController) newObject() * controller0x21caa6.LogoutController {
	return & controller0x21caa6.LogoutController {}
}

func (inst * comFactory4pComLogoutController) castObject(instance application.ComponentInstance) * controller0x21caa6.LogoutController {
	return instance.Get().(*controller0x21caa6.LogoutController)
}

func (inst * comFactory4pComLogoutController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLogoutController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLogoutController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLogoutController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLogoutController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLogoutController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	return context.LastError()
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComLogoutController) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-controller0x21caa6.LogoutController")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




