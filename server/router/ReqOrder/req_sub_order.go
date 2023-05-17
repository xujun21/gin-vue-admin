package ReqOrder

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReqSubOrderRouter struct {
}

// InitReqSubOrderRouter 初始化 ReqSubOrder 路由信息
func (s *ReqSubOrderRouter) InitReqSubOrderRouter(Router *gin.RouterGroup) {
	reqSubOrderRouter := Router.Group("reqSubOrder").Use(middleware.OperationRecord())
	reqSubOrderRouterWithoutRecord := Router.Group("reqSubOrder")
	var reqSubOrderApi = v1.ApiGroupApp.ReqorderApiGroup.ReqSubOrderApi
	{
		reqSubOrderRouter.POST("createReqSubOrder", reqSubOrderApi.CreateReqSubOrder)             // 新建ReqSubOrder
		reqSubOrderRouter.DELETE("deleteReqSubOrder", reqSubOrderApi.DeleteReqSubOrder)           // 删除ReqSubOrder
		reqSubOrderRouter.DELETE("deleteReqSubOrderByIds", reqSubOrderApi.DeleteReqSubOrderByIds) // 批量删除ReqSubOrder
		reqSubOrderRouter.PUT("updateReqSubOrder", reqSubOrderApi.UpdateReqSubOrder)              // 更新ReqSubOrder
		reqSubOrderRouter.PUT("addReqSubOrderByProductIds", reqSubOrderApi.AddReqSubOrderByProductIds)
	}
	{
		reqSubOrderRouterWithoutRecord.GET("findReqSubOrder", reqSubOrderApi.FindReqSubOrder)       // 根据ID获取ReqSubOrder
		reqSubOrderRouterWithoutRecord.GET("getReqSubOrderList", reqSubOrderApi.GetReqSubOrderList) // 获取ReqSubOrder列表
	}
}
