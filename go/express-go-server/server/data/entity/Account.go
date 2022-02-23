package entity

type AccountID int64

type Account struct {
	Base

	ID AccountID `gorm:"primaryKey"`

	Avatar   string
	Email    string `gorm:"unique"`
	Nickname string
	Password string
	Roles    string
	Username string `gorm:"unique"`
}
