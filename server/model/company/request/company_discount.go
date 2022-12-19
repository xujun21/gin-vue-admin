package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
)

type CompanyDiscountSearch struct {
	company.CompanyDiscount
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type CompanyIdAndIdsReq struct {
	CompanyId string `json:"companyId" form:"companyId"`
	Ids       []int  `json:"ids" form:"ids"`
}
