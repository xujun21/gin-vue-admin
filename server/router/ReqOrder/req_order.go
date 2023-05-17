package ReqOrder

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RequireOrderRouter struct {
}

// InitRequireOrderRouter 初始化 RequireOrder 路由信息
func (s *RequireOrderRouter) InitRequireOrderRouter(Router *gin.RouterGroup) {
	ReqOrdRouter := Router.Group("ReqOrd").Use(middleware.OperationRecord())
	ReqOrdRouterWithoutRecord := Router.Group("ReqOrd")
	var ReqOrdApi = v1.ApiGroupApp.ReqorderApiGroup.RequireOrderApi
	{
		ReqOrdRouter.POST("createRequireOrder", ReqOrdApi.CreateRequireOrder)             // 新建RequireOrder
		ReqOrdRouter.DELETE("deleteRequireOrder", ReqOrdApi.DeleteRequireOrder)           // 删除RequireOrder
		ReqOrdRouter.DELETE("deleteRequireOrderByIds", ReqOrdApi.DeleteRequireOrderByIds) // 批量删除RequireOrder
		ReqOrdRouter.PUT("updateRequireOrder", ReqOrdApi.UpdateRequireOrder)              // 更新RequireOrder
		ReqOrdRouter.POST("exportRequireExcel", ReqOrdApi.ExportRequireExcel)
	}
	{
		ReqOrdRouterWithoutRecord.GET("findRequireOrder", ReqOrdApi.FindRequireOrder)       // 根据ID获取RequireOrder
		ReqOrdRouterWithoutRecord.GET("getRequireOrderList", ReqOrdApi.GetRequireOrderList) // 获取RequireOrder列表
	}
}
