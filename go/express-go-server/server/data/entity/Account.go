package entity

import (
	"strconv"

	"github.com/bitwormhole/starter/util"
)

// AccountID 表示一个账号的ID
type AccountID int64

// Account 表示一个账号
type Account struct {
	Base

	ID AccountID `gorm:"primaryKey"`

	Avatar   string
	Email    string `gorm:"index:idx_name,unique"`
	Nickname string
	Password string // sha256sum hex string |plain text
	Roles    string
	Salt     util.Base64
	Username string `gorm:"index:idx_name,unique"`
}

// TableName 设置表名
func (Account) TableName() string {
	return TableNameAccount
}

// Int64 把 AccountID 转换成 int64
func (id AccountID) Int64() int64 {
	return int64(id)
}

// String 把 AccountID 转换成 string
func (id AccountID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}
