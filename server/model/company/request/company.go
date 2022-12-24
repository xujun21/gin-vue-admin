package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	"time"
)

type CompanySearch struct {
	company.Company
	Company_name   string     `json:"customer_company_name" form:"customer_company_name"`
	Contact_name   string     `json:"customer_contact_name" form:"customer_contact_name"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
