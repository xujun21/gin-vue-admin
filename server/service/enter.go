package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/company"
	"github.com/flipped-aurora/gin-vue-admin/server/service/customer"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/order"
	"github.com/flipped-aurora/gin-vue-admin/server/service/product"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	ProductServiceGroup  product.ServiceGroup
	OrderServiceGroup    order.ServiceGroup
	CustomerServiceGroup customer.ServiceGroup
	CompanyServiceGroup  company.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
