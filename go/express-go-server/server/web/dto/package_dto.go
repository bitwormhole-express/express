package dto

import "github.com/bitwormhole/starter-restful/api/dto"

type PackageID string // 快递包裹编号（运单号）

type Package struct {
	dto.BaseDTO

	ID PackageID `json:"id"` // 运单号

	FromAddress string `json:"from_address"` // 寄件人邮箱地址
	ToAddress   string `json:"to_address"`   // 收件人邮箱地址
	Comment     string `json:"comment"`      // 备注信息
	Transport   string `json:"transport"`    // 运输方式
	Status      string `json:"status"`       // 当前状态
}
