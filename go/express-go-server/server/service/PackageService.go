package service

import "bitwomrhole.com/djaf/express-go-server/server/web/dto"

// PackageService 是处理包裹对象的服务
// 【inject:"#express-PackageService"】
type PackageService interface {
	Find(o *dto.Package) ([]*dto.Package, error)

	FindOne(id dto.PackageID) (*dto.Package, error)

	Insert(o *dto.Package) (*dto.Package, error)

	Remove(id dto.PackageID) error

	Update(id dto.PackageID, o *dto.Package) (*dto.Package, error)
}
