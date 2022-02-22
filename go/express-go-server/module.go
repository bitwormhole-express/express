package expressgoserver

import (
	"embed"

	"bitwomrhole.com/djaf/express-go-server/gen/app"
	"github.com/bitwormhole/starter"
	ginstarter "github.com/bitwormhole/starter-gin"
	starterginsecurity "github.com/bitwormhole/starter-gin-security"
	gormstarter "github.com/bitwormhole/starter-gorm"
	mysql "github.com/bitwormhole/starter-gorm-mysql"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "bitwomrhole.com/djaf/express-go-server"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

//go:embed "src/main/resources"
var theModuleSrcMainRes embed.FS

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
	mb.Dependency(starterginsecurity.Module())

	mb.Dependency(ginstarter.ModuleWithDevtools())

	return mb.Create()
}
