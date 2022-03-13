package impl

import (
	"time"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// AutoMigrator 实现自动建表的接口
// 【inject:".express-server-data-auto-migrator"】
type AutoMigrator interface {
	AutoMigrate() error
}

type AutoMigratorManager struct {
	markup.Component `initMethod:"Init"`

	DS    datasource.Source `inject:"#gorm-datasource-default"`
	Items []AutoMigrator    `inject:".express-server-data-auto-migrator"`
}

func (inst *AutoMigratorManager) Init() error {
	inst.startDelay(time.Second * 2)
	return nil
}

func (inst *AutoMigratorManager) startDelay(t time.Duration) {
	go func() {
		time.Sleep(t)
		err := inst.migrate()
		if err != nil {
			vlog.Error(err)
		}
	}()
}

func (inst *AutoMigratorManager) migrate() error {

	defer func() {
		e := recover()
		if e != nil {
			vlog.Error(e)
		}
	}()

	all := inst.Items
	for _, item := range all {
		err := item.AutoMigrate()
		if err != nil {
			return err
		}
	}
	return nil
}
