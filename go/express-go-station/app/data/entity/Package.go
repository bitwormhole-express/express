package entity

import (
	"bitwomrhole.com/djaf/express-go-server/server/data/dxo"
)

////////////////////////////////////////////////////////////////////////////////

// 定义传输方式
const (
	TransportLocalhost = "localhost" // 本地直连
	TransportAliyun    = "aliyun"    // 阿里云
	TransportTencent   = "tencent"   // 腾讯云
	TransportBaidu     = "baidu"     // 百度云
	TransportHuawei    = "huawei"    // 华为云
)

////////////////////////////////////////////////////////////////////////////////

// Package 表示一个包裹
type Package struct {
	Base

	ID dxo.PackageID `gorm:"primaryKey"`

	Num         string // 运单号
	FromAddress string
	ToAddress   string
	Comment     string
	Status      string
	Transport   string
	Manifest    dxo.FileManifestValue
}

// TableName 设置表名
func (Package) TableName() string {
	return TableNamePackage
}

////////////////////////////////////////////////////////////////////////////////
