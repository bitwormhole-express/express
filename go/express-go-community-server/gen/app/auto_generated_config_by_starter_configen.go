// (todo:gen2.template) 
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

	// component: com1-impl0x31d227.AutoMigratorManager
	cominfobuilder.Next()
	cominfobuilder.ID("com1-impl0x31d227.AutoMigratorManager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAutoMigratorManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-data-bucket-dao
	cominfobuilder.Next()
	cominfobuilder.ID("express-data-bucket-dao").Class("express-server-data-auto-migrator").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBucketDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-data-bucket-agent
	cominfobuilder.Next()
	cominfobuilder.ID("express-data-bucket-agent").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBucketAgentImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-data-package-dao
	cominfobuilder.Next()
	cominfobuilder.ID("express-data-package-dao").Class("express-server-data-auto-migrator").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPackageDaoImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-security0xe21654.EmailAuthenticator
	cominfobuilder.Next()
	cominfobuilder.ID("com5-security0xe21654.EmailAuthenticator").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEmailAuthenticator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-security0xe21654.PasswordAuthenticator
	cominfobuilder.Next()
	cominfobuilder.ID("com6-security0xe21654.PasswordAuthenticator").Class("keeper-authenticator-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPasswordAuthenticator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-EmailVerificationService
	cominfobuilder.Next()
	cominfobuilder.ID("express-EmailVerificationService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEmailVerificationServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-PackageService
	cominfobuilder.Next()
	cominfobuilder.ID("express-PackageService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPackageServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-PasswordService
	cominfobuilder.Next()
	cominfobuilder.ID("express-PasswordService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPasswordServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: express-UploadService
	cominfobuilder.Next()
	cominfobuilder.ID("express-UploadService").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUploadServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: the-uuid-generator
	cominfobuilder.Next()
	cominfobuilder.ID("the-uuid-generator").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUUIDGeneratorImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com12-controller0xc95f51.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com12-controller0xc95f51.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com13-controller0xc95f51.BucketController
	cominfobuilder.Next()
	cominfobuilder.ID("com13-controller0xc95f51.BucketController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBucketController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com14-controller0xc95f51.EmailVerificationController
	cominfobuilder.Next()
	cominfobuilder.ID("com14-controller0xc95f51.EmailVerificationController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEmailVerificationController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com15-controller0xc95f51.ExampleController
	cominfobuilder.Next()
	cominfobuilder.ID("com15-controller0xc95f51.ExampleController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComExampleController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com16-controller0xc95f51.LogoutController
	cominfobuilder.Next()
	cominfobuilder.ID("com16-controller0xc95f51.LogoutController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLogoutController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com17-controller0xc95f51.PackageController
	cominfobuilder.Next()
	cominfobuilder.ID("com17-controller0xc95f51.PackageController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPackageController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com18-controller0xc95f51.PasswordController
	cominfobuilder.Next()
	cominfobuilder.ID("com18-controller0xc95f51.PasswordController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPasswordController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com19-controller0xc95f51.UploadController
	cominfobuilder.Next()
	cominfobuilder.ID("com19-controller0xc95f51.UploadController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUploadController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com20-interceptor0x488949.DebugInterceptor
	cominfobuilder.Next()
	cominfobuilder.ID("com20-interceptor0x488949.DebugInterceptor").Class("rest-interceptor-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDebugInterceptor{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAccountDaoImpl : the factory of component: express-data-account-dao
type comFactory4pComAccountDaoImpl struct {

    mPrototype * impl0x31d227.AccountDaoImpl

	
	mDSSelector config.InjectionSelector
	mUUIDGeneratorSelector config.InjectionSelector

}

func (inst * comFactory4pComAccountDaoImpl) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)
	inst.mUUIDGeneratorSelector = config.NewInjectionSelector("#the-uuid-generator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAccountDaoImpl) newObject() * impl0x31d227.AccountDaoImpl {
	return & impl0x31d227.AccountDaoImpl {}
}

func (inst * comFactory4pComAccountDaoImpl) castObject(instance application.ComponentInstance) * impl0x31d227.AccountDaoImpl {
	return instance.Get().(*impl0x31d227.AccountDaoImpl)
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
	obj.UUIDGenerator = inst.getterForFieldUUIDGeneratorSelector(context)
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

//getterForFieldUUIDGeneratorSelector
func (inst * comFactory4pComAccountDaoImpl) getterForFieldUUIDGeneratorSelector (context application.InstanceContext) service0x602df7.UUIDGenerator {

	o1 := inst.mUUIDGeneratorSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.UUIDGenerator)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-account-dao")
		eb.Set("field", "UUIDGenerator")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.UUIDGenerator")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAutoMigratorManager : the factory of component: com1-impl0x31d227.AutoMigratorManager
type comFactory4pComAutoMigratorManager struct {

    mPrototype * impl0x31d227.AutoMigratorManager

	
	mDSSelector config.InjectionSelector
	mItemsSelector config.InjectionSelector

}

func (inst * comFactory4pComAutoMigratorManager) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)
	inst.mItemsSelector = config.NewInjectionSelector(".express-server-data-auto-migrator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAutoMigratorManager) newObject() * impl0x31d227.AutoMigratorManager {
	return & impl0x31d227.AutoMigratorManager {}
}

func (inst * comFactory4pComAutoMigratorManager) castObject(instance application.ComponentInstance) * impl0x31d227.AutoMigratorManager {
	return instance.Get().(*impl0x31d227.AutoMigratorManager)
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
		eb.Set("com", "com1-impl0x31d227.AutoMigratorManager")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldItemsSelector
func (inst * comFactory4pComAutoMigratorManager) getterForFieldItemsSelector (context application.InstanceContext) []impl0x31d227.AutoMigrator {
	list1 := inst.mItemsSelector.GetList(context)
	list2 := make([]impl0x31d227.AutoMigrator, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(impl0x31d227.AutoMigrator)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBucketDaoImpl : the factory of component: express-data-bucket-dao
type comFactory4pComBucketDaoImpl struct {

    mPrototype * impl0x31d227.BucketDaoImpl

	
	mDSSelector config.InjectionSelector
	mUUIDGeneratorSelector config.InjectionSelector

}

func (inst * comFactory4pComBucketDaoImpl) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)
	inst.mUUIDGeneratorSelector = config.NewInjectionSelector("#the-uuid-generator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBucketDaoImpl) newObject() * impl0x31d227.BucketDaoImpl {
	return & impl0x31d227.BucketDaoImpl {}
}

func (inst * comFactory4pComBucketDaoImpl) castObject(instance application.ComponentInstance) * impl0x31d227.BucketDaoImpl {
	return instance.Get().(*impl0x31d227.BucketDaoImpl)
}

func (inst * comFactory4pComBucketDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBucketDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBucketDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBucketDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBucketDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBucketDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DS = inst.getterForFieldDSSelector(context)
	obj.UUIDGenerator = inst.getterForFieldUUIDGeneratorSelector(context)
	return context.LastError()
}

//getterForFieldDSSelector
func (inst * comFactory4pComBucketDaoImpl) getterForFieldDSSelector (context application.InstanceContext) datasource0x68a737.Source {

	o1 := inst.mDSSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Source)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-bucket-dao")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGeneratorSelector
func (inst * comFactory4pComBucketDaoImpl) getterForFieldUUIDGeneratorSelector (context application.InstanceContext) service0x602df7.UUIDGenerator {

	o1 := inst.mUUIDGeneratorSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.UUIDGenerator)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-bucket-dao")
		eb.Set("field", "UUIDGenerator")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.UUIDGenerator")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBucketAgentImpl : the factory of component: express-data-bucket-agent
type comFactory4pComBucketAgentImpl struct {

    mPrototype * impl0x31d227.BucketAgentImpl

	
	mBucketDriverSelector config.InjectionSelector
	mBMSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComBucketAgentImpl) init() application.ComponentFactory {

	
	inst.mBucketDriverSelector = config.NewInjectionSelector("${bucket.default.driver}",nil)
	inst.mBMSelector = config.NewInjectionSelector("#buckets.Manager",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBucketAgentImpl) newObject() * impl0x31d227.BucketAgentImpl {
	return & impl0x31d227.BucketAgentImpl {}
}

func (inst * comFactory4pComBucketAgentImpl) castObject(instance application.ComponentInstance) * impl0x31d227.BucketAgentImpl {
	return instance.Get().(*impl0x31d227.BucketAgentImpl)
}

func (inst * comFactory4pComBucketAgentImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBucketAgentImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBucketAgentImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBucketAgentImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComBucketAgentImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBucketAgentImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.BucketDriver = inst.getterForFieldBucketDriverSelector(context)
	obj.BM = inst.getterForFieldBMSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldBucketDriverSelector
func (inst * comFactory4pComBucketAgentImpl) getterForFieldBucketDriverSelector (context application.InstanceContext) string {
    return inst.mBucketDriverSelector.GetString(context)
}

//getterForFieldBMSelector
func (inst * comFactory4pComBucketAgentImpl) getterForFieldBMSelector (context application.InstanceContext) buckets0xc61cfb.Manager {

	o1 := inst.mBMSelector.GetOne(context)
	o2, ok := o1.(buckets0xc61cfb.Manager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-bucket-agent")
		eb.Set("field", "BM")
		eb.Set("type1", "?")
		eb.Set("type2", "buckets0xc61cfb.Manager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComBucketAgentImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPackageDaoImpl : the factory of component: express-data-package-dao
type comFactory4pComPackageDaoImpl struct {

    mPrototype * impl0x31d227.PackageDaoImpl

	
	mDSSelector config.InjectionSelector
	mUUIDGeneratorSelector config.InjectionSelector

}

func (inst * comFactory4pComPackageDaoImpl) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#gorm-datasource-default",nil)
	inst.mUUIDGeneratorSelector = config.NewInjectionSelector("#the-uuid-generator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPackageDaoImpl) newObject() * impl0x31d227.PackageDaoImpl {
	return & impl0x31d227.PackageDaoImpl {}
}

func (inst * comFactory4pComPackageDaoImpl) castObject(instance application.ComponentInstance) * impl0x31d227.PackageDaoImpl {
	return instance.Get().(*impl0x31d227.PackageDaoImpl)
}

func (inst * comFactory4pComPackageDaoImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPackageDaoImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPackageDaoImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPackageDaoImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageDaoImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageDaoImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DS = inst.getterForFieldDSSelector(context)
	obj.UUIDGenerator = inst.getterForFieldUUIDGeneratorSelector(context)
	return context.LastError()
}

//getterForFieldDSSelector
func (inst * comFactory4pComPackageDaoImpl) getterForFieldDSSelector (context application.InstanceContext) datasource0x68a737.Source {

	o1 := inst.mDSSelector.GetOne(context)
	o2, ok := o1.(datasource0x68a737.Source)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-package-dao")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "datasource0x68a737.Source")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUUIDGeneratorSelector
func (inst * comFactory4pComPackageDaoImpl) getterForFieldUUIDGeneratorSelector (context application.InstanceContext) service0x602df7.UUIDGenerator {

	o1 := inst.mUUIDGeneratorSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.UUIDGenerator)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-data-package-dao")
		eb.Set("field", "UUIDGenerator")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.UUIDGenerator")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEmailAuthenticator : the factory of component: com5-security0xe21654.EmailAuthenticator
type comFactory4pComEmailAuthenticator struct {

    mPrototype * security0xe21654.EmailAuthenticator

	
	mAccountDAOSelector config.InjectionSelector
	mEmailVeriServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComEmailAuthenticator) init() application.ComponentFactory {

	
	inst.mAccountDAOSelector = config.NewInjectionSelector("#express-data-account-dao",nil)
	inst.mEmailVeriServiceSelector = config.NewInjectionSelector("#express-EmailVerificationService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEmailAuthenticator) newObject() * security0xe21654.EmailAuthenticator {
	return & security0xe21654.EmailAuthenticator {}
}

func (inst * comFactory4pComEmailAuthenticator) castObject(instance application.ComponentInstance) * security0xe21654.EmailAuthenticator {
	return instance.Get().(*security0xe21654.EmailAuthenticator)
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
	
	obj := inst.castObject(instance)
	obj.AccountDAO = inst.getterForFieldAccountDAOSelector(context)
	obj.EmailVeriService = inst.getterForFieldEmailVeriServiceSelector(context)
	return context.LastError()
}

//getterForFieldAccountDAOSelector
func (inst * comFactory4pComEmailAuthenticator) getterForFieldAccountDAOSelector (context application.InstanceContext) dao0xf91732.Account {

	o1 := inst.mAccountDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf91732.Account)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-security0xe21654.EmailAuthenticator")
		eb.Set("field", "AccountDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf91732.Account")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEmailVeriServiceSelector
func (inst * comFactory4pComEmailAuthenticator) getterForFieldEmailVeriServiceSelector (context application.InstanceContext) service0x602df7.EmailVerificationService {

	o1 := inst.mEmailVeriServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.EmailVerificationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-security0xe21654.EmailAuthenticator")
		eb.Set("field", "EmailVeriService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.EmailVerificationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPasswordAuthenticator : the factory of component: com6-security0xe21654.PasswordAuthenticator
type comFactory4pComPasswordAuthenticator struct {

    mPrototype * security0xe21654.PasswordAuthenticator

	
	mAccountDAOSelector config.InjectionSelector
	mPasswordServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPasswordAuthenticator) init() application.ComponentFactory {

	
	inst.mAccountDAOSelector = config.NewInjectionSelector("#express-data-account-dao",nil)
	inst.mPasswordServiceSelector = config.NewInjectionSelector("#express-PasswordService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPasswordAuthenticator) newObject() * security0xe21654.PasswordAuthenticator {
	return & security0xe21654.PasswordAuthenticator {}
}

func (inst * comFactory4pComPasswordAuthenticator) castObject(instance application.ComponentInstance) * security0xe21654.PasswordAuthenticator {
	return instance.Get().(*security0xe21654.PasswordAuthenticator)
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
	obj.PasswordService = inst.getterForFieldPasswordServiceSelector(context)
	return context.LastError()
}

//getterForFieldAccountDAOSelector
func (inst * comFactory4pComPasswordAuthenticator) getterForFieldAccountDAOSelector (context application.InstanceContext) dao0xf91732.Account {

	o1 := inst.mAccountDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf91732.Account)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-security0xe21654.PasswordAuthenticator")
		eb.Set("field", "AccountDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf91732.Account")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPasswordServiceSelector
func (inst * comFactory4pComPasswordAuthenticator) getterForFieldPasswordServiceSelector (context application.InstanceContext) service0x602df7.PasswordService {

	o1 := inst.mPasswordServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.PasswordService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-security0xe21654.PasswordAuthenticator")
		eb.Set("field", "PasswordService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.PasswordService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEmailVerificationServiceImpl : the factory of component: express-EmailVerificationService
type comFactory4pComEmailVerificationServiceImpl struct {

    mPrototype * impl0x79e5ab.EmailVerificationServiceImpl

	
	mContextSelector config.InjectionSelector
	mMailTemplateNameSelector config.InjectionSelector
	mMailTitleSelector config.InjectionSelector
	mSenderSelector config.InjectionSelector
	mSenderAddrSelector config.InjectionSelector

}

func (inst * comFactory4pComEmailVerificationServiceImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mMailTemplateNameSelector = config.NewInjectionSelector("${email.verification.template}",nil)
	inst.mMailTitleSelector = config.NewInjectionSelector("${email.verification.title}",nil)
	inst.mSenderSelector = config.NewInjectionSelector("#mail.Sender",nil)
	inst.mSenderAddrSelector = config.NewInjectionSelector("${mail.sender.address}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEmailVerificationServiceImpl) newObject() * impl0x79e5ab.EmailVerificationServiceImpl {
	return & impl0x79e5ab.EmailVerificationServiceImpl {}
}

func (inst * comFactory4pComEmailVerificationServiceImpl) castObject(instance application.ComponentInstance) * impl0x79e5ab.EmailVerificationServiceImpl {
	return instance.Get().(*impl0x79e5ab.EmailVerificationServiceImpl)
}

func (inst * comFactory4pComEmailVerificationServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComEmailVerificationServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComEmailVerificationServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComEmailVerificationServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComEmailVerificationServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEmailVerificationServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.MailTemplateName = inst.getterForFieldMailTemplateNameSelector(context)
	obj.MailTitle = inst.getterForFieldMailTitleSelector(context)
	obj.Sender = inst.getterForFieldSenderSelector(context)
	obj.SenderAddr = inst.getterForFieldSenderAddrSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComEmailVerificationServiceImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldMailTemplateNameSelector
func (inst * comFactory4pComEmailVerificationServiceImpl) getterForFieldMailTemplateNameSelector (context application.InstanceContext) string {
    return inst.mMailTemplateNameSelector.GetString(context)
}

//getterForFieldMailTitleSelector
func (inst * comFactory4pComEmailVerificationServiceImpl) getterForFieldMailTitleSelector (context application.InstanceContext) string {
    return inst.mMailTitleSelector.GetString(context)
}

//getterForFieldSenderSelector
func (inst * comFactory4pComEmailVerificationServiceImpl) getterForFieldSenderSelector (context application.InstanceContext) mail0xcd88fb.Sender {

	o1 := inst.mSenderSelector.GetOne(context)
	o2, ok := o1.(mail0xcd88fb.Sender)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-EmailVerificationService")
		eb.Set("field", "Sender")
		eb.Set("type1", "?")
		eb.Set("type2", "mail0xcd88fb.Sender")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSenderAddrSelector
func (inst * comFactory4pComEmailVerificationServiceImpl) getterForFieldSenderAddrSelector (context application.InstanceContext) string {
    return inst.mSenderAddrSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPackageServiceImpl : the factory of component: express-PackageService
type comFactory4pComPackageServiceImpl struct {

    mPrototype * impl0x79e5ab.PackageServiceImpl

	
	mPackageDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComPackageServiceImpl) init() application.ComponentFactory {

	
	inst.mPackageDAOSelector = config.NewInjectionSelector("#express-data-package-dao",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPackageServiceImpl) newObject() * impl0x79e5ab.PackageServiceImpl {
	return & impl0x79e5ab.PackageServiceImpl {}
}

func (inst * comFactory4pComPackageServiceImpl) castObject(instance application.ComponentInstance) * impl0x79e5ab.PackageServiceImpl {
	return instance.Get().(*impl0x79e5ab.PackageServiceImpl)
}

func (inst * comFactory4pComPackageServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPackageServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPackageServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPackageServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PackageDAO = inst.getterForFieldPackageDAOSelector(context)
	return context.LastError()
}

//getterForFieldPackageDAOSelector
func (inst * comFactory4pComPackageServiceImpl) getterForFieldPackageDAOSelector (context application.InstanceContext) dao0xf91732.PackageDAO {

	o1 := inst.mPackageDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf91732.PackageDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-PackageService")
		eb.Set("field", "PackageDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf91732.PackageDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPasswordServiceImpl : the factory of component: express-PasswordService
type comFactory4pComPasswordServiceImpl struct {

    mPrototype * impl0x79e5ab.PasswordServiceImpl

	
	mAccountDAOSelector config.InjectionSelector
	mEmailVeriServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPasswordServiceImpl) init() application.ComponentFactory {

	
	inst.mAccountDAOSelector = config.NewInjectionSelector("#express-data-account-dao",nil)
	inst.mEmailVeriServiceSelector = config.NewInjectionSelector("#express-EmailVerificationService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPasswordServiceImpl) newObject() * impl0x79e5ab.PasswordServiceImpl {
	return & impl0x79e5ab.PasswordServiceImpl {}
}

func (inst * comFactory4pComPasswordServiceImpl) castObject(instance application.ComponentInstance) * impl0x79e5ab.PasswordServiceImpl {
	return instance.Get().(*impl0x79e5ab.PasswordServiceImpl)
}

func (inst * comFactory4pComPasswordServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPasswordServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPasswordServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPasswordServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AccountDAO = inst.getterForFieldAccountDAOSelector(context)
	obj.EmailVeriService = inst.getterForFieldEmailVeriServiceSelector(context)
	return context.LastError()
}

//getterForFieldAccountDAOSelector
func (inst * comFactory4pComPasswordServiceImpl) getterForFieldAccountDAOSelector (context application.InstanceContext) dao0xf91732.Account {

	o1 := inst.mAccountDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf91732.Account)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-PasswordService")
		eb.Set("field", "AccountDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf91732.Account")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEmailVeriServiceSelector
func (inst * comFactory4pComPasswordServiceImpl) getterForFieldEmailVeriServiceSelector (context application.InstanceContext) service0x602df7.EmailVerificationService {

	o1 := inst.mEmailVeriServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.EmailVerificationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-PasswordService")
		eb.Set("field", "EmailVeriService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.EmailVerificationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUploadServiceImpl : the factory of component: express-UploadService
type comFactory4pComUploadServiceImpl struct {

    mPrototype * impl0x79e5ab.UploadServiceImpl

	
	mSubjectsSelector config.InjectionSelector
	mBucketDriversSelector config.InjectionSelector
	mBucketDAOSelector config.InjectionSelector

}

func (inst * comFactory4pComUploadServiceImpl) init() application.ComponentFactory {

	
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)
	inst.mBucketDriversSelector = config.NewInjectionSelector("#buckets.Manager",nil)
	inst.mBucketDAOSelector = config.NewInjectionSelector("#express-data-bucket-dao",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUploadServiceImpl) newObject() * impl0x79e5ab.UploadServiceImpl {
	return & impl0x79e5ab.UploadServiceImpl {}
}

func (inst * comFactory4pComUploadServiceImpl) castObject(instance application.ComponentInstance) * impl0x79e5ab.UploadServiceImpl {
	return instance.Get().(*impl0x79e5ab.UploadServiceImpl)
}

func (inst * comFactory4pComUploadServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUploadServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUploadServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUploadServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	obj.BucketDrivers = inst.getterForFieldBucketDriversSelector(context)
	obj.BucketDAO = inst.getterForFieldBucketDAOSelector(context)
	return context.LastError()
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComUploadServiceImpl) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-UploadService")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldBucketDriversSelector
func (inst * comFactory4pComUploadServiceImpl) getterForFieldBucketDriversSelector (context application.InstanceContext) buckets0xc61cfb.Manager {

	o1 := inst.mBucketDriversSelector.GetOne(context)
	o2, ok := o1.(buckets0xc61cfb.Manager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-UploadService")
		eb.Set("field", "BucketDrivers")
		eb.Set("type1", "?")
		eb.Set("type2", "buckets0xc61cfb.Manager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldBucketDAOSelector
func (inst * comFactory4pComUploadServiceImpl) getterForFieldBucketDAOSelector (context application.InstanceContext) dao0xf91732.BucketDAO {

	o1 := inst.mBucketDAOSelector.GetOne(context)
	o2, ok := o1.(dao0xf91732.BucketDAO)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "express-UploadService")
		eb.Set("field", "BucketDAO")
		eb.Set("type1", "?")
		eb.Set("type2", "dao0xf91732.BucketDAO")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUUIDGeneratorImpl : the factory of component: the-uuid-generator
type comFactory4pComUUIDGeneratorImpl struct {

    mPrototype * impl0x79e5ab.UUIDGeneratorImpl

	

}

func (inst * comFactory4pComUUIDGeneratorImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUUIDGeneratorImpl) newObject() * impl0x79e5ab.UUIDGeneratorImpl {
	return & impl0x79e5ab.UUIDGeneratorImpl {}
}

func (inst * comFactory4pComUUIDGeneratorImpl) castObject(instance application.ComponentInstance) * impl0x79e5ab.UUIDGeneratorImpl {
	return instance.Get().(*impl0x79e5ab.UUIDGeneratorImpl)
}

func (inst * comFactory4pComUUIDGeneratorImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUUIDGeneratorImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUUIDGeneratorImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUUIDGeneratorImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUUIDGeneratorImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUUIDGeneratorImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAuthController : the factory of component: com12-controller0xc95f51.AuthController
type comFactory4pComAuthController struct {

    mPrototype * controller0xc95f51.AuthController

	

}

func (inst * comFactory4pComAuthController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAuthController) newObject() * controller0xc95f51.AuthController {
	return & controller0xc95f51.AuthController {}
}

func (inst * comFactory4pComAuthController) castObject(instance application.ComponentInstance) * controller0xc95f51.AuthController {
	return instance.Get().(*controller0xc95f51.AuthController)
}

func (inst * comFactory4pComAuthController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAuthController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAuthController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAuthController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBucketController : the factory of component: com13-controller0xc95f51.BucketController
type comFactory4pComBucketController struct {

    mPrototype * controller0xc95f51.BucketController

	
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComBucketController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBucketController) newObject() * controller0xc95f51.BucketController {
	return & controller0xc95f51.BucketController {}
}

func (inst * comFactory4pComBucketController) castObject(instance application.ComponentInstance) * controller0xc95f51.BucketController {
	return instance.Get().(*controller0xc95f51.BucketController)
}

func (inst * comFactory4pComBucketController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBucketController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBucketController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBucketController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBucketController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBucketController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComBucketController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com13-controller0xc95f51.BucketController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEmailVerificationController : the factory of component: com14-controller0xc95f51.EmailVerificationController
type comFactory4pComEmailVerificationController struct {

    mPrototype * controller0xc95f51.EmailVerificationController

	
	mEmailVeriServiceSelector config.InjectionSelector
	mResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComEmailVerificationController) init() application.ComponentFactory {

	
	inst.mEmailVeriServiceSelector = config.NewInjectionSelector("#express-EmailVerificationService",nil)
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEmailVerificationController) newObject() * controller0xc95f51.EmailVerificationController {
	return & controller0xc95f51.EmailVerificationController {}
}

func (inst * comFactory4pComEmailVerificationController) castObject(instance application.ComponentInstance) * controller0xc95f51.EmailVerificationController {
	return instance.Get().(*controller0xc95f51.EmailVerificationController)
}

func (inst * comFactory4pComEmailVerificationController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComEmailVerificationController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComEmailVerificationController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComEmailVerificationController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEmailVerificationController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEmailVerificationController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.EmailVeriService = inst.getterForFieldEmailVeriServiceSelector(context)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	return context.LastError()
}

//getterForFieldEmailVeriServiceSelector
func (inst * comFactory4pComEmailVerificationController) getterForFieldEmailVeriServiceSelector (context application.InstanceContext) service0x602df7.EmailVerificationService {

	o1 := inst.mEmailVeriServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.EmailVerificationService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com14-controller0xc95f51.EmailVerificationController")
		eb.Set("field", "EmailVeriService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.EmailVerificationService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldResponderSelector
func (inst * comFactory4pComEmailVerificationController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com14-controller0xc95f51.EmailVerificationController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComExampleController : the factory of component: com15-controller0xc95f51.ExampleController
type comFactory4pComExampleController struct {

    mPrototype * controller0xc95f51.ExampleController

	

}

func (inst * comFactory4pComExampleController) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComExampleController) newObject() * controller0xc95f51.ExampleController {
	return & controller0xc95f51.ExampleController {}
}

func (inst * comFactory4pComExampleController) castObject(instance application.ComponentInstance) * controller0xc95f51.ExampleController {
	return instance.Get().(*controller0xc95f51.ExampleController)
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

// comFactory4pComLogoutController : the factory of component: com16-controller0xc95f51.LogoutController
type comFactory4pComLogoutController struct {

    mPrototype * controller0xc95f51.LogoutController

	
	mSubjectsSelector config.InjectionSelector

}

func (inst * comFactory4pComLogoutController) init() application.ComponentFactory {

	
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLogoutController) newObject() * controller0xc95f51.LogoutController {
	return & controller0xc95f51.LogoutController {}
}

func (inst * comFactory4pComLogoutController) castObject(instance application.ComponentInstance) * controller0xc95f51.LogoutController {
	return instance.Get().(*controller0xc95f51.LogoutController)
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
		eb.Set("com", "com16-controller0xc95f51.LogoutController")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPackageController : the factory of component: com17-controller0xc95f51.PackageController
type comFactory4pComPackageController struct {

    mPrototype * controller0xc95f51.PackageController

	
	mResponderSelector config.InjectionSelector
	mPackageServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPackageController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mPackageServiceSelector = config.NewInjectionSelector("#express-PackageService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPackageController) newObject() * controller0xc95f51.PackageController {
	return & controller0xc95f51.PackageController {}
}

func (inst * comFactory4pComPackageController) castObject(instance application.ComponentInstance) * controller0xc95f51.PackageController {
	return instance.Get().(*controller0xc95f51.PackageController)
}

func (inst * comFactory4pComPackageController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPackageController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPackageController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPackageController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	obj.PackageService = inst.getterForFieldPackageServiceSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComPackageController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com17-controller0xc95f51.PackageController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPackageServiceSelector
func (inst * comFactory4pComPackageController) getterForFieldPackageServiceSelector (context application.InstanceContext) service0x602df7.PackageService {

	o1 := inst.mPackageServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.PackageService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com17-controller0xc95f51.PackageController")
		eb.Set("field", "PackageService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.PackageService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPasswordController : the factory of component: com18-controller0xc95f51.PasswordController
type comFactory4pComPasswordController struct {

    mPrototype * controller0xc95f51.PasswordController

	
	mResponderSelector config.InjectionSelector
	mPasswordServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComPasswordController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mPasswordServiceSelector = config.NewInjectionSelector("#express-PasswordService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPasswordController) newObject() * controller0xc95f51.PasswordController {
	return & controller0xc95f51.PasswordController {}
}

func (inst * comFactory4pComPasswordController) castObject(instance application.ComponentInstance) * controller0xc95f51.PasswordController {
	return instance.Get().(*controller0xc95f51.PasswordController)
}

func (inst * comFactory4pComPasswordController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPasswordController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPasswordController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPasswordController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPasswordController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	obj.PasswordService = inst.getterForFieldPasswordServiceSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComPasswordController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com18-controller0xc95f51.PasswordController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPasswordServiceSelector
func (inst * comFactory4pComPasswordController) getterForFieldPasswordServiceSelector (context application.InstanceContext) service0x602df7.PasswordService {

	o1 := inst.mPasswordServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.PasswordService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com18-controller0xc95f51.PasswordController")
		eb.Set("field", "PasswordService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.PasswordService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUploadController : the factory of component: com19-controller0xc95f51.UploadController
type comFactory4pComUploadController struct {

    mPrototype * controller0xc95f51.UploadController

	
	mResponderSelector config.InjectionSelector
	mUploadServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComUploadController) init() application.ComponentFactory {

	
	inst.mResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mUploadServiceSelector = config.NewInjectionSelector("#express-UploadService",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUploadController) newObject() * controller0xc95f51.UploadController {
	return & controller0xc95f51.UploadController {}
}

func (inst * comFactory4pComUploadController) castObject(instance application.ComponentInstance) * controller0xc95f51.UploadController {
	return instance.Get().(*controller0xc95f51.UploadController)
}

func (inst * comFactory4pComUploadController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUploadController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUploadController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUploadController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUploadController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Responder = inst.getterForFieldResponderSelector(context)
	obj.UploadService = inst.getterForFieldUploadServiceSelector(context)
	return context.LastError()
}

//getterForFieldResponderSelector
func (inst * comFactory4pComUploadController) getterForFieldResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com19-controller0xc95f51.UploadController")
		eb.Set("field", "Responder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldUploadServiceSelector
func (inst * comFactory4pComUploadController) getterForFieldUploadServiceSelector (context application.InstanceContext) service0x602df7.UploadService {

	o1 := inst.mUploadServiceSelector.GetOne(context)
	o2, ok := o1.(service0x602df7.UploadService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com19-controller0xc95f51.UploadController")
		eb.Set("field", "UploadService")
		eb.Set("type1", "?")
		eb.Set("type2", "service0x602df7.UploadService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDebugInterceptor : the factory of component: com20-interceptor0x488949.DebugInterceptor
type comFactory4pComDebugInterceptor struct {

    mPrototype * interceptor0x488949.DebugInterceptor

	

}

func (inst * comFactory4pComDebugInterceptor) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDebugInterceptor) newObject() * interceptor0x488949.DebugInterceptor {
	return & interceptor0x488949.DebugInterceptor {}
}

func (inst * comFactory4pComDebugInterceptor) castObject(instance application.ComponentInstance) * interceptor0x488949.DebugInterceptor {
	return instance.Get().(*interceptor0x488949.DebugInterceptor)
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




