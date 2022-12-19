package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UploadSubOrderRouter struct {
}

// InitUploadSubOrderRouter 初始化 UploadSubOrder 路由信息
func (s *UploadSubOrderRouter) InitUploadSubOrderRouter(Router *gin.RouterGroup) {
	upSubOrdRouter := Router.Group("upSubOrd").Use(middleware.OperationRecord())
	upSubOrdRouterWithoutRecord := Router.Group("upSubOrd")
	var upSubOrdApi = v1.ApiGroupApp.OrderApiGroup.UploadSubOrderApi
	{
		upSubOrdRouter.POST("createUploadSubOrder", upSubOrdApi.CreateUploadSubOrder)             // 新建UploadSubOrder
		upSubOrdRouter.DELETE("deleteUploadSubOrder", upSubOrdApi.DeleteUploadSubOrder)           // 删除UploadSubOrder
		upSubOrdRouter.DELETE("deleteUploadSubOrderByIds", upSubOrdApi.DeleteUploadSubOrderByIds) // 批量删除UploadSubOrder
		upSubOrdRouter.PUT("updateUploadSubOrder", upSubOrdApi.UpdateUploadSubOrder)              // 更新UploadSubOrder
		upSubOrdRouter.POST("uploadSubOrderExcel", upSubOrdApi.UploadSubOrderExcel)               // 上传EXCEL
		upSubOrdRouter.POST("doExecImport", upSubOrdApi.DoExecImport)                             // 确认导入
	}
	{
		upSubOrdRouterWithoutRecord.GET("findUploadSubOrder", upSubOrdApi.FindUploadSubOrder)       // 根据ID获取UploadSubOrder
		upSubOrdRouterWithoutRecord.GET("getUploadSubOrderList", upSubOrdApi.GetUploadSubOrderList) // 获取UploadSubOrder列表
	}
}
