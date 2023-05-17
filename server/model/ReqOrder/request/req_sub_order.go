package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReqSubOrderSearch struct {
	ReqOrder.ReqSubOrder
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ReqOrderIdAndIdsReq struct {
	ReqOrderId string `json:"reqOrderId" form:"reqOrderId"`
	Ids        []int  `json:"ids" form:"ids"`
}

type ReqOrderIdReq struct {
	ReqOrderId string `json:"reqOrderId" form:"reqOrderId"`
}
