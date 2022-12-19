package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"time"
)

type SubOrderSearch struct {
	order.SubOrder
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartExp_date  *time.Time `json:"startExp_date" form:"startExp_date"`
	EndExp_date    *time.Time `json:"endExp_date" form:"endExp_date"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type OrderIdAndIdsReq struct {
	OrderId string `json:"orderId" form:"orderId"`
	Ids     []int  `json:"ids" form:"ids"`
}

type OrderIdReq struct {
	OrderId string `json:"order_id" form:"order_id"`
}
