package dao

import (
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/data/entity"
	"github.com/bitwormhole/starter-object-bucket/buckets"
)

// Account 是访问 entity.Account 的 DAO
// 【inject:"#express-data-account-dao"】
type Account interface {
	// 增
	Insert(e *entity.Account) (*entity.Account, error)

	// 改
	Update(id dxo.AccountID, e *entity.Account) (*entity.Account, error)

	// 删除
	Remove(id dxo.AccountID) error

	// 通过email查找账号
	FindByEmail(email string) (*entity.Account, error)

	// 通过ID查找账号
	FindByID(id dxo.AccountID) (*entity.Account, error)

	// 通过用户名查找账号
	FindByUserName(username string) (*entity.Account, error)

	// 通过(email 或 id 或 username)查找账号
	Find(word string) (*entity.Account, error)
}

// BucketAgent 为app提供访问云对象存储的接口
// 【inject:"#express-data-bucket-agent"】
type BucketAgent interface {
	GetBucket() buckets.Connection
}

// BucketDAO  是用来存取bucket配置的DAO
// 【inject:"#express-data-bucket-dao"】
type BucketDAO interface {
	ListBucketsByOwner(owner dxo.AccountID) ([]*entity.Bucket, error)
}

// PackageDAO 是用来存取包裹对象的DAO
// 【inject:"#express-data-package-dao"】
type PackageDAO interface {

	// 增
	Insert(o *entity.Package) (*entity.Package, error)

	// 删
	Delete(id dxo.PackageID) error

	// 改
	Update(id dxo.PackageID, o *entity.Package) (*entity.Package, error)

	// 查：列表
	FindList(o *entity.Package) ([]*entity.Package, error)

	// 查：单个对象
	FindOne(id dxo.PackageID) (*entity.Package, error)
}
