package entity

// Package 表示一个包裹的ID，即运单号
type PackageID string

// Package 表示一个包裹
type Package struct {
	Base

	ID PackageID `gorm:"primaryKey"`

	FromAddress string
	ToAddress   string
	Comment     string
	Status      string
	Transport   string
}

func (id PackageID) String() string {
	return string(id)
}
