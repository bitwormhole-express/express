package dao

import (
	"crypto/rand"
	"errors"
	"strconv"
	"strings"
	"sync"

	"bitwomrhole.com/djaf/express-go-server/server/data/entity"
	"bitwomrhole.com/djaf/express-go-server/server/service"
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// PackageDAO 是用来存取包裹对象的DAO
// 【inject:"#express-data-package-dao"】
type PackageDAO interface {

	// 增
	Insert(o *entity.Package) (*entity.Package, error)

	// 删
	Delete(id entity.PackageID) error

	// 改
	Update(id entity.PackageID, o *entity.Package) (*entity.Package, error)

	// 查：列表
	FindList(o *entity.Package) ([]*entity.Package, error)

	// 查：单个对象
	FindOne(id entity.PackageID) (*entity.Package, error)
}

///////////////////////////////////////////////////////////////////////////////

type PackageDaoImpl struct {
	markup.Component `id:"express-data-package-dao" class:"express-server-data-auto-migrator"`

	DS            datasource.Source     `inject:"#gorm-datasource-default"`
	UUIDGenerator service.UUIDGenerator `inject:"#the-uuid-generator"`

	idgen innerPackageIDGen
}

func (inst *PackageDaoImpl) _Impl() (PackageDAO, AutoMigrator) {
	return inst, inst
}

func (inst *PackageDaoImpl) AutoMigrate() error {
	db := inst.DS.DB()
	result := db.AutoMigrate(&entity.Package{})
	return result
}

// 增
func (inst *PackageDaoImpl) Insert(o *entity.Package) (*entity.Package, error) {

	o.UUID = inst.UUIDGenerator.GenUUID("entity.Package", o.ID.String())
	o.ID = inst.idgen.makeID(o)

	db := inst.DS.DB()
	result := db.Create(o)
	if result.Error != nil {
		return nil, result.Error
	}
	return o, nil
}

// 删
func (inst *PackageDaoImpl) Delete(id entity.PackageID) error {
	db := inst.DS.DB()
	result := db.Delete(&entity.Package{}, id)
	return result.Error
}

// 改
func (inst *PackageDaoImpl) Update(id entity.PackageID, o1 *entity.Package) (*entity.Package, error) {
	db := inst.DS.DB()
	o2, err := inst.FindOne(id)
	if err != nil {
		return nil, err
	}

	o2.Comment = o1.Comment
	o2.Status = o1.Status
	o2.Transport = o2.Transport

	result := db.Save(o2)
	if result.Error != nil {
		return nil, result.Error
	}
	return o2, nil
}

// 查：列表
func (inst *PackageDaoImpl) FindList(o *entity.Package) ([]*entity.Package, error) {

	// db := inst.DS.DB()
	// o2 := &entity.Package{}
	// result := db.Where().First(o2, id)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return o2, nil

	return nil, errors.New("no impl")
}

// 查：单个对象
func (inst *PackageDaoImpl) FindOne(id entity.PackageID) (*entity.Package, error) {
	db := inst.DS.DB()
	o2 := &entity.Package{}
	result := db.First(o2, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return o2, nil
}

///////////////////////////////////////////////////////////////////////////////

type innerPackageIDGen struct {
	mutex sync.Mutex
	index int64
}

func (inst *innerPackageIDGen) nextIndex() int64 {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()
	inst.index++
	return inst.index
}

func (inst *innerPackageIDGen) makeID(o *entity.Package) entity.PackageID {

	nonce := make([]byte, 7)
	index := inst.nextIndex()
	now := util.Now().Int64()

	rand.Reader.Read(nonce)

	builder := strings.Builder{}
	builder.WriteString(strconv.FormatInt(now, 10))
	builder.WriteString("-")
	builder.WriteString(strconv.FormatInt(index, 10))
	builder.WriteString("-")
	builder.WriteString(util.StringifyBytes(nonce))

	str := builder.String()
	return entity.PackageID(str)
}
