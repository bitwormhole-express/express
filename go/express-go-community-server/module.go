package communityserver

import (
	"embed"

	"github.com/bitwomrhole-express/express/community-server/gen/app"
	"github.com/bitwormhole/starter"
	ginstarter "github.com/bitwormhole/starter-gin"
	starterginsecurity "github.com/bitwormhole/starter-gin-security"
	gormstarter "github.com/bitwormhole/starter-gorm"
	mysql "github.com/bitwormhole/starter-gorm-mysql"
	sqlserver "github.com/bitwormhole/starter-gorm-sqlserver"
	startermail "github.com/bitwormhole/starter-mail"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwomrhole-express/express/community-server"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

//go:embed "src/main/resources"
var theModuleSrcMainRes embed.FS

// Module 导出 app 主模块
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.Resources(collection.LoadEmbedResources(&theModuleSrcMainRes, "src/main/resources"))
	mb.OnMount(app.ExportConfigForExpressGoServerApp)

	mb.Dependency(starter.Module())
	mb.Dependency(ginstarter.Module())
	mb.Dependency(gormstarter.Module())
	mb.Dependency(mysql.DriverModule())
	mb.Dependency(sqlserver.DriverModule())
	mb.Dependency(starterginsecurity.Module())
	mb.Dependency(startermail.Module())

	mb.Dependency(ginstarter.ModuleWithDevtools())

	return mb.Create()
}
