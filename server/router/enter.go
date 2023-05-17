package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ReqOrder"
	"github.com/flipped-aurora/gin-vue-admin/server/router/company"
	"github.com/flipped-aurora/gin-vue-admin/server/router/customer"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/order"
	"github.com/flipped-aurora/gin-vue-admin/server/router/product"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Product  product.RouterGroup
	Order    order.RouterGroup
	Customer customer.RouterGroup
	Company  company.RouterGroup
	Reqorder ReqOrder.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
