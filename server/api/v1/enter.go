package v1

import (
	reqOrder "github.com/flipped-aurora/gin-vue-admin/server/api/v1/ReqOrder"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/company"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/customer"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/order"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/product"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	ProductApiGroup  product.ApiGroup
	OrderApiGroup    order.ApiGroup
	CustomerApiGroup customer.ApiGroup
	CompanyApiGroup  company.ApiGroup
	ReqorderApiGroup reqOrder.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
