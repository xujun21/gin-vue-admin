package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"time"
)

type OrderSearch struct {
	order.Order
	StartCreatedAt     *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt       *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartOrder_date    *time.Time `json:"startOrder_date" form:"startOrder_date"`
	EndOrder_date      *time.Time `json:"endOrder_date" form:"endOrder_date"`
	StartInvoice_date  *time.Time `json:"startInvoice_date" form:"startInvoice_date"`
	EndInvoice_date    *time.Time `json:"endInvoice_date" form:"endInvoice_date"`
	StartQuantity      *int       `json:"startQuantity" form:"startQuantity"`
	EndQuantity        *int       `json:"endQuantity" form:"endQuantity"`
	StartSubtotal      *float64   `json:"startSubtotal" form:"startSubtotal"`
	EndSubtotal        *float64   `json:"endSubtotal" form:"endSubtotal"`
	StartVat           *float64   `json:"startVat" form:"startVat"`
	EndVat             *float64   `json:"endVat" form:"endVat"`
	StartDiscount      *float64   `json:"startDiscount" form:"startDiscount"`
	EndDiscount        *float64   `json:"endDiscount" form:"endDiscount"`
	StartOrder_total   *float64   `json:"startOrder_total" form:"startOrder_total"`
	EndOrder_total     *float64   `json:"endOrder_total" form:"endOrder_total"`
	StartHand_discount *float64   `json:"startHand_discount" form:"startHand_discount"`
	EndHand_discount   *float64   `json:"endHand_discount" form:"endHand_discount"`
	StartCustomer_id   *int       `json:"startCustomer_id" form:"startCustomer_id"`
	EndCustomer_id     *int       `json:"endCustomer_id" form:"endCustomer_id"`
	request.PageInfo
	Sort     string `json:"sort" form:"sort"`
	OrderStr string `json:"order" form:"order"`
}

type OrderExcelInfo struct {
	//FileName string `json:"fileName" form:"fileName"`
	OrderId *int `json:"orderId" form:"orderId"`
}
