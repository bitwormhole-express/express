package impl

import (
	"strconv"

	"github.com/bitwomrhole-express/express/community-server/app/data/dao"
	"github.com/bitwomrhole-express/express/community-server/app/data/dxo"
	"github.com/bitwomrhole-express/express/community-server/app/data/entity"
	"github.com/bitwomrhole-express/express/community-server/app/service"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
)

///////////////////////////////////////////////////////////////////////////////

type AccountDaoImpl struct {
	markup.Component `id:"express-data-account-dao" class:"express-server-data-auto-migrator"`

	DS            datasource.Source     `inject:"#gorm-datasource-default"`
	UUIDGenerator service.UUIDGenerator `inject:"#the-uuid-generator"`
}

func (inst *AccountDaoImpl) _Impl() (dao.Account, AutoMigrator) {
	return inst, inst
}

func (inst *AccountDaoImpl) AutoMigrate() error {
	db := inst.DS.DB()
	result := db.AutoMigrate(&entity.Account{})
	return result
}

// 增
func (inst *AccountDaoImpl) Insert(e *entity.Account) (*entity.Account, error) {

	e.UUID = inst.UUIDGenerator.GenUUID("entity.Account", e.Email)

	db := inst.DS.DB()
	result := db.Create(e)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

// 改
func (inst *AccountDaoImpl) Update(id dxo.AccountID, o1 *entity.Account) (*entity.Account, error) {

	db := inst.DS.DB()
	o2, err := inst.FindByID(id)
	if err != nil {
		return nil, err
	}

	o2.Avatar = o1.Avatar
	o2.Nickname = o1.Nickname
	o2.Password = o1.Password
	o2.Roles = o1.Roles
	o2.Salt = o1.Salt

	result := db.Save(o2)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return o2, nil
}

// 删除
func (inst *AccountDaoImpl) Remove(id dxo.AccountID) error {
	db := inst.DS.DB()
	result := db.Delete(&entity.Account{}, id)
	return result.Error
}

// 通过ID查找账号
func (inst *AccountDaoImpl) FindByID(id dxo.AccountID) (*entity.Account, error) {
	e := &entity.Account{}
	db := inst.DS.DB()
	result := db.First(e, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return e, nil
}

// 通过email查找账号
func (inst *AccountDaoImpl) FindByEmail(email string) (*entity.Account, error) {
	e := &entity.Account{}
	db := inst.DS.DB()
	result := db.Where("email = ?", email).First(e)
	if result.Error != nil {
		return nil, result.Error
	}
	return e, nil
}

// 通过用户名查找账号
func (inst *AccountDaoImpl) FindByUserName(username string) (*entity.Account, error) {
	e := &entity.Account{}
	db := inst.DS.DB()
	result := db.Where("username = ?", username).First(e)
	if result.Error != nil {
		return nil, result.Error
	}
	return e, nil
}

// 通过email或者id查找账号
func (inst *AccountDaoImpl) Find(word string) (*entity.Account, error) {

	// by id
	n, err := strconv.ParseInt(word, 10, 64)
	if err == nil {
		account, err := inst.FindByID(dxo.AccountID(n))
		if err == nil {
			return account, nil
		}
	}

	// by username
	account, err := inst.FindByUserName(word)
	if err == nil {
		return account, nil
	}

	// by email
	return inst.FindByEmail(word)
}
