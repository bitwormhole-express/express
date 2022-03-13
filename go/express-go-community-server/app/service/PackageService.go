package service

import (
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/web/dto"
)

// PackageService 是处理包裹对象的服务
// 【inject:"#express-PackageService"】
type PackageService interface {
	Find(o *dto.Package) ([]*dto.Package, error)

	FindOne(id dxo.PackageID) (*dto.Package, error)

	Insert(o *dto.Package) (*dto.Package, error)

	Remove(id dxo.PackageID) error

	Update(id dxo.PackageID, o *dto.Package) (*dto.Package, error)
}
