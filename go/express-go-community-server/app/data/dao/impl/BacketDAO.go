package impl

import (
	"errors"

	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/data/entity"
	"github.com/bitwomrhole-express/express/community-server/app/service"
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
)

type BucketDaoImpl struct {
	markup.Component `id:"express-data-bucket-dao" class:"express-server-data-auto-migrator"`

	DS            datasource.Source     `inject:"#gorm-datasource-default"`
	UUIDGenerator service.UUIDGenerator `inject:"#the-uuid-generator"`
}

func (inst *BucketDaoImpl) _Impl() (dao.BucketDAO, AutoMigrator) {
	return inst, inst
}

func (inst *BucketDaoImpl) AutoMigrate() error {
	db := inst.DS.DB()
	return db.AutoMigrate(&entity.Bucket{})
}

func (inst *BucketDaoImpl) ListBucketsByOwner(owner dxo.AccountID) ([]*entity.Bucket, error) {
	return nil, errors.New("no impl")
}
