package order

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SubOrderRouter struct {
}

// InitSubOrderRouter 初始化 SubOrder 路由信息
func (s *SubOrderRouter) InitSubOrderRouter(Router *gin.RouterGroup) {
	subOrdRouter := Router.Group("subOrd").Use(middleware.OperationRecord())
	subOrdRouterWithoutRecord := Router.Group("subOrd")
	var subOrdApi = v1.ApiGroupApp.OrderApiGroup.SubOrderApi
	{
		subOrdRouter.POST("createSubOrder", subOrdApi.CreateSubOrder)                  // 新建SubOrder
		subOrdRouter.DELETE("deleteSubOrder", subOrdApi.DeleteSubOrder)                // 删除SubOrder
		subOrdRouter.DELETE("deleteSubOrderByIds", subOrdApi.DeleteSubOrderByIds)      // 批量删除SubOrder
		subOrdRouter.PUT("updateSubOrder", subOrdApi.UpdateSubOrder)                   // 更新SubOrder
		subOrdRouter.PUT("addSubOrderByProductIds", subOrdApi.AddSubOrderByProductIds) // 根据商品IDs，添加商品项
	}
	{
		subOrdRouterWithoutRecord.GET("findSubOrder", subOrdApi.FindSubOrder)       // 根据ID获取SubOrder
		subOrdRouterWithoutRecord.GET("getSubOrderList", subOrdApi.GetSubOrderList) // 获取SubOrder列表
	}
}
