package dto

import (
	"bitwomrhole.com/djaf/express-go-server/server/data/dxo"
	"github.com/bitwormhole/starter-restful/api/dto"
)

// type PackageID PackageID // 快递包裹编号（运单号）

// Package 表示一件包裹
type Package struct {
	dto.BaseDTO

	ID  dxo.PackageID `json:"id"`
	Num string        `json:"num"` // 运单号

	FromAddress string `json:"from_address"` // 寄件人邮箱地址
	ToAddress   string `json:"to_address"`   // 收件人邮箱地址
	Comment     string `json:"comment"`      // 备注信息
	Transport   string `json:"transport"`    // 运输方式
	Status      string `json:"status"`       // 当前状态
}
