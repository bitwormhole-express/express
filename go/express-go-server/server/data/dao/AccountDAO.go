package dao

import (
	"strconv"

	"bitwomrhole.com/djaf/express-go-server/server/data/entity"
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
)

// Account 是访问 entity.Account 的 DAO
// 【inject:"#express-data-account-dao"】
type Account interface {
	// 增
	Insert(e *entity.Account) (*entity.Account, error)

	// 改
	Update(id entity.AccountID, e *entity.Account) (*entity.Account, error)

	// 删除
	Remove(id entity.AccountID) error

	// 通过email查找账号
	FindByEmail(email string) (*entity.Account, error)

	// 通过ID查找账号
	FindByID(id entity.AccountID) (*entity.Account, error)

	// 通过用户名查找账号
	FindByUserName(username string) (*entity.Account, error)

	// 通过(email 或 id 或 username)查找账号
	Find(word string) (*entity.Account, error)
}

///////////////////////////////////////////////////////////////////////////////

type AccountDaoImpl struct {
	markup.Component `id:"express-data-account-dao" class:"express-server-data-auto-migrator"`

	DS datasource.Source `inject:"#gorm-datasource-default"`
}

func (inst *AccountDaoImpl) _Impl() (Account, AutoMigrator) {
	return inst, inst
}

func (inst *AccountDaoImpl) AutoMigrate() error {
	db := inst.DS.DB()
	result := db.AutoMigrate(&entity.Account{})
	return result
}

// 增
func (inst *AccountDaoImpl) Insert(e *entity.Account) (*entity.Account, error) {
	db := inst.DS.DB()
	result := db.Create(e)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

// 改
func (inst *AccountDaoImpl) Update(id entity.AccountID, e *entity.Account) (*entity.Account, error) {
	e.ID = id
	db := inst.DS.DB()
	result := db.Save(e)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

// 删除
func (inst *AccountDaoImpl) Remove(id entity.AccountID) error {
	db := inst.DS.DB()
	result := db.Delete(&entity.Account{}, id)
	return result.Error
}

// 通过ID查找账号
func (inst *AccountDaoImpl) FindByID(id entity.AccountID) (*entity.Account, error) {
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
		account, err := inst.FindByID(entity.AccountID(n))
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
