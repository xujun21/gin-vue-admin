package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

// InitOrderRouter 初始化 Order 路由信息
func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	ordRouter := Router.Group("ord").Use(middleware.OperationRecord())
	ordRouterWithoutRecord := Router.Group("ord")
	var ordApi = v1.ApiGroupApp.OrderApiGroup.OrderApi
	{
		ordRouter.POST("createOrder", ordApi.CreateOrder)                         // 新建Order
		ordRouter.DELETE("deleteOrder", ordApi.DeleteOrder)                       // 删除Order
		ordRouter.DELETE("deleteOrderByIds", ordApi.DeleteOrderByIds)             // 批量删除Order
		ordRouter.PUT("updateOrder", ordApi.UpdateOrder)                          // 更新Order
		ordRouter.POST("exportInvoiceExcel", ordApi.ExportInvoiceExcel)           // 下载发票
		ordRouter.POST("exportConfirmExcel", ordApi.ExportConfirmExcel)           // 下载发票
		ordRouter.POST("exportDeliveryNoteExcel", ordApi.ExportDeliveryNoteExcel) // 下载发票
	}
	{
		ordRouterWithoutRecord.GET("findOrder", ordApi.FindOrder)       // 根据ID获取Order
		ordRouterWithoutRecord.GET("getOrderList", ordApi.GetOrderList) // 获取Order列表
	}
}
