package request

import (
	"time"

	ReqOrder "github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RequireOrderSearch struct {
	ReqOrder.RequireOrder
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	// StartOrder_date  *time.Time  `json:"startOrder_date" form:"startOrder_date"`
	// EndOrder_date  *time.Time  `json:"endOrder_date" form:"endOrder_date"`
	PONumber        string   `json:"poNumber" form:"poNumber"`
	StartQuantity   *int     `json:"startQuantity" form:"startQuantity"`
	EndQuantity     *int     `json:"endQuantity" form:"endQuantity"`
	StartTotalCbm   *float64 `json:"startTotalCbm" form:"startTotalCbm"`
	EndTotalCbm     *float64 `json:"endTotalCbm" form:"endTotalCbm"`
	StartTotalPrice *float64 `json:"startTotalPrice" form:"startTotalPrice"`
	EndTotalPrice   *float64 `json:"endTotalPrice" form:"endTotalPrice"`
	request.PageInfo
}

type ReqOrderExcelInfo struct {
	ReqOrderId *int `json:"reqOrderId" form:"reqOrderId"`
}
