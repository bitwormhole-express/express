package entity

import (
	"bitwomrhole.com/djaf/express-go-server/server/data/dxo"
	"github.com/bitwormhole/starter/util"
)

// Account 表示一个账号
type Account struct {
	Base

	ID dxo.AccountID `gorm:"primaryKey"`

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
